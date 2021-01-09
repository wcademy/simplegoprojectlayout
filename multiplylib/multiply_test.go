package multiplylib_test

import (
	"testing"

	sgpl "github.com/wcademy/simplegoprojectlayout"
	"github.com/wcademy/simplegoprojectlayout/multiplylib"
)

func TestMultiply(t *testing.T) {
	type args struct {
		ints sgpl.Ints
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2x2",
			args{sgpl.Ints{2, 2}},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplylib.Multiply(tt.args.ints); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
