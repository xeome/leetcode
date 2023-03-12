package main

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{nums: []int{-10, -3, 0, 5, 9}},
			want: createTree([]int{0, -3, 9, -10, -1, 5}),
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 3}},
			want: createTree([]int{3, 1}),
		},
		{
			name: "Example 3",
			args: args{nums: []int{}},
			want: createTree([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
