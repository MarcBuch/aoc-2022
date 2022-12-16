package main

type Stack []string

// Pop removes n elements from the right of the Stack and returns a new Stack
func (s *Stack) Pop(n int) Stack {
	r := *s
	value := r[len(r)-n:]
	*s = r[:len(r)-n]
	return value
}

// Push appends crates to the Stack
func (s *Stack) Push(c Stack) {
	r := *s
	r = append(r, c...)
	*s = r
}

// Reverse changes the order of a stack
// needed for moving one crate at a time
func (s *Stack) Reverse() {
	r := *s
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
