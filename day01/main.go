package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jlbrooks/advent-2022/shared"
)

func main() {
	lines := shared.ReadLines("day01/input.txt")
	curSum := 0
	biggest := 0
	var elves []int
	for _, l := range lines {
		num, err := strconv.Atoi(l)
		if err == nil {
			curSum += num
		} else {
			elves = append(elves, curSum)
			if curSum > biggest {
				biggest = curSum
			}
			curSum = 0
		}
	}
	elves = append(elves, curSum)

	fmt.Printf("Biggest: %d\n", biggest)
	sort.Ints(elves)

	sum := 0
	arrLen := len(elves)
	for _, num := range elves[arrLen-3 : arrLen] {
		sum += num
	}

	fmt.Printf("Top 3: %d\n", sum)
}
