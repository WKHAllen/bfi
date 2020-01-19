package main

import (
	"fmt"
)

// Interpreter return codes
const (
	BFIDone = iota
	BFIInput = iota
	BFIOutput = iota
)

// BFInterpreter : Brainfuck interpreter object
type BFInterpreter struct {
	tape *BFTape
	code string
	index int
	whileStack *Stack
}

// NewBFInterpreter : Create a new interpreter object
func NewBFInterpreter(code string) *BFInterpreter {
	bfi := &BFInterpreter{
		tape: NewBFTape(),
		code: code,
		index: 0,
		whileStack: NewStack(),
	}
	return bfi
}

// Interpret : Interpret brainfuck code
func (bfi *BFInterpreter) Interpret() (int, error) {
	// TODO: interpret the brainfuck code
	// when '[': add the current index to the whileStack
	// when ']': pop from the whileStack and return to that position
	// return BFIDone when done
	// return BFIInput when user input is required (',')
	// return BFIOutput when '.'
	// return an error if ']' is found and whileStack is empty
	// return an error if tape head is moved left of zero
	return 0, fmt.Errorf("error")
}
