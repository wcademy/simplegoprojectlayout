package comandlineutils_test

import (
	"reflect"
	"testing"

	sgpl "github.com/wcademy/simplegoprojectlayout"
	"github.com/wcademy/simplegoprojectlayout/internal/comandlineutils"
)

func TestGetArgs(t *testing.T) {
	type args struct {
		allOSArgs []string
	}

	tests := []struct {
		name    string
		args    args
		want    sgpl.Ints
		wantErr bool
	}{
		{
			"success case with 1 2 3",
			args{[]string{"1", "2", "3"}},
			sgpl.Ints{1, 2, 3},
			false,
		},
		{
			"error case with wrong input",
			args{[]string{"1", "not a number", "3"}},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := comandlineutils.GetArgs(tt.args.allOSArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
