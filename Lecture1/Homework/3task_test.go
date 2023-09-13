package Homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareWithSort(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{[]int{5, 4, 2, 1, 4}, []int{1, 4, 5, 2, 4}},
			want: true,
		},
		{
			name: "should return false",
			args: args{[]int{7, 5, 3, 2, 1, 6, 8, 1}, []int{1, 3, 2, 1, 6, 9, 6, 5}},
			want: false,
		},
		{
			name: "should return true",
			args: args{[]int{1, 1, 2, 2, 3, 4, 5}, []int{5, 3, 4, 2, 2, 1, 1}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareWithSort(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)

			got = compareWithoutSort(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
