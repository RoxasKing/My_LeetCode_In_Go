package main

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"1", args{"abcd", "abcde"}, 'e'},
		{"2", args{"", "y"}, 'y'},
		{"3", args{"a", "aa"}, 'a'},
		{"4", args{"ae", "aea"}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
