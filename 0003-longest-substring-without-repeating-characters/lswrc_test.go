package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.args.s))
		})
	}
}
