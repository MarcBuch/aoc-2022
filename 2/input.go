package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var inputSlice [][]string

	// Parse the file contents
	reader := bufio.NewReader(f)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		// Check if line is not empty
		if len(line) > 0 {
			// Convert byte array to string
			lineStr := string(line[:])
			inputSlice = append(inputSlice, strings.Split(lineStr, " "))
		}
	}

	return inputSlice
}
