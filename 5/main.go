package main

import "fmt"

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

var cargo = map[int][]string{
	1: {"F", "C", "J", "P", "H", "T", "W"},
	2: {"G", "R", "V", "F", "Z", "J", "B", "H"},
	3: {"H", "P", "T", "R"},
	4: {"Z", "S", "N", "P", "H", "T"},
	5: {"N", "V", "F", "Z", "H", "J", "C", "D"},
	6: {"P", "M", "G", "F", "W", "D", "Z"},
	7: {"M", "V", "Z", "W", "S", "J", "D", "P"},
	8: {"N", "D", "S"},
	9: {"D", "Z", "S", "F", "M"},
}

func main() {
	partOne()
}

func partOne() {
	// TODO: Read input instructions
	fmt.Println(cargo)
	// Test instructions
	// move 1 from 8 to 2
	from := cargo[8]
	to := cargo[2]
	i := 2
	from, s := deletetItem(from, i)
	fmt.Printf("Moving %s from cargo[%d]\n", s, 8)

	// Write changes back
	cargo[8] = from
	to = append(to, s)
	cargo[2] = to
	fmt.Println(cargo[8])
	fmt.Println(cargo[2])
	// Get item from position
	//
	//	for i := 1; i <= amount; i++ {
	//		item := cargo[from][len(cargo[from])-i:]
	//		fmt.Println(item)
	//	}
}

func deletetItem(slice []string, i int) ([]string, string) {
	item := slice[i]
	slice[i] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]

	return slice, item
}
