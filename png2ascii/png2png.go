package png2ascii

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/transcriptaze/png2ascii/png2ascii/profile"
)

type Png2Png struct {
	profile profile.Profile
	font    font.Face
}

func NewPng2Png(profile profile.Profile) (*Png2Png, error) {
	if face, err := profile.Font.Load(); err != nil {
		return nil, err
	} else {
		return &Png2Png{
			profile: profile,
			font:    face,
		}, nil
	}
}

func (m Png2Png) Convert(img image.Image, dest string, debug bool) error {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	if debug {
		ratio := float64(w) / float64(h)

		fmt.Printf("  ORIGINAL\n")
		fmt.Printf("    image size:   %vx%v\n", w, h)
		fmt.Printf("    aspect ratio: %.3f\n", ratio)
		fmt.Println()
	}

	var buffer []string

	fh := m.font.Metrics().Height
	fw, _ := m.font.GlyphAdvance('X')
	width := fixed.I(w).Mul(fw).Ceil()
	height := fixed.I(h).Mul(fh).Ceil()

	squoosh := m.profile.Squoosh

	if squoosh.Enabled || squoosh.Width != nil {
		scaled := m.squoosh(img, m.font, squoosh, debug)
		bounds := scaled.Bounds()
		w := bounds.Dx()
		h := bounds.Dy()
		width = fixed.I(w).Mul(fw).Ceil()
		height = fixed.I(h).Mul(fh).Ceil()

		if debug {
			fmt.Printf("  RENDER\n")
			fmt.Printf("    width:        %v\n", width)
			fmt.Printf("    height:       %v\n", height)
			fmt.Printf("    aspect ratio: %.3f\n", float64(width)/float64(height))
			fmt.Println()
		}

		buffer = toAscii(scaled, m.profile.Charset)
	} else {
		buffer = toAscii(img, m.profile.Charset)
	}

	origin := image.Point{X: 0, Y: 0}
	background := image.NewUniform(color.White)
	dst := image.NewGray(image.Rect(0, 0, width, height))

	pen := font.Drawer{
		Dst:  dst,
		Src:  image.Black,
		Face: m.font,
		Dot:  fixed.P(origin.X, origin.Y),
	}

	draw.Copy(dst, image.Point{0, 0}, background, dst.Bounds(), draw.Over, nil)

	for i, row := range buffer {
		x := origin.X
		y := origin.Y + fixed.I(i+1).Mul(fh).Ceil()

		pen.Dot = fixed.P(x, y)
		pen.DrawString(row)
	}

	if dest != "" {
		f, err := os.Create(dest)
		if err != nil {
			return err
		}

		defer f.Close()

		return png.Encode(f, dst)
	}

	return nil
}

func (m Png2Png) squoosh(img image.Image, font font.Face, squoosh profile.Squoosh, debug bool) image.Image {
	metrics := font.Metrics()
	fh := metrics.Height
	fw := fixed.I(0)

	for _, ch := range m.profile.Charset {
		if advance, _ := font.GlyphAdvance(ch); advance > fw {
			fw = advance
		}
	}

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

	w := fixed.I(W).Mul(fw).Ceil()
	h := fixed.I(H).Mul(fh).Ceil()

	hh := int(math.Round(float64(w*H) / float64(W)))
	width := W
	height := int(math.Round(float64(H*hh) / float64(h)))

	if debug {
		fmt.Printf("  SQUOOSH\n")
		fmt.Printf("    font.metrics      %+v\n", metrics)
		fmt.Printf("    font.wh           %vx%v  (%vx%v)\n", fw, fh, fw.Ceil(), fh.Ceil())
		fmt.Printf("    image.wh          %vx%v\n", W, H)
		fmt.Printf("    image.ratio       %.3f\n", float64(W)/float64(H))
		fmt.Printf("    unsquooshed.wh    %vx%v\n", w, h)
		fmt.Printf("    unsquooshed.ratio %.3f\n", float64(w)/float64(h))
		fmt.Printf("    squooshed.wh      %vx%v\n", width, height)
		fmt.Println()
	}

	rect := image.Rect(0, 0, width, height)
	squooshed := image.NewRGBA(rect)

	draw.CatmullRom.Scale(squooshed, squooshed.Rect, img, img.Bounds(), draw.Over, nil)

	return squooshed
}
