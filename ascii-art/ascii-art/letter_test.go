package main

import (
	"slices"
	"testing"
)

func TestFindThe(t *testing.T) {
	type args struct {
		fn     func([]int) int
		letter Letter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{fn: slices.Max[[]int], letter: Letter{"a", "b  ", "c"}}, 3},
		{"Test 2", args{fn: slices.Max[[]int], letter: Letter{"a      ", "b  ", "c"}}, 7},
		{"Test 3", args{fn: slices.Max[[]int], letter: Letter{"a", "b  ", "c           "}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindThe(tt.args.fn, tt.args.letter); got != tt.want {
				t.Errorf("FindThe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSpaces(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{num: 0}, ""},
		{"Test 2", args{num: 1}, " "},
		{"Test 3", args{num: 2}, "  "},
		{"Test 4", args{num: 3}, "   "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSpaces(tt.args.num); got != tt.want {
				t.Errorf("GenerateSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
