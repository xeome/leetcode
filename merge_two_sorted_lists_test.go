package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				list1: createLinkedList([]int{1, 2, 4}),
				list2: createLinkedList([]int{1, 3, 4}),
			},
			want: createLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "Example 2",
			args: args{
				list1: createLinkedList([]int{}),
				list2: createLinkedList([]int{}),
			},
			want: createLinkedList([]int{}),
		},
		{
			name: "Example 3",
			args: args{
				list1: createLinkedList([]int{}),
				list2: createLinkedList([]int{0}),
			},
			want: createLinkedList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
