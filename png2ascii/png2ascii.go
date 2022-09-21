package png2ascii

import (
	"image"

	"golang.org/x/image/draw"
)

type Codec interface {
	Convert(img image.Image, dest string, debug bool) error
}

func toAscii(img image.Image, charset string) []string {
	N := len(charset)
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	gray := toGrayScale(img)

	buffer := []string{}
	for y := 0; y < height; y++ {
		row := []rune{}
		for x := 0; x < width; x++ {
			pixel := gray.GrayAt(x, y)
			brightness := float64(pixel.Y) / 256.0
			ix := int(brightness * float64(N))

			ch := charset[ix]
			row = append(row, rune(ch))
		}

		buffer = append(buffer, string(row))
	}

	return buffer
}

func toGrayScale(img image.Image) *image.Gray {
	bounds := img.Bounds()

	gray := image.NewGray(bounds)
	draw.Draw(gray, bounds, img, bounds.Min, draw.Src)

	return gray
}
