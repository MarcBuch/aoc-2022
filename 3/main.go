package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var inputSlice []string

	// Parse the file contents
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		rucksack := sc.Text()
		inputSlice = append(inputSlice, rucksack)
	}

	return inputSlice
}

func priority(item rune) int {
	// ASCII defines 128 characters
	// lowercase items have priorities 1 - 26, from example: (p) 112 - x = 16, where x is 96
	// uppercase items have priorities 27 - 52, from example: (L) 76 - x = 38, where x is 38

	// Uppercase ASCII items end at 90
	if int(item) > 90 {
		return int(item) - 96
	}
	return int(item) - 38
}

func partOne(inputPath string) int {
	input := readFile(inputPath)

	partOneTotal := 0
	for _, rucksack := range input {
		// Split rucksack into two compartments
		firstHalf := rucksack[:len(rucksack)/2]
		secondHalf := rucksack[len(rucksack)/2:]

		for _, item := range firstHalf {
			// if item is in both compartments
			if strings.Contains(secondHalf, string(item)) {
				// figure out its priority
				partOneTotal += priority(item)
				break
			}
		}
	}

	return partOneTotal
}

func partTwo(inputPath string) int {
	// Figure out which item type is common between all 3 elves
	// every set of three lines (input) correspond to a single group
	f, err := os.Open(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	var input [][]string

	// Parse the file contents
	sc := bufio.NewScanner(f)
	var buffer []string
	for sc.Scan() {
		rucksack := sc.Text()

		// Split elfes into groups of 3
		buffer = append(buffer, string(rucksack))
		if len(buffer) == 3 {
			input = append(input, buffer)
			buffer = []string{}
		}
	}

	// Figre out which item is common between all 3 elves
	var total int
	for _, rucksack := range input {
		firstRucksack := rucksack[0]
		secondRucksack := rucksack[1]
		thirdRucksack := rucksack[2]

		var badge rune
		for _, item := range firstRucksack {
			if strings.Contains(secondRucksack, string(item)) && strings.Contains(thirdRucksack, string(item)) {
				badge = item
				break
			}
		}
		total += priority(badge)
	}
	return total
}

func main() {
	fmt.Printf("Part One: %d\n", partOne("./input.txt"))
	fmt.Printf("Part Two: %d", partTwo("./input.txt"))
}
