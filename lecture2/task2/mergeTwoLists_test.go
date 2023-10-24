package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
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
			name: "2 lists",
			args: args{
				list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
				list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		}, {
			name: "1 list is empty",
			args: args{
				list1: nil,
				list2: &ListNode{Val: 0, Next: nil},
			},
			want: &ListNode{Val: 0, Next: nil},
		},
		{
			name: "both lists are empty",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "one list is longer",
			args: args{
				list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				list2: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		}, {
			name: "duplicate values in both lists",
			args: args{
				list1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}},
				list2: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}},
			},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: nil}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			assert.Equal(t, tt.want, got)
		})
	}
}
