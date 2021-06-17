package main

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{1, 7, 11}, []int{2, 4, 6}, 3}, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{"2", args{[]int{1, 1, 2}, []int{1, 2, 3}, 2}, [][]int{{1, 1}, {1, 1}}},
		{"3", args{[]int{1, 2}, []int{3}, 3}, [][]int{{1, 3}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
