package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
    [H]         [D]     [P]
[W] [B]         [C] [Z] [D]
[T] [J]     [T] [J] [D] [J]
[H] [Z]     [H] [H] [W] [S]     [M]
[P] [F] [R] [P] [Z] [F] [W]     [F]
[J] [V] [T] [N] [F] [G] [Z] [S] [S]
[C] [R] [P] [S] [V] [M] [V] [D] [Z]
[F] [G] [H] [Z] [N] [P] [M] [N] [D]
 1   2   3   4   5   6   7   8   9
*/

// Puzzle input contains instructions
// Crates are moved one at a time
// After the rearrangement procedure completes, what crate ends up on top of each stack?

func main() {
	input := readFile("./input.txt")
	fmt.Printf("Result Part One: %s\n", partOne(input))
	fmt.Printf("Result Part Two: %s\n", partTwo(input))
}

func example() string {
	input := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	crates := []Stack{
		{},
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}

	op, err := NewOperator(input, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := Warehouse{
		Operator: op,
		Cargo:    crates,
	}

	w.OrganizeCrates()

	return w.GetCratesOnTop()
}

func partOne(input []string) string {
	crates := []Stack{
		{"F", "C", "J", "P", "H", "T", "W"},
		{"G", "R", "V", "F", "Z", "J", "B", "H"},
		{"H", "P", "T", "R"},
		{"Z", "S", "N", "P", "H", "T"},
		{"N", "V", "F", "Z", "H", "J", "C", "D"},
		{"P", "M", "G", "F", "W", "D", "Z"},
		{"M", "V", "Z", "W", "S", "J", "D", "P"},
		{"N", "D", "S"},
		{"D", "Z", "S", "F", "M"},
	}

	op, err := NewOperator(input, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := Warehouse{
		Operator: op,
		Cargo:    crates,
	}

	w.OrganizeCrates()

	return w.GetCratesOnTop()
}

func partTwo(input []string) string {
	crates := []Stack{
		{"F", "C", "J", "P", "H", "T", "W"},
		{"G", "R", "V", "F", "Z", "J", "B", "H"},
		{"H", "P", "T", "R"},
		{"Z", "S", "N", "P", "H", "T"},
		{"N", "V", "F", "Z", "H", "J", "C", "D"},
		{"P", "M", "G", "F", "W", "D", "Z"},
		{"M", "V", "Z", "W", "S", "J", "D", "P"},
		{"N", "D", "S"},
		{"D", "Z", "S", "F", "M"},
	}

	op, err := NewOperator(input, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := Warehouse{
		Operator: op,
		Cargo:    crates,
	}

	w.OrganizeCrates()

	return w.GetCratesOnTop()
}

type Warehouse struct {
	Operator Operator
	Cargo    []Stack
}

func (w *Warehouse) OrganizeCrates() {
	for _, v := range w.Operator.Instructions {
		// Decrementing for slices starting at 0
		v.From--
		v.To--
		w.Operator.Rearrange(&w.Cargo[v.From], &w.Cargo[v.To], v.Count)
	}
}

func (w *Warehouse) GetCratesOnTop() string {
	var str string
	for _, v := range w.Cargo {
		if len(v) == 0 {
			continue
		}
		l := len(v) - 1
		str += v[l]
	}
	return str
}

type Instructions struct {
	Count int
	From  int
	To    int
}

type Operator struct {
	Instructions []Instructions
	Stacking     bool
}

func (op *Operator) Rearrange(s, t *Stack, n int) {
	c := s.Pop(n)
	// Reverse is only needed if crates can be moved one by one.
	if op.Stacking == false {
		c.Reverse()
	}
	t.Push(c)
}

func NewOperator(input []string, stacking bool) (Operator, error) {
	operator := Operator{
		[]Instructions{},
		stacking,
	}

	for _, v := range input {
		inst, err := processLine(v)
		if err != nil {
			return Operator{}, err
		}
		operator.Instructions = append(operator.Instructions, inst)
	}

	return operator, nil
}

// processLine parses a line and returns instructions
func processLine(l string) (Instructions, error) {
	var s string
	var a, b, c int
	_, err := fmt.Sscanf(l, "%s %d %s %d %s %d", &s, &a, &s, &b, &s, &c)
	if err != nil {
		return Instructions{}, fmt.Errorf("invalid input\n")
	}
	return Instructions{Count: a, From: b, To: c}, nil
}

func readFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var inputSlice []string

	// Parse the file contents
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		inputSlice = append(inputSlice, sc.Text())
	}

	return inputSlice
}
