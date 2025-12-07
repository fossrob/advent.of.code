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

func puzzle1() int {
	return 0
}

func puzzle2() int {
	return 0
}

func main() {

	input, err := readInput("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ranges, items := []string{}, []string{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			ranges = append(ranges, line)
		} else {
			items = append(items, line)
		}
	}

	fmt.Println("Advent of Code - Day 05")
	fmt.Println()

	fmt.Println("Number of ranges, items:", len(ranges), len(items))
	fmt.Println()

	fmt.Println("Puzzle 1:")
	fmt.Println()

	fmt.Println("Answer:", puzzle1())
	fmt.Println()

	fmt.Println("Puzzle 2:")
	fmt.Println()

	fmt.Println("Answer:", puzzle2())
	fmt.Println()

}
