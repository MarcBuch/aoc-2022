package main

import "testing"

func runTestPartOne(t *testing.T, input []string, want int) {
	got := partOne(input)
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}

func TestPartOne(t *testing.T) {
	runTestPartOne(t, []string{"C", "Y"}, 2)
	runTestPartOne(t, []string{"B", "Y"}, 5)
	runTestPartOne(t, []string{"A", "Z"}, 3)
	runTestPartOne(t, []string{"A", "X"}, 4)
}

func runTestPartTwo(t *testing.T, input []string, want int) {
	got := partTwo(input)
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}

func TestPartTwo(t *testing.T) {
	// Opponent plays Scissors and guide says draw, 3 + 3 = 6
	runTestPartTwo(t, []string{"C", "Y"}, 6)

	// Opponent plays rock and guide says lose, 3 + 0 = 3
	runTestPartTwo(t, []string{"A", "X"}, 3)
}
