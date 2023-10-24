package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return MCMXCIV string",
			args: args{1994},
			want: "MCMXCIV",
		},
		{
			name: "should return LVIII string",
			args: args{58},
			want: "LVIII",
		},
		{
			name: "should return III string",
			args: args{3},
			want: "III",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}
