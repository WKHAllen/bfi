package main

import (
	"fmt"
)

// Stack : Stack object
type Stack []int

// NewStack : Create a new stack object
func NewStack() *Stack {
	s := &Stack{}
	return s
}

// Push : Push to a stack
func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

// Pop : Pop from a stack
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
    res := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return res, nil
}
