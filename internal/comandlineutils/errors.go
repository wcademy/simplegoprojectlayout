package comandlineutils

import "fmt"

// IntConversionErr сигнализирует об ошибке преобразование аргументов командной строки

type IntConversionErr struct {
	Arg string
	Err error
}

func (e IntConversionErr) Unwrap() error {
	return e.Err
}

func (e IntConversionErr) Error() string {
	return fmt.Sprintf("can't convert %v to int: %s", e.Arg, e.Err.Error())
}
