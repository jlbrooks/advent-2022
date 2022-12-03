package main

import (
	"testing"
)

func compareMaps(s1, s2 map[item]bool) bool {
	for i := range s1 {
		if !s2[i] {
			return false
		}
	}
	return true
}

func TestNewRucksack(t *testing.T) {
	ruck := newRucksack("abcd")
	expectedComp1 := map[item]bool{'a': true, 'b': true}
	expectedComp2 := map[item]bool{'c': true, 'd': true}
	if !compareMaps(ruck.comp1, expectedComp1) {
		t.Fail()
	}

	if !compareMaps(ruck.comp2, expectedComp2) {
		t.Fail()
	}
}

func TestFindCommon(t *testing.T) {
	ruck := newRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	common, err := ruck.findCommon()
	if err != nil {
		t.Fail()
	}
	if common != 'p' {
		t.Fatalf("Found common %q instead of 'a'", common)
	}
}

func TestFindCommonWithCapitals(t *testing.T) {
	ruck := newRucksack("TZZjzzZLfZbzgzZNNJZjwCVbwMmhwCbBpCMMBCbM")
	common, err := ruck.findCommon()
	if err != nil {
		t.Fail()
	}
	if common != 'b' {
		t.Fatalf("Found common %q instead of 'c'", common)
	}
}

func TestValues(t *testing.T) {
	initAlphabet()
	if item('a').value() != 1 {
		t.Fatalf("Got %d expected 1", item('a').value())
	}

	if item('z').value() != 26 {
		t.Fatalf("Got %d expected 26", item('z').value())
	}

	if item('B').value() != 28 {
		t.Fatalf("Got %d expected 28", item('B').value())
	}
}
