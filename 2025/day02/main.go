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

	// 3. Split the string by "," to get a slice of ranges.
	lines := strings.Split(fileString, ",")

	return lines, nil
}

func analyseInput(sequence []string) []int {
	idLengths := []int{}

	for _, idRangeStr := range sequence {

		idRange := strings.Split(idRangeStr, "-")
		startStr, endStr := idRange[0], idRange[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			chars := len(idStr)
			idLengths = append(idLengths, chars)
		}
	}

	slices.Sort(idLengths)

	return slices.Compact(idLengths)

}

func compareString(s string, v int) (repeats bool) {

	// none of these cases provide segments of equal length to compare
	if len(s) < 2 || (len(s)/v < 2) || (len(s)%v > 0) {
		return false
	}

	splitStr := []string{}
	segments := len(s) / v

	// split the string into a slice of segments, of length: v
	pos := 0
	for segment := 1; segment <= segments; segment++ {
		splitStr = append(splitStr, s[pos:(pos+v)])
		pos += v
	}

	// compare each value in the slice returning at the first instance of a mismatch
	for i := 1; i < len(splitStr); i++ {
		if splitStr[i] != splitStr[0] {
			return false
		}
	}

	// if no patterns have been found above then the id is valid
	return true
}

// take a sequence of ranges, for each number in the range look for numbers solely consisting of repeating patterns, sum each of these
// examples: 55 (5 twice), 6464 (64 twice), and 123123 (123 twice)
func puzzle1(sequence []string) (invalidSum int) {

	for _, idRangeStr := range sequence {

		idRange := strings.Split(idRangeStr, "-")
		startStr, endStr := idRange[0], idRange[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			chars := len(idStr)

			// skip ids with an odd number of characters because they can't consist solely of repeating patterns
			if chars%2 > 0 {
				continue
			}
			// split the id into two parts, compare to see if they match
			part1, part2 := idStr[0:chars/2], idStr[chars-(chars/2):]
			if part1 == part2 {
				invalidSum += id
			}

		}
	}

	return invalidSum
}

// invalid ids can now include those made up of any repeating numbers, repeated at least once
// examples: 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times)
// we know from analysing the input, that we have ids ranging in length from 1 - 10 characters: [1 2 3 4 5 6 7 8 9 10]
// 1  - can be excluded obviously, every other number could be made up of the same repeating single digit, plus combinations of 2, 3, 4 or 5 numbers
// 2  - 2|2
// 3  - 3|3|3
// 4  - 4|4|4|4, 44|44
// 5  - 5|5|5|5|5
// 6  - 6|6|6|6|6|6, 66|66|66, 666|666
// 7  - 7|7|7|7|7|7|7
// 8  - 8|8|8|8|8|8|8|8, 88|88|88|88, 8888|8888
// 9  - 9|9|9|9|9|9|9|9|9, 999|999|999
// 10 - 1|1|1|1|1|1|1|1|1|1, 11|11|11|11|11, 11111|11111
func puzzle2(sequence []string) (invalidSum int) {

	invalids := []int{}

	// loop through each of the ranges
	for _, idRangeStr := range sequence {

		idRange := strings.Split(idRangeStr, "-")
		startStr, endStr := idRange[0], idRange[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)

		// loop through each of the ids in the range
		for id := start; id <= end; id++ {
			invalidId := false
			idStr := strconv.Itoa(id)

			// check for patterns of length 1-5
			for v := 1; v <= 5; v++ {
				invalidId = compareString(idStr, v)
				if invalidId {
					invalids = append(invalids, id)
					break
				}
			}
		}
	}

	// sum each of the invalid ids identified
	for _, id := range invalids {
		invalidSum += id
	}

	return invalidSum
}

func main() {

	sequence, err := readInput("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code - Day 02")
	fmt.Println()

	fmt.Println("ID lengths: ", analyseInput(sequence))
	fmt.Println()

	fmt.Println("Puzzle 1:")
	fmt.Println()

	fmt.Println("Answer:", puzzle1(sequence))
	fmt.Println()

	fmt.Println("Puzzle 2:")
	fmt.Println()

	fmt.Println("Answer:", puzzle2(sequence))
	fmt.Println()

}
