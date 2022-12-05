package main

import (
	"github.com/jlbrooks/advent-2022/shared"
)

type state uint8

const (
	parseCrates state = iota
	parseMoves
)

type crate []rune

type crane struct {
	crates []crate
}

type move struct {
	number, from, to int
}

func (c crate) toString() string {
	return string(c)
}

func makeCrates(lines []string) []crate {
	crates := make([]crate, len(lines[0])/4+1)
	println(len(crates))
	for i := 0; i < len(crates); i++ {
		crates[i] = make(crate, 0)
	}

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			// 1, 5, 9 are the numbers
			if (j-1)%4 == 0 {
				crateNum := (j - 1) / 4
				char := rune(line[j])
				if char != ' ' {
					crates[crateNum] = append(crates[(j-1)/4], char)
				}
			}
		}
	}

	return crates
}

func parseMove(line string) *move {
	return &move{}
}

func (c *crane) evalMove(m move) {

}

func main() {
	lines := shared.ReadLines("day05/input.txt")
	st := parseCrates
	crateLines := make([]string, 10)
	cr := &crane{
		crates: make([]crate, 0),
	}
	for _, l := range lines {
		if st == parseCrates {
			crateLines = append(crateLines, l)
		} else {
			m := parseMove(l)
			cr.evalMove(*m)
		}

		// got to the end of the parse phase
		if l == "" {
			cr.crates = makeCrates(crateLines)
			st = parseMoves
		}
	}

	output := make([]rune, len(cr.crates))
	for _, cra := range cr.crates {
		output = append(output, cra[0])
	}

	println(string(output))
}
