package main

import (
	"fmt"
)

// Interpreter return codes
const (
	BFIDone = 0
	BFIInput = 1
	BFIOutput = 2
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
func (bfi *BFInterpreter) Interpret() (int, byte, error) {
	for ; bfi.index < len(bfi.code); bfi.index++ {
		switch (bfi.code[bfi.index]) {
			case '<':
				err := bfi.tape.MoveLeft()
				if err != nil {
					return -1, ' ', err
				}
			case '>':
				bfi.tape.MoveRight()
			case '-':
				bfi.tape.Dec()
			case '+':
				bfi.tape.Inc()
			case '[':
				if bfi.tape.Get() == 0 {
					err := bfi.JumpForward()
					if err != nil {
						return -1, ' ', err
					}
				} else {
					bfi.whileStack.Push(bfi.index)
				}
			case ']':
				index, err := bfi.whileStack.Peek()
				if err != nil {
					return -1, ' ', fmt.Errorf("no corresponding '[' found")
				}
				if bfi.tape.Get() == 0 {
					bfi.whileStack.Pop()
				} else {
					bfi.index = index
				}
			case '.':
				return BFIOutput, byte(bfi.tape.Get()), nil
			case ',':
				return BFIInput, ' ', nil
		}
	}
	return BFIDone, ' ', nil
}

// JumpForward : Jump to the corresponding ']'
func (bfi *BFInterpreter) JumpForward() error {
	bfi.index++
	for level := 1; level > 0; bfi.index++ {
		if bfi.index >= len(bfi.code) {
			return fmt.Errorf("no corresponding ']' found")
		}
		switch (bfi.code[bfi.index]) {
			case '[':
				level++
			case ']':
				level--
		}
	}
	return nil
}
