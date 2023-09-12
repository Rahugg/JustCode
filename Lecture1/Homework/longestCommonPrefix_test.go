package Homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return a common prefix with size of 2",
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "should return an empty string",
			args: args{[]string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "should return a common prefix size of 1",
			args: args{[]string{"a"}},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.args.strs)
			//fmt.Println("The results are:", got)
			assert.Equal(t, tt.want, got)
		})
	}
}
