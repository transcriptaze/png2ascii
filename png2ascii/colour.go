package png2ascii

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"

	"golang.org/x/image/colornames"
)

var Black = Colour{
	colour: colornames.Black,
}

var White = Colour{
	colour: colornames.White,
}

type Colour struct {
	colour color.RGBA
}

func (c Colour) String() string {
	return fmt.Sprintf("#%02x%02x%02x%02x", c.colour.R, c.colour.G, c.colour.B, c.colour.A)
}

func (u *Colour) Set(s string) error {
	if match := regexp.MustCompile("^#([a-fAF0-9]{6})$").FindStringSubmatch(s); match != nil {
		if v, err := strconv.ParseUint(match[1], 16, 32); err != nil {
			return err
		} else {
			u.colour = color.RGBA{
				R: uint8((v & 0x00ff0000) >> 16),
				G: uint8((v & 0x0000ff00) >> 8),
				B: uint8((v & 0x000000ff) >> 0),
				A: 0xff,
			}

			return nil
		}
	}

	if match := regexp.MustCompile("^#([a-fAF0-9]{8})$").FindStringSubmatch(s); match != nil {
		if v, err := strconv.ParseUint(match[1], 16, 32); err != nil {
			return err
		} else {
			u.colour = color.RGBA{
				R: uint8((v & 0xff000000) >> 24),
				G: uint8((v & 0x00ff0000) >> 16),
				B: uint8((v & 0x0000ff00) >> 8),
				A: uint8((v & 0x000000ff) >> 0),
			}

			return nil
		}

		return nil
	}

	return fmt.Errorf("invalid colour")
}
