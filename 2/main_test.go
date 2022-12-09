package main

import "testing"

func runTest(t *testing.T, input []string, want int) {
	got := calculateScore(input)
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}

func TestCalculateScore(t *testing.T) {
	runTest(t, []string{"C", "Y"}, 2)
	runTest(t, []string{"B", "Y"}, 5)
	runTest(t, []string{"A", "Z"}, 3)
	runTest(t, []string{"A", "X"}, 4)
}
