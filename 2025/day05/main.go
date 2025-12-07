package main

import (
	"fmt"
	"os"
	"sort"
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

	// sort the ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	prevSlice := []int{0, 0}
	// check for overlaps and add the range for subsequent ranges
	for key, rangeSlice := range ranges {
		// fmt.Print(rangeSlice)
		if rangeSlice[1] <= prevSlice[1] {
			// this entire range exists inside the previous, skip it
			// fmt.Print(" ", rangeSlice[1], " <= ", prevSlice[1], " skip")
			continue
		} else if rangeSlice[0] <= prevSlice[1] {
			// range overlaps, new range starts after old range ends
			freshItems += rangeSlice[1] - prevSlice[1]
			// fmt.Print(" ", prevSlice[1]+1, " - ", rangeSlice[1], " overlap (", rangeSlice[1]-prevSlice[1], ")")
		} else {
			// new range
			freshItems += (rangeSlice[1] - rangeSlice[0]) + 1
			// fmt.Print(" ", rangeSlice[0], " - ", rangeSlice[1], " new (", (rangeSlice[1]-rangeSlice[0])+1, ")")
		}
		prevSlice = ranges[key]
		// fmt.Println()
	}

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
