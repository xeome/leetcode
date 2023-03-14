package main

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root:      createTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 0, 1}),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root:      createTree([]int{1, 2, 3}),
				targetSum: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
