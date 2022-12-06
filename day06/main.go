package main

import (
	"fmt"

	"github.com/jlbrooks/advent-2022/shared"
)

type sequence []rune

func (seq sequence) isDistinct() bool {
	chars := make(map[rune]bool, len(seq))
	for _, c := range seq {
		if _, exists := chars[c]; exists {
			return false
		}
		chars[c] = true
	}

	return true
}

func (seq sequence) findStart() int {
	for i := 4; i < len(seq); i++ {
		window := seq[i-4 : i]
		if window.isDistinct() {
			return i
		}
	}

	return -1
}

func (seq sequence) findFirstPacket() int {
	for i := 14; i < len(seq); i++ {
		window := seq[i-14 : i]
		if window.isDistinct() {
			return i
		}
	}

	return -1
}

func main() {
	lines := shared.ReadLines("day06/input.txt")
	input := sequence(lines[0])

	// part 1
	start := input.findStart()
	fmt.Printf("Start: %d\n", start)

	// part 2
	packet := input.findFirstPacket()
	fmt.Printf("Packet: %d\n", packet)
}
