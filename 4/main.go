package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Every section has a unique ID and each elf is assigned a range of sections
// assignments overlap
// list may contain
// TODO: In how many assignment pairs does one range fully contain the other?

func main() {
	fmt.Println(partOne("./input.txt"))
	fmt.Println(partTwo("./input.txt"))
}

func partOne(inputFile string) int {
	input := readFile(inputFile)

	var total int

	for _, pair := range input {
		firstPair := resolve(pair[0])
		secondPair := resolve(pair[1])

		if contains(firstPair, secondPair) {
			total++
		}
	}

	return total
}

func partTwo(inputFile string) int {
	input := readFile(inputFile)

	var totalPairsOverlap int

	for _, pair := range input {
		firstPair := resolve(pair[0])
		secondPair := resolve(pair[1])

		for key := range firstPair {
			_, isPresent := secondPair[key]

			if isPresent {
				// it is only required to know if the pairs overlap, not how often
				totalPairsOverlap++
				break
			}
		}
	}
	return totalPairsOverlap
}

func contains(pairA map[int]struct{}, pairB map[int]struct{}) bool {
	for key := range pairA {
		_, isPresent := pairB[key]

		if isPresent {
			delete(pairA, key)
			delete(pairB, key)
		}
	}

	if len(pairA) == 0 || len(pairB) == 0 {
		return true
	}
	return false
}

func resolve(assignment string) map[int]struct{} {
	sp := strings.Split(assignment, "-")
	start, _ := strconv.Atoi(sp[0])
	end, _ := strconv.Atoi(sp[1])

	set := make(map[int]struct{})

	for i := start; i <= end; i++ {
		set[i] = struct{}{}
	}
	// f, _ := strconv.Atoi(included)
	return set
}

func readFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var inputSlice [][]string

	// Parse the file contents
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		rucksack := strings.Split(sc.Text(), ",")
		inputSlice = append(inputSlice, rucksack)
	}

	return inputSlice
}
