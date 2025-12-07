package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLinesSimple(path string) ([]string, error) {
	// 1. Read the entire file content into a byte slice.
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 2. Convert the byte slice to a string.
	fileString := string(content)

	// 3. Split the string by the newline character (\n) to get a slice of lines.
	lines := strings.Split(fileString, "\n")

	// Note: This method might produce an empty string at the end of the slice
	// if the file ends with a newline. You may need a small check here.

	return lines, nil
}

// take a sequence of inputs in the form Lx / Rx representing the direction and the number of steps to turn the dial on the safe
// the dial has 100 positions from 0 - 99 and starts at 50
// count how many times the dial rests on 0 after a move
func puzzle1(sequence []string) (zeros int) {
	var position = 50

	for _, move := range sequence {
		direction := move[0:1]
		number, _ := strconv.Atoi(move[1:])

		// some numbers of moves are greater than 100 making anything other than the remainder superfluos
		if number > 100 {
			number %= 100
		}

		switch direction {
		case "L":
			// fmt.Println("Move left", number)
			if number > position {
				number -= position
				position = 100 - number
			} else {
				position -= number
			}
		case "R":
			// fmt.Println("Move right", number)
			if number >= (100 - position) {
				number -= (100 - position)
				position = 0 + number
			} else {
				position += number
			}
		}

		if position == 0 {
			zeros++
		}
		// fmt.Println("New positon", position)
	}
	return zeros
}

// take a sequence of inputs in the form Lx / Rx representing the direction and the number of steps to turn the dial on the safe
// the dial has 100 positions from 0 - 99 and starts at 50
// count how many times the dial passes 0, or rests on 0 after a move
func puzzle2(sequence []string) (zeros int) {
	var position = 50

	for _, move := range sequence {
		direction := move[0:1]
		number, _ := strconv.Atoi(move[1:])

		// some numbers of moves are greater than 100 making anything other than the remainder superfluos
		for number >= 100 {
			number -= 100
			zeros++
		}

		// fmt.Print(direction, number, position, zeros)

		switch direction {
		case "L":
			// fmt.Println("Move left", number)
			if number > position {
				// if starting position is 0 then don't count the transition "past" zero
				if position != 0 {
					zeros++
				}
				number -= position
				position = 100 - number
			} else {
				position -= number
			}
		case "R":
			// fmt.Println("Move right", number)
			if number >= (100 - position) {
				number -= (100 - position)
				position = 0 + number
				// if end position is 0 then don't count the transition "past" zero
				if position != 0 {
					zeros++
				}
			} else {
				position += number
			}
		}

		if position == 0 {
			zeros++
		}

		// fmt.Println(" -> ", direction, number, position, zeros)

	}
	return zeros
}

func main() {

	sequence, err := readLinesSimple("puzzle.input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code - Day 01")
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
