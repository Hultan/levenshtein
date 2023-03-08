package levenshtein

import (
	"testing"
)

func TestLev(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"/apple=5", args{"", "apple"}, 5},
		{"apple/=5", args{"apple", ""}, 5},
		{"apple/apple=0", args{"apple", "apple"}, 0},
		{"apple/dapple=1", args{"apple", "dapple"}, 1},
		{
			"alfabet/alfabet-1=1",
			args{"abcdefghijklmnopqrstuvxyzåäö", "abcdefghijklmnopqrstuvxyzåä"}, 1,
		},
		{"long/long=24", args{"sdfpioudfglkjaweoiulknsdasdasd", "cxvpouwejk,mnasdljkhasdlkhasdiu"}, 24},
		{"cattle/apple=1", args{"cattle", "apple"}, 3},
		{"foo/bar=3", args{"foo", "bar"}, 3},
		{"foo/foooooo=4", args{"foo", "foooooo"}, 4},
		{"foooooo/foo=4", args{"foooooo", "foo"}, 4},
		{"世界世界世界/foo=6", args{"世界世界世界", "foo"}, 6},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Lev(tt.args.s1, tt.args.s2); got != tt.want {
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
