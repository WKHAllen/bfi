package main

import (
	"fmt"
)

// BFTape : Brainfuck tape object
type BFTape struct {
	tape []int8
	head int
	size int
}

// NewBFTape : Create a new tape object
func NewBFTape() *BFTape {
	bft := &BFTape{size: 1}
	bft.tape = append(bft.tape, 0)
	return bft
}

// Get : Get the value at the tape head
func (bft *BFTape) Get() int8 {
	return bft.tape[bft.head]
}

// Inc : Increment the value at the tape head
func (bft *BFTape) Inc() {
	bft.tape[bft.head]++
}

// Dec : Decrement the value at the tape head
func (bft *BFTape) Dec() {
	bft.tape[bft.head]--
}

// MoveLeft : Move the tape head to the left
func (bft *BFTape) MoveLeft() error {
	if bft.head == 0 {
		return fmt.Errorf("head cannot move left")
	}
	bft.head--
	return nil
}

// MoveRight : Move the tape head to the right
func (bft *BFTape) MoveRight() {
	if bft.head == bft.size {
		bft.tape = append(bft.tape, 0)
	}
	bft.head++
}
