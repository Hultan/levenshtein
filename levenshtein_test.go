package levenshtein

import (
	"fmt"
	"testing"
	"time"
)

func TestLevPerf(t *testing.T) {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		Lev("Skruvmejsel", "hammare")
	}

	fmt.Println(time.Since(start).Milliseconds())
}

func TestLev(t *testing.T) {
	type args struct {
		source string
		dest   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"identical", args{"apple", "apple"}, 0},
		{"empty source", args{"", "apple"}, 5},
		{"empty dest", args{"apple", ""}, 5},
		{"add (before)", args{"apple", "dapple"}, 1},
		{"add (after)", args{"apple", "apples"}, 1},
		{"remove (before)", args{"apples", "pples"}, 1},
		{"remove (after)", args{"apples", "apple"}, 1},
		{"substitute (before)", args{"apples", "opples"}, 1},
		{"substitute (after)", args{"apples", "applet"}, 1},
		{"long", args{"sdfpioudfglkjaweoiulknsdasdasd", "cxvpouwejk,mnasdljkhasdlkhasdiu"}, 23},
		{"different", args{"foo", "bar"}, 3},
		{"unicode", args{"世界世界世界", "foo"}, 6},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Lev(tt.args.source, tt.args.dest); got != tt.want {
					t.Errorf("Lev() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_min(t *testing.T) {
	type args struct {
		lev1, lev2, lev3 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,1,1", args{1, 1, 1}, 1},
		{"3,1,2", args{3, 1, 2}, 1},
		{"3,2,1", args{3, 2, 1}, 1},
		{"1,2,3", args{1, 2, 3}, 1},
		{"-1,-1,-1", args{-1, -1, -1}, -1},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := min(tt.args.lev1, tt.args.lev2, tt.args.lev3); got != tt.want {
					t.Errorf("min() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
