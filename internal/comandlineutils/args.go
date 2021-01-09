package comandlineutils

import (
	"strconv"

	sgpl "github.com/wcademy/simplegoprojectlayout"
)

func GetArgs(allOSArgs []string) (sgpl.Ints, error) {
	ints := make(sgpl.Ints, len(allOSArgs))

	for i, arg := range allOSArgs {
		number, err := strconv.Atoi(arg)
		if err != nil {
			return nil, IntConversionErr{arg, err}
		}

		ints[i] = number
	}

	return ints, nil
}
