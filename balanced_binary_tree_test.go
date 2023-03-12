package main

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{root: createTree([]int{3, 9, 20, -1, -1, 15, 7})},
			want: true,
		},
		{
			name: "Example 2",
			args: args{root: createTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})},
			want: false,
		},
		{
			name: "Example 3",
			args: args{root: createTree([]int{})},
			want: true,
		},
		{
			name: "Example 4",
			args: args{root: createTree([]int{1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
