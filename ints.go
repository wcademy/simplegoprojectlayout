package simplegoprojectlayout

//  Ints представляет собой слайс целых чисел
type Ints []int

// SliceToInts преобразует слайс интов во внутренний тип данных Ints
func SliceToInts(slice []int) Ints {
	ints := make(Ints, len(slice))

	for i := range slice {
		ints[i] = slice[i]
	}

	return slice
}
