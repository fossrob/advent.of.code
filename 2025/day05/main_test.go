package main

import "testing"

func Test_puzzle1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ranges [][]int
		items  []int
		want   int
	}{
		{
			name: "Test Puzzle 1",
			ranges: [][]int{
				{3, 5},
				{10, 14},
				{16, 20},
				{12, 18},
			},
			items: []int{
				1,
				5,
				8,
				11,
				17,
				32,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle1(tt.ranges, tt.items)
			if got != tt.want {
				t.Errorf("puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_puzzle2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ranges [][]int
		want   int
	}{
		{
			name: "Test Puzzle 2",
			ranges: [][]int{
				{3, 5},
				{10, 14},
				{16, 20},
				{12, 18},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle2(tt.ranges)
			if got != tt.want {
				t.Errorf("puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
