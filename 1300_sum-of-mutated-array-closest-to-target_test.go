package leetcode

import (
	"testing"
)

func Test_findBestValue(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 9, 3}, 10}, 3},
		{"2", args{[]int{2, 3, 5}, 10}, 5},
		{"3", args{[]int{60864, 25176, 27249, 21296, 20204}, 56803}, 11361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestValue(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("findBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findBestValue2(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 9, 3}, 10}, 3},
		{"2", args{[]int{2, 3, 5}, 10}, 5},
		{"3", args{[]int{60864, 25176, 27249, 21296, 20204}, 56803}, 11361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestValue2(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("findBestValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}
