package main

import "testing"

func Test_puzzle1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sequence []string
		want     int
	}{
		{
			name: "Test Puzzle 1",
			sequence: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle1(tt.sequence)
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
		sequence []string
		want     int
	}{
		{
			name: "Test Puzzle 2",
			sequence: []string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			want: 3121910778619,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle2(tt.sequence)
			if got != tt.want {
				t.Errorf("puzzle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
