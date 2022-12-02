package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "ab",
				t: "a",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				s: "a",
				t: "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAnagram(tt.args.s, tt.args.t))
		})
	}
}
