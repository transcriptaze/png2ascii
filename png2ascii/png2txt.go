package png2ascii

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"os"

	"golang.org/x/image/draw"

	"github.com/transcriptaze/png2ascii/png2ascii/profile"
)

type Png2Txt struct {
	profile profile.Profile
}

func NewPng2Txt(profile profile.Profile) (Png2Txt, error) {
	return Png2Txt{
		profile: profile,
	}, nil
}

func (m Png2Txt) Convert(img image.Image, dest string, debug bool) error {
	if debug {
		bounds := img.Bounds()
		w := bounds.Dx()
		h := bounds.Dy()
		ratio := float64(w) / float64(h)

		fmt.Printf("  ORIGINAL\n")
		fmt.Printf("    image size:   %vx%v\n", w, h)
		fmt.Printf("    aspect ratio: %.3f\n", ratio)
		fmt.Println()
	}

	var buffer []string

	squoosh := m.profile.Squoosh

	if squoosh.Enabled || squoosh.Width != nil {
		squashed := m.squoosh(img, squoosh, debug)
		buffer = toAscii(squashed, m.profile.Charset)
	} else {
		buffer = toAscii(img, m.profile.Charset)
	}

	var b bytes.Buffer

	for _, line := range buffer {
		fmt.Fprintf(&b, "%v\n", line)
	}

	if dest != "" {
		return os.WriteFile(dest, b.Bytes(), 0660)
	}

	fmt.Printf("%v", string(b.Bytes()))

	return nil
}

func (m Png2Txt) squoosh(img image.Image, squoosh profile.Squoosh, debug bool) image.Image {
	bounds := img.Bounds()
	W := bounds.Dx()
	H := bounds.Dy()

	if squoosh.Width != nil {
		w := float64(W)
		h := float64(H)
		r := h / w
		w = float64(*squoosh.Width)
		h = math.Round(r * w)

		W = int(math.Round(w))
		H = int(math.Round(h))
	}

	width := W
	height := H / 2

	if debug {
		fmt.Printf("  SQUOOSH\n")
		fmt.Printf("    image.wh    %vx%v\n", W, H)
		fmt.Printf("    image.ratio %.3f\n", float64(W)/float64(H))
		fmt.Printf("    squashed.wh %vx%v\n", width, height)
		fmt.Println()
	}

	rect := image.Rect(0, 0, width, height)
	squashed := image.NewRGBA(rect)

	draw.CatmullRom.Scale(squashed, squashed.Rect, img, img.Bounds(), draw.Over, nil)

	return squashed
}
