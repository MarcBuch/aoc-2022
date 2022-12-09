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
// ---
// Part Two
// ---
// X lose
// Y draw
// Z win

// calcScore takes in the two choices players made (Rock + Paper) and returns the score for the player.
// - A is the opponents choice
// - B is the players choice
func calcScore(a string, b string) int {
	lose := 0
	draw := 3
	win := 6

	score := 0
	// Get score per shape & convert to same set of strings
	switch b {
	// Rock
	case "A":
		score += 1
	// Paper
	case "B":
		score += 2
	// Scissors
	case "C":
		score += 3
	}

	// Calc outcome of round
	if a == "A" {
		switch b {
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
	} else if a == "B" {
		switch b {
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
	} else if a == "C" {
		switch b {
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

func partOne(choices []string) int {
	opponent := choices[0]
	you := choices[1]
	var youConvert string

	// Get score per shape & convert to same set of strings
	switch you {
	// Rock
	case "X":
		youConvert = "A"
	// Paper
	case "Y":
		youConvert = "B"
	// Scissors
	case "Z":
		youConvert = "C"
	}

	return calcScore(opponent, youConvert)
}

func partTwo(choices []string) int {
	opponent := choices[0]
	plannedMove := choices[1]

	// Debug purpose
	switch opponent {
	case "A":
		fmt.Print("Opponent plays rock, ")
	case "B":
		fmt.Print("Opponent plays paper, ")
	case "C":
		fmt.Print("Opponent plays scissor, ")
	}
	switch plannedMove {
	case "X":
		fmt.Print("guide says lose, ")
	case "Y":
		fmt.Print("guide says draw, ")
	case "Z":
		fmt.Print("guide says win, ")
	}

	lose := "X"
	win := "Z"
	draw := "Y"
	dic := map[string]map[string]string{
		"A": {
			win:  "B",
			draw: "A",
			lose: "C",
		},
		"B": {
			win:  "C",
			draw: "B",
			lose: "A",
		},
		"C": {
			win:  "A",
			draw: "C",
			lose: "B",
		},
	}

	// Example, opponent plays rock (A), guide says win (Z)
	actualMove := dic[opponent][plannedMove]

	// Debug
	switch actualMove {
	case "A":
		fmt.Print("playing rock\n")
	case "B":
		fmt.Print("playing paper\n")
	case "C":
		fmt.Print("playing scissor\n")
	}

	return calcScore(opponent, actualMove)
}

func main() {
	inputs := readFile("./input.txt")

	totalScore := 0

	for _, round := range inputs {
		totalScore += partTwo(round)
	}
	fmt.Println(totalScore)
}
