package sumlib_test

import (
	"testing"

	sgpl "github.com/wcademy/simplegoprojectlayout"
	"github.com/wcademy/simplegoprojectlayout/sumlib"
)

func TestSum(t *testing.T) {
	type args struct {
		ints sgpl.Ints
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2 + 2",
			args{sgpl.Ints{2, 2}},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumlib.Sum(tt.args.ints); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
