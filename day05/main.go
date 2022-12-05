package main

import (
	"fmt"
	"regexp"
	"strconv"

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

var moveRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func (c crate) toString() string {
	return string(c)
}

func makeCrates(lines []string) []crate {
	crates := make([]crate, (len(lines[0])/4)+1)
	println(lines[0])
	println((len(lines[0]) / 4) + 1)
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
					crates[crateNum] = append(crates[crateNum], char)
				}
			}
		}
	}

	return crates
}

func parseMove(line string) *move {
	matches := moveRegex.FindStringSubmatch(line)
	if len(matches) != 4 {
		panic(fmt.Sprintf("Invalid matches found for move: %q", matches))
	}
	num, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	return &move{
		number: num,
		from:   from,
		to:     to,
	}
}

func (c *crane) evalMove(m move) {
	for i := 0; i < m.number; i++ {
		fromCrate := c.crates[m.from-1]
		toCrate := c.crates[m.to-1]
		if len(fromCrate) == 0 {
			continue
		}

		elem, fromCrate := fromCrate[0], fromCrate[1:]
		toCrate = append(crate{elem}, toCrate...)

		c.crates[m.from-1] = fromCrate
		c.crates[m.to-1] = toCrate
	}
}

func main() {
	lines := shared.ReadLines("day05/input.txt")
	st := parseCrates
	crateLines := make([]string, 0)
	cr := &crane{
		crates: make([]crate, 0),
	}
	for _, l := range lines {
		if st == parseCrates {
			crateLines = append(crateLines, l)
		} else {
			if l != "" {
				m := parseMove(l)
				cr.evalMove(*m)
			}
		}

		// got to the end of the parse phase
		if l == "" {
			cr.crates = makeCrates(crateLines[:len(crateLines)-1])
			st = parseMoves
		}
	}

	output := make([]rune, len(cr.crates))
	for _, cra := range cr.crates {
		output = append(output, cra[0])
	}

	println(string(output))
}
