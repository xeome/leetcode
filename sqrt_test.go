package main

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 1},
		{"4", args{4}, 2},
		{"5", args{5}, 2},
		{"6", args{6}, 2},
		{"7", args{7}, 2},
		{"8", args{8}, 2},
		{"9", args{9}, 3},
		{"10", args{10}, 3},
		{"11", args{11}, 3},
		{"12", args{12}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
