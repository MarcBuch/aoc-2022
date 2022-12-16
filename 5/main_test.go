package main

import (
	"testing"
)

func TestProcessLine(t *testing.T) {
	want := Instructions{Count: 1, From: 2, To: 3}
	got, _ := processLine("move 1 from 2 to 3")

	if got != want {
		t.Errorf("got %v want %v", got, want)
		t.Fail()
	}
}

func TestStack(t *testing.T) {
	stack := Stack{"A", "B", "C"}

	t.Log("Testing stack.Pop")
	want := 1
	got := stack.Pop(want)

	// Pop should return a single element
	if len(got) != want {
		t.Errorf("got %d want %d", len(got), want)
		t.Fail()
	}
	// And the returned element should be a C
	if got[0] != "C" {
		t.Errorf("got %s, want C", got)
		t.Fail()
	}

	t.Log("Testing stack.Push")
	stack.Push(Stack{"C"})

	// Push should add an element
	if len(stack) != 3 {
		t.Errorf("got %d want 3", len(stack))
		t.Fail()
	}
	// And the last item should be a C
	if stack[len(stack)-1] != "C" {
		t.Errorf("got %s want C", stack[len(stack)-1])
		t.Fail()
	}
}
