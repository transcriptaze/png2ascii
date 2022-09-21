package profile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Squoosh struct {
	Enabled bool
	Width   *uint32
}

func (u Squoosh) String() string {
	if u.Width != nil {
		return fmt.Sprintf("width:%v", *u.Width)
	} else if u.Enabled {
		return "yes"
	} else {
		return "no"
	}
}

func (u *Squoosh) Set(s string) error {
	switch {
	case s == "yes":
		u.Enabled = true

	case s == "no":
		u.Enabled = false

	case strings.HasPrefix(s, "width"):
		if match := regexp.MustCompile("^width:([0-9]+)$").FindStringSubmatch(s); match == nil {
			return fmt.Errorf("invalid squoosh width (%v)", s)
		} else if len(match) < 2 {
			return fmt.Errorf("invalid squoosh width (%v)", s)
		} else if v, err := strconv.ParseUint(match[1], 10, 32); err != nil {
			return err
		} else {
			width := uint32(v)
			u.Width = &width
		}

	default:
		if enabled, err := strconv.ParseBool(s); err != nil {
			return err
		} else {
			u.Enabled = enabled
		}
	}

	return nil
}
