package png2ascii

import (
	"fmt"
	"strconv"
)

type Uint32 struct {
	v  uint32
	ok bool
}

func (u Uint32) String() string {
	if u.ok {
		return fmt.Sprintf("%v", u.v)
	} else {
		return ""
	}
}

func (u *Uint32) Set(s string) error {
	if v, err := strconv.ParseUint(s, 10, 32); err != nil {
		return err
	} else {
		u.v = uint32(v)
		u.ok = true
	}

	return nil
}
