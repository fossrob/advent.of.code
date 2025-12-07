package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(path string) ([]string, error) {
	// 1. Read the entire file content into a byte slice.
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 2. Convert the byte slice to a string.
	fileString := string(content)

	// 3. Split the string by the newline character (\n) to get a slice of lines.
	lines := strings.Split(fileString, "\n")

	return lines, nil
}

func countRolls(grid [][]string, row int, col int) (neighbors int) {

	for r := row - 1; r <= row+1; r++ {
		if r < 0 || r == len(grid) {
			// row is out of bounds
			continue
		}
		for c := col - 1; c <= col+1; c++ {
			if c < 0 || c == len(grid[r]) {
				// column is out of bounds
				continue
			} else if r == row && c == col {
				// this is the roll we are checking so skip
				continue
			} else {
				// check for neighboring roll
				if grid[r][c] == "@" {
					neighbors++
				}
			}
		}
	}
	return neighbors
}

func removeRolls(gridIn [][]string, removedIn int) (removedOut int, gridOut [][]string) {

	// create a new grid as a copy of the old because in Go assigning a slice to another slice still uses the same underlying array!
	// this also means that each of the inner slices / rows are referencing the same arrays even if you use copy() or append() with gridIn
	width, height := len(gridIn[0]), len(gridIn)
	gridOut = make([][]string, height)
	for i := range gridOut {
		gridOut[i] = make([]string, width)
	}

	for row := range len(gridIn) {
		for col := range len(gridIn[row]) {
			// see above, make a copy of the individual item to avoid slices referencing the same array
			gridOut[row][col] = gridIn[row][col]
			if gridIn[row][col] == "@" {
				if countRolls(gridIn, row, col) < 4 {
					gridOut[row][col] = "x"
					removedOut++
				}
			}
		}
		// fmt.Println(gridOut[row])
	}
	// fmt.Println()
	// fmt.Println("Removed:", removedOut, "Total:", removedIn+removedOut)
	// fmt.Println()

	if removedOut == 0 {
		// this pass removed zero rolls, we're done
		return removedIn + removedOut, gridOut
	} else {
		// try again
		return removeRolls(gridOut, removedIn+removedOut)
	}
}

func puzzle1(input []string) int {

	accessibleRolls := 0

	// create a two dimensional grid to hold the input
	width, height := len(input[0]), len(input)
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}

	// populate the grid slice from the input
	for row := range height {
		for col := range width {
			grid[row][col] = input[row][col : col+1]
		}
	}

	// process grid, counting rolls that can be removed
	for row := range height {
		for col := range width {
			if grid[row][col] == "@" {
				if countRolls(grid, row, col) < 4 {
					accessibleRolls++
				}
			}
		}
	}
	return accessibleRolls
}

func puzzle2(input []string) int {
	removedRolls := 0

	// create a two dimensional grid to hold the input
	width, height := len(input[0]), len(input)
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}

	// populate the grid slice from the input
	for row := range height {
		for col := range width {
			grid[row][col] = input[row][col : col+1]
		}
	}

	removedRolls, _ = removeRolls(grid, removedRolls)

	return removedRolls
}

func main() {

	input, err := readInput("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code - Day 04")
	fmt.Println()

	fmt.Println("Number of items:", len(input))
	fmt.Println()

	fmt.Println("Puzzle 1:")
	fmt.Println()

	fmt.Println("Answer:", puzzle1(input))
	fmt.Println()

	fmt.Println("Puzzle 2:")
	fmt.Println()

	fmt.Println("Answer:", puzzle2(input))
	fmt.Println()

}
