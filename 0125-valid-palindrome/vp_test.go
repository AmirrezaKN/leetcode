package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: " ",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.s))
		})
	}
}
