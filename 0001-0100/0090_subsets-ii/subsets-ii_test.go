package main

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{1, 2, 2},
			},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{2, 2},
				{1, 2, 2},
			},
		},
		{
			"2",
			args{
				[]int{1, 2, 2, 2},
			},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{2, 2},
				{1, 2, 2},
				{2, 2, 2},
				{1, 2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
