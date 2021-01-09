package multiplylib

import (
	sgpl "github.com/wcademy/simplegoprojectlayout"
)

func Multiply(ints sgpl.Ints) int {
	product := 1

	for _, val := range ints {
		product *= val
	}

	return product
}
