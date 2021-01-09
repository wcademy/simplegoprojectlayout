package sumlib

import (
	sgpl "github.com/wcademy/simplegoprojectlayout"
)

func Sum(ints sgpl.Ints) int {
	s := 0

	for _, val := range ints {
		s += val
	}

	return s
}
