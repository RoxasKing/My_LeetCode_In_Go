package My_LeetCode_In_Go

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"",
			args{
				[][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}},
			},
			[][]int{{1, 3}, {4, 7}},
		},
		{
			"",
			args{[][]int{{1, 4}, {0, 2}, {3, 5}}},
			[][]int{{0, 5}},
		},
		{
			"",
			args{[][]int{{1, 4}, {1, 5}}},
			[][]int{{1, 5}},
		},
		{
			"",
			args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"",
			args{[][]int{{1, 4}, {4, 5}}},
			[][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
