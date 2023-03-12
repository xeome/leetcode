package main

import "testing"

func Test_minDepth(t *testing.T) {
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
			want: 2,
		},
		{
			name: "Example 2",
			args: args{root: createTree([]int{2, -1, 3, -1, 4, -1, 5, -1, 6})},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
