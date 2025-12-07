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
			name:     "Puzzle1, Test 1",
			sequence: []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"},
			want:     1227775554,
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
			name:     "Puzzle2, Test 1",
			sequence: []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"},
			want:     4174379265,
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

func Test_compareString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		v    int
		want bool
	}{
		{
			name: "Test compareString length: 0",
			s:    "",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 1",
			s:    "1",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 2, pattern, value 2",
			s:    "11",
			v:    2,
			want: false,
		},
		{
			name: "Test compareString length: 2, no pattern",
			s:    "12",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 3, no pattern",
			s:    "123",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 4, no pattern",
			s:    "1234",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 5, no pattern",
			s:    "12345",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 6, no pattern",
			s:    "123456",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 7, no pattern",
			s:    "1234567",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 8, no pattern",
			s:    "12345678",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 9, no pattern",
			s:    "123456789",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 10, no pattern",
			s:    "1234567890",
			v:    1,
			want: false,
		},
		{
			name: "Test compareString length: 2, pattern, value 1",
			s:    "11",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 3, pattern, value 1",
			s:    "111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 4, pattern, value 1",
			s:    "1111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 4, pattern, value 2",
			s:    "1010",
			v:    2,
			want: true,
		},
		{
			name: "Test compareString length: 5, pattern, value 1",
			s:    "11111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 6, pattern, value 1",
			s:    "111111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 6, pattern, value 2",
			s:    "565656",
			v:    2,
			want: true,
		},
		{
			name: "Test compareString length: 6, pattern, value 3",
			s:    "446446",
			v:    3,
			want: true,
		},
		{
			name: "Test compareString length: 7, pattern, value 1",
			s:    "1111111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 8, pattern, value 1",
			s:    "11111111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 8, pattern, value 2",
			s:    "35353535",
			v:    2,
			want: true,
		},
		{
			name: "Test compareString length: 8, pattern, value 4",
			s:    "38593859",
			v:    4,
			want: true,
		},
		{
			name: "Test compareString length: 9, pattern, value 1",
			s:    "111111111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 9, pattern, value 3",
			s:    "824824824",
			v:    3,
			want: true,
		},
		{
			name: "Test compareString length: 10, pattern, value 1",
			s:    "11111111",
			v:    1,
			want: true,
		},
		{
			name: "Test compareString length: 10, pattern, value 2",
			s:    "2121212121",
			v:    2,
			want: true,
		},
		{
			name: "Test compareString length: 10, pattern, value 5",
			s:    "1188511885",
			v:    5,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareString(tt.s, tt.v)
			if got != tt.want {
				t.Errorf("compareString() = %v, want %v", got, tt.want)
			}
		})
	}
}
