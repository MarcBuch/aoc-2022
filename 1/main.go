package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getHighestCalories(inputFile *os.File) int {
	var inputSlice []int
	var buffer int

	// Parse the file contents
	reader := bufio.NewReader(inputFile)
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
	
	// Get highest int
	var highest, temp int
	for _, element := range inputSlice {
		if element > temp {
			temp = element
			highest = temp
		}
	}
	return highest
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(getHighestCalories(f))
}