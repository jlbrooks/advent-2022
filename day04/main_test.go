package main

import "testing"

func TestOverlap(t *testing.T) {
	if !overlaps(parseShift("3-5"), parseShift("4-6")) {
		t.Fail()
	}

	if !overlaps(parseShift("5-7"), parseShift("7-9")) {
		t.Fail()
	}

	if overlaps(parseShift("5-7"), parseShift("8-9")) {
		t.Fail()
	}

	if !overlaps(parseShift("7-9"), parseShift("5-7")) {
		t.Fail()
	}
}
