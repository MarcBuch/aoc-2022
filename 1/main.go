package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func readFile(path string) []int {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var inputSlice []int
	var buffer int

	// Parse the file contents
	reader := bufio.NewReader(f)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		// Check if line is not empty
		if len(line) > 0 {
			// Convert byte array to string
			lineStr := string(line[:])
			// Convert string to int
			lineInt, _ := strconv.Atoi(lineStr)
			buffer += lineInt
			// If line is empty (empty line has been declared as delimiter)
		} else {
			// append the buffer int to the inputSlice and clear the buffer
			inputSlice = append(inputSlice, buffer)
			buffer = 0
		}
	}
	return inputSlice
}

func getHighestCalories(input []int) int {
	// Get highest int
	var highest, temp int
	for _, element := range input {
		if element > temp {
			temp = element
			highest = temp
		}
	}
	return highest
}

func topThreeCaloriesInTotal(input []int) int {
	// Sort the slice to descending
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	return input[0] + input[1] + input[2]
}

func main() {

	input := readFile("./input.txt")

	fmt.Printf("The elf with the highest calories has: %d calories in his bag.\n", getHighestCalories(input))
	fmt.Println("--------")
	fmt.Printf("The top three elfs are carying %d calories in total.\n", topThreeCaloriesInTotal(input))
}
