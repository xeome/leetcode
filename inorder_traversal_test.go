package main

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{root: createTree([]int{1, -1, 2, 3})},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			args: args{root: createTree([]int{})},
			want: []int{},
		},
		{
			name: "Example 3",
			args: args{root: createTree([]int{1})},
			want: []int{1},
		},
		{
			name: "Example 4",
			args: args{root: createTree([]int{1, 2})},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
