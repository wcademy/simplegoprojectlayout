package main

import (
	"fmt"

	"github.com/wcademy/simplegoprojectlayout"
	"github.com/wcademy/simplegoprojectlayout/sumlib"
)

func main() {
	ints := simplegoprojectlayout.Ints{1, 2, 3}
	result := sumlib.Sum(ints)
	fmt.Println(result)
}
