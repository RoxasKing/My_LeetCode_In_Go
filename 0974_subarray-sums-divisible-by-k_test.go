package leetcode

import (
	"testing"
)

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraysDivByK2(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK2(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK2() = %v, want %v", got, tt.want)
			}
		})
	}
}