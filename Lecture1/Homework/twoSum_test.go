package Homework

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//type of data of arguments
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return two different indexes",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{0, 1},
		},
		{
			name: "should return two same indexes",
			args: args{[]int{3, 3}, 6},
			want: []int{0, 1},
		},
		{
			name: "should return empty array",
			args: args{[]int{2, 7, 11, 15, 25, 65, 765}, 3},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// calling function
			got := twoSumSlower(tt.args.nums, tt.args.target)

			sort.Ints(tt.want)
			sort.Ints(got)

			fmt.Println("The result is for Slower :", got)
			assert.Equal(t, tt.want, got)

			got = twoSumFaster(tt.args.nums, tt.args.target)

			sort.Ints(got)

			fmt.Println("The result is for Faster :", got)
			assert.Equal(t, tt.want, got)
		})
	}
}
