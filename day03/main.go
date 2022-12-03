package main

import (
	"errors"
	"strings"

	"github.com/jlbrooks/advent-2022/shared"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var alphabetItems = []item(alphabet)
var itemValues = make(map[item]int, 52)

func initAlphabet() {
	for i := 0; i < len(alphabetItems); i++ {
		itemValues[alphabetItems[i]] = i + 1
		upper := strings.ToUpper(string(alphabetItems[i]))
		itemValues[[]item(upper)[0]] = i + 1 + 26
	}
}

type item rune
type compartment map[item]bool

type rucksack struct {
	comp1, comp2, combined compartment
	len                    int
}

func (i item) value() int {
	return itemValues[i]
}

func newRucksack(contents string) *rucksack {
	contentsArr := []item(contents)
	halfIndex := len(contents) / 2

	comp1 := make(compartment, halfIndex)
	comp2 := make(compartment, halfIndex)
	combined := make(compartment, len(contents))
	for i := 0; i < len(contents); i++ {
		if i < halfIndex {
			comp1[contentsArr[i]] = true
		} else {
			comp2[contentsArr[i]] = true
		}
		combined[contentsArr[i]] = true
	}

	return &rucksack{
		comp1:    comp1,
		comp2:    comp2,
		combined: combined,
		len:      halfIndex,
	}
}

func commonItem(c1, c2 compartment) (item, error) {
	for i := range c1 {
		if c2[i] {
			return i, nil
		}
	}

	return 0, errors.New("No common items")
}

func commonBadge(c1, c2, c3 compartment) (item, error) {
	for i := range c1 {
		if c2[i] && c3[i] {
			return i, nil
		}
	}

	return 0, errors.New("No common items for badge")
}

func (r rucksack) findCommon() (item, error) {
	return commonItem(r.comp1, r.comp2)
}

func main() {
	initAlphabet()

	sum := 0
	groupSum := 0
	curGroup := make([]*rucksack, 3)
	groupCount := 0
	for _, l := range shared.ReadLines("day03/input.txt") {
		if l == "" {
			continue
		}

		ruck := newRucksack(l)

		// part 1
		common, err := ruck.findCommon()
		if err == nil {
			sum += common.value()
		}

		// part 2
		curGroup[groupCount] = ruck
		if groupCount == 2 {
			badge, err := commonBadge(curGroup[0].combined, curGroup[1].combined, curGroup[2].combined)
			if err != nil {
				panic("No badge for group")
			}
			groupSum += badge.value()
			groupCount = 0
		} else {
			groupCount += 1
		}
	}

	println(sum)
	println(groupSum)
}
