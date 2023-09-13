package Homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		array []int
	}
	//tests := []struct {
	//	name string
	//	args args
	//	want []int
	//}{
	//	{
	//		name: "array should be sorted",
	//		args: args{
	//			boxes: struct {
	//				array []int
	//			}{
	//				array: []int{9, 3, 4, 2, 1, 8},
	//			},
	//		},
	//		want: []int{1, 2, 3, 4, 8, 9},
	//	},
	//	{
	//		name: "empty array should result in an empty sorted array",
	//		args: args{
	//			boxes: struct {
	//				array []int
	//			}{
	//				array: []int{},
	//			},
	//		},
	//		want: []int{},
	//	},
	//	{
	//		name: "already sorted array should remain the same",
	//		args: args{
	//			boxes: struct {
	//				array []int
	//			}{
	//				array: []int{1, 2, 3, 4, 8, 9},
	//			},
	//		},
	//		want: []int{1, 2, 3, 4, 8, 9},
	//	},
	//}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "array should be sorted",
			args: args{[]int{9, 3, 4, 2, 1, 8}},
			want: []int{1, 2, 3, 4, 8, 9},
		},
		{
			name: "empty array should result in an empty sorted array",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "already sorted array should remain the same",
			args: args{[]int{1, 2, 3, 4, 8, 9}},
			want: []int{1, 2, 3, 4, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoxes(tt.args.array)
			b.QuickSort()

			assert.Equal(t, tt.want, b.array)
		})
	}
}
