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
			name:     "Puzzle 1, Test 1",
			sequence: []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			want:     3,
		},
		{
			name:     "Puzzle 1, Test 2",
			sequence: []string{"L68", "L130", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			want:     3,
		}}
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
			name:     "Puzzle 2, Test 1",
			sequence: []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			want:     6,
		},
		{
			name:     "Puzzle 2, Test 2",
			sequence: []string{"L68", "L30", "R248", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			want:     8,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := puzzle2(tt.sequence)
			if got != tt.want {
				t.Errorf("puzzle1() = %v, want %v", got, tt.want)
			}
		})
	}
}
