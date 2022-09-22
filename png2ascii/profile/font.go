package profile

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	Typeface string  `json:"typeface"`
	Size     float64 `json:"size"`
	DPI      float64 `json:"dpi"`
}

func (f Font) String() string {
	return fmt.Sprintf("%v:%v:%v", f.Typeface, f.Size, f.DPI)
}

func (f *Font) Set(s string) error {
	tokens := strings.Split(s, ":")

	if len(tokens) > 0 && strings.TrimSpace(tokens[0]) != "" {
		f.Typeface = strings.TrimSpace(tokens[0])
	}

	if len(tokens) > 1 && strings.TrimSpace(tokens[1]) != "" {
		if v, err := strconv.ParseFloat(tokens[1], 64); err != nil {
			return err
		} else if v < 0.0 {
			return fmt.Errorf("invalid font size (%v)", tokens[1])
		} else {
			f.Size = v
		}
	}

	if len(tokens) > 2 && strings.TrimSpace(tokens[2]) != "" {
		if v, err := strconv.ParseFloat(tokens[2], 64); err != nil {
			return err
		} else if v < 0.0 {
			return fmt.Errorf("invalid DPI (%v)", tokens[2])
		} else {
			f.DPI = v
		}
	}

	return nil
}

func (f *Font) Load() (font.Face, error) {
	options := opentype.FaceOptions{
		Size:    f.Size,
		DPI:     f.DPI,
		Hinting: font.HintingNone,
	}

	var bytes []byte
	var err error

	switch f.Typeface {
	case "gomono":
		bytes = gomono.TTF

	case "gomonobold":
		bytes = gomonobold.TTF

	case "gomonoitalic":
		bytes = gomonoitalic.TTF

	case "gomonobolditalic":
		bytes = gomonobolditalic.TTF

	default:
		if bytes, err = os.ReadFile(f.Typeface); err != nil {
			return nil, err
		}
	}

	if font, err := opentype.Parse(bytes); err != nil {
		return nil, err
	} else {
		return opentype.NewFace(font, &options)
	}
}
