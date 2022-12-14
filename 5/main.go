package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Ship struct {
	cargo []string
}

func NewShip(slice []string) *Ship {
	return &Ship{
		cargo: slice,
	}
}

func (s *Ship) pop() string {
	// pop removes the last item of the ship's cargo
	index := len(s.cargo) - 1
	item := s.cargo[index]
	s.cargo = append(s.cargo[:index], s.cargo[index+1:]...)

	return item
}

func (s *Ship) add(str string) {
	s.cargo = append(s.cargo, str)
}

var shipsexampleShips = map[int]*Ship{
	1: NewShip([]string{"Z", "N"}),
	2: NewShip([]string{"M", "C", "D"}),
	3: NewShip([]string{"P"}),
}

var ships = map[int]*Ship{
	1: NewShip([]string{"F", "C", "J", "P", "H", "T", "W"}),
	2: NewShip([]string{"G", "R", "V", "F", "Z", "J", "B", "H"}),
	3: NewShip([]string{"H", "P", "T", "R"}),
	4: NewShip([]string{"Z", "S", "N", "P", "H", "T"}),
	5: NewShip([]string{"N", "V", "F", "Z", "H", "J", "C", "D"}),
	6: NewShip([]string{"P", "M", "G", "F", "W", "D", "Z"}),
	7: NewShip([]string{"M", "V", "Z", "W", "S", "J", "D", "P"}),
	8: NewShip([]string{"N", "D", "S"}),
	9: NewShip([]string{"D", "Z", "S", "F", "M"}),
}

func main() {
	partOne()

}

func partOne() {
	input := readFile("./input.txt")

	for _, instruction := range input {
		sp := strings.Split(instruction, " ")

		move, _ := strconv.Atoi(sp[1])
		from, _ := strconv.Atoi(sp[3])
		to, _ := strconv.Atoi(sp[5])

		for i := 1; i <= move; i++ {
			item := ships[from].pop()
			ships[to].add(item)
		}
	}

	fmt.Println("The last items for each ship are:")
	for i := 1; i <= len(ships); i++ {
		last := ships[i].cargo
		fmt.Printf("ship[%d]: %v\n", i, last[len(last)-1:])
	}
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
