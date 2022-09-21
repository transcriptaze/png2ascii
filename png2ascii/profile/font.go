package profile

import (
	"fmt"
	"strconv"
	"strings"
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
