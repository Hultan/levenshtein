package levenshtein

import (
	"testing"
)

func TestLevRecursive(t *testing.T) {
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
		{"long/long=24", args{"sdfpioudfglkjaweoiulknsdasdasd", "cxvpouwejk,mnasdljkhasdlkhasdiu"}, 23},
		{"cattle/apple=1", args{"cattle", "apple"}, 3},
		{"foo/bar=3", args{"foo", "bar"}, 3},
		{"foo/foooooo=4", args{"foo", "foooooo"}, 4},
		{"foooooo/foo=4", args{"foooooo", "foo"}, 4},
		{"世界世界世界/foo=6", args{"世界世界世界", "foo"}, 6},
		{"Test #1", args{"apple", "dapple"}, 1},
		{"Test #2", args{"dapple", "apple"}, 1},
		{"Identical", args{"apple", "apple"}, 0},
		{"Add", args{"apple", "apples"}, 1},
		{"Remove", args{"apples", "apple"}, 1},
		{"Substitute", args{"apples", "applet"}, 1},
		{"Test #1", args{"apple", "dapple"}, 1},
		{"Test #2", args{"dapple", "apple"}, 1},
		{"Test #3", args{"apple", "opple"}, 1},
		{"Test #4", args{"kitten", "sitting"}, 3},
		{"Test #5", args{"sunday", "saturday"}, 3},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := LevRecursive(tt.args.s1, tt.args.s2); got != tt.want {
					t.Errorf("LevRecursive() = %v, want %v", got, tt.want)
				}
			},
		)
	}
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
		{"/apple=5", args{"", "apple"}, 5},
		{"apple/=5", args{"apple", ""}, 5},
		{"apple/apple=0", args{"apple", "apple"}, 0},
		{"apple/dapple=1", args{"apple", "dapple"}, 1},
		{
			"alfabet/alfabet-1=1",
			args{"abcdefghijklmnopqrstuvxyzåäö", "abcdefghijklmnopqrstuvxyzåä"}, 1,
		},
		{"long/long=24", args{"sdfpioudfglkjaweoiulknsdasdasd", "cxvpouwejk,mnasdljkhasdlkhasdiu"}, 23},
		{"cattle/apple=1", args{"cattle", "apple"}, 3},
		{"foo/bar=3", args{"foo", "bar"}, 3},
		{"foo/foooooo=4", args{"foo", "foooooo"}, 4},
		{"foooooo/foo=4", args{"foooooo", "foo"}, 4},
		{"世界世界世界/foo=6", args{"世界世界世界", "foo"}, 6},
		{"Test #1", args{"apple", "dapple"}, 1},
		{"Test #2", args{"dapple", "apple"}, 1},
		{"Identical", args{"apple", "apple"}, 0},
		{"Add", args{"apple", "apples"}, 1},
		{"Remove", args{"apples", "apple"}, 1},
		{"Substitute", args{"apples", "applet"}, 1},
		{"Test #1", args{"apple", "dapple"}, 1},
		{"Test #2", args{"dapple", "apple"}, 1},
		{"Test #3", args{"apple", "opple"}, 1},
		{"Test #4", args{"kitten", "sitting"}, 3},
		{"Test #5", args{"sunday", "saturday"}, 3},
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
