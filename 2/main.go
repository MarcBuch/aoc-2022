package main

import "fmt"

// What would your total score be if everything goes exactly according to your strategy guide?
// Can you trust the strategy guide?
// ---
// If opponent chooses
// A for rock
// B for paper
// C for scissors
// ---
// Respond with
// X for rock
// Y for Paper
// Z for scissors
// ---
// Score per shape
// 1 for rock
// 2 for paper
// 3 for scissors
// ---
// Plus score for the outcome of the round
// 0 if lost
// 3 if draw
// 6 if won

func calculateScore(choices []string) int {
	lose := 0
	draw := 3
	win := 6
	score := 0
	opponent := choices[0]
	you := choices[1]
	var youConvert string

	// Get score per shape & convert to same set of strings
	switch you {
	// Rock
	case "X":
		score += 1
		youConvert = "A"
	// Paper
	case "Y":
		score += 2
		youConvert = "B"
	// Scissors
	case "Z":
		score += 3
		youConvert = "C"
	}

	// Calc outcome of round
	if opponent == "A" {
		switch youConvert {
		case "A":
			// Rock + Rock = draw
			score += draw
		case "B":
			// Rock + Paper = win
			score += win
		case "C":
			// Rock + Scissors = lose
			score += lose
		}
	} else if opponent == "B" {
		switch youConvert {
		case "A":
			// Paper + Rock = lose
			score += lose
		case "B":
			// Paper + Paper = draw
			score += draw
		case "C":
			// Paper + Scissors = win
			score += win
		}
	} else if opponent == "C" {
		switch youConvert {
		case "A":
			// Scissors + Rock = win
			score += win
		case "B":
			// Scissors + Paper = lose
			score += lose
		case "C":
			// Scissors + Scissors = draw
			score += draw
		}
	}

	return score
}

func main() {
	inputs := readFile("./input.txt")

	totalScore := 0

	for _, round := range inputs {
		totalScore += calculateScore(round)
	}
	fmt.Println(totalScore)
}
