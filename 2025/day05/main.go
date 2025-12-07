package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
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

func puzzle1(ranges [][]int, items []int) (freshItems int) {

	for _, item := range items {
		for _, r := range ranges {
			if r[0] <= item && item <= r[1] {
				freshItems++
				break
			}
		}
	}
	return freshItems
}

func puzzle2(ranges [][]int) (freshItems int) {

	freshSlice := []int{}

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			freshSlice = append(freshSlice, i)
		}
	}

	slices.Sort(freshSlice)
	freshItems = len(slices.Compact(freshSlice))

	return freshItems
}

func main() {

	input, err := readInput("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ranges, items := [][]int{}, []int{}

	for _, line := range input {
		if strings.Contains(line, "-") {
			splitStr := strings.Split(line, "-")
			startInt, _ := strconv.Atoi(splitStr[0])
			endInt, _ := strconv.Atoi(splitStr[1])
			ranges = append(ranges, []int{startInt, endInt})
		} else {
			lineInt, _ := strconv.Atoi(line)
			items = append(items, lineInt)
		}
	}

	fmt.Println("Advent of Code - Day 05")
	fmt.Println()

	fmt.Println("Number of ranges, items:", len(ranges), len(items))
	fmt.Println()

	fmt.Println("Puzzle 1:")
	fmt.Println()

	fmt.Println("Answer:", puzzle1(ranges, items))
	fmt.Println()

	fmt.Println("Puzzle 2:")
	fmt.Println()

	fmt.Println("Answer:", puzzle2(ranges))
	fmt.Println()

}
