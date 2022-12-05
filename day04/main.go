package main

import (
	"strconv"
	"strings"

	"github.com/jlbrooks/advent-2022/shared"
)

type shift struct {
	start, end int
}

func parseShift(s string) *shift {
	spl := strings.Split(s, "-")
	start, _ := strconv.Atoi(spl[0])
	end, _ := strconv.Atoi(spl[1])
	return &shift{
		start: start,
		end:   end,
	}
}

func contains(s1, s2 *shift) bool {
	return s1.start >= s2.start && s1.end <= s2.end
}

func (s shift) contains(i int) bool {
	return (i >= s.start && i <= s.end)
}

func overlaps(s1, s2 *shift) bool {
	return s1.contains(s2.start) || s1.contains(s2.end) || s2.contains(s1.start) || s2.contains(s1.end)
}

func main() {
	lines := shared.ReadLines("day04/input.txt")

	sum1 := 0
	sum2 := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		spl := strings.Split(l, ",")
		shift1 := parseShift(spl[0])
		shift2 := parseShift(spl[1])

		// part 1
		if contains(shift1, shift2) || contains(shift2, shift1) {
			sum1 += 1
		}

		// part 2
		if overlaps(shift1, shift2) {
			sum2 += 1
		}
	}
	println(sum1)
	println(sum2)
}
