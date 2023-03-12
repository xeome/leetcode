package main

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{root: createTree([]int{3, 9, 20, -1, -1, 15, 7})},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{root: createTree([]int{1, -1, 2})},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{root: createTree([]int{})},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{root: createTree([]int{0})},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
