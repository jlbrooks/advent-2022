package main

import (
	"fmt"
	"testing"
)

func TestMakeCrates(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}

	crates := makeCrates(input)
	if string(crates[0]) != "NZ" {
		t.Fail()
	}

	if string(crates[1]) != "DCM" {
		println(crates[1].toString())
		t.Fail()
	}

	if string(crates[2]) != "P" {
		t.Fail()
	}
}

func TestParseMove(t *testing.T) {
	m := parseMove("move 2 from 3 to 5")
	if m.number != 2 || m.from != 3 || m.to != 5 {
		fmt.Printf("%q\n", m)
		t.Fail()
	}
}
