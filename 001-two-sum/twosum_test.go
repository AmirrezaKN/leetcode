package twosum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		wants [][]int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			wants: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			wants: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			wants: [][]int{
				{0, 1},
				{1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, want := range tt.wants {
				if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, want) {
					t.Logf("twoSum() = %v, want %v", got, want)
					continue
				}
			}
		})
	}
}
