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

func puzzle1(sequence []string) int {

	joltage := 0
	for _, bank := range sequence {
		// create a slice of integers representing each battery
		batteries := make([]int, len(bank))
		for b := 0; b < len(bank); b++ {
			battery, _ := strconv.Atoi(bank[b : b+1])
			batteries[b] = battery
		}
		// fmt.Print(batteries)

		pair := make([]int, 2)
		for position, battery := range batteries {
			if battery == slices.Max(batteries) {
				if position < len(batteries)-1 {
					pair[0] = battery
					secondSlice := batteries[position+1:]
					// fmt.Print(" ", position, " ", secondSlice)
					pair[1] = slices.Max(secondSlice)
				} else {
					pair[1] = battery
					secondSlice := batteries[0 : len(batteries)-1]
					// fmt.Print(" ", position, " ", secondSlice)
					pair[0] = slices.Max(secondSlice)
				}
				break
			}
		}

		// fmt.Println(" ", pair)
		pairStr := strconv.Itoa(pair[0]) + strconv.Itoa(pair[1])
		pairInt, _ := strconv.Atoi(pairStr)
		joltage += pairInt

	}
	return joltage
}

func puzzle2(sequence []string) int {

	joltage := 0

	for _, bank := range sequence {
		// create a slice of integers representing each battery
		batteries := make([]int, len(bank))
		for b := 0; b < len(bank); b++ {
			battery, _ := strconv.Atoi(bank[b : b+1])
			batteries[b] = battery
		}

		b, c, batSlice := 0, 0, make([]int, 12)
		for i := 12; i > 0; i-- {
			// find the best combination of 12 batteries, b and c are the start and end of the range to search for the max value in each loop
			c = len(batteries) - i + 1 // always exclude the last N batteries, i.e. the last 11 for the first loop, 10 for the second etc
			for b <= c {               // c is adjusted after each loop but b is not and carries on from it's last position
				bat := batteries[b]
				if bat == slices.Max(batteries[b:c]) {
					batSlice[12-i] = bat
					b++
					break
				}
				b++
			}
		}

		batStr := ""
		for _, bat := range batSlice {
			batStr += strconv.Itoa(bat)
		}
		batInt, _ := strconv.Atoi(batStr)
		joltage += batInt
	}

	return joltage
}

func main() {

	sequence, err := readInput("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code - Day 03")
	fmt.Println()

	fmt.Println("Number of battery banks:", len(sequence))
	fmt.Println()

	fmt.Println("Puzzle 1:")
	fmt.Println()

	fmt.Println("Total output joltage:", puzzle1(sequence))
	fmt.Println()

	fmt.Println("Puzzle 2:")
	fmt.Println()

	fmt.Println("Total output joltage:", puzzle2(sequence))
	fmt.Println()

}
