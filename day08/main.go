package main

import (
	"fmt"
	"strconv"

	"github.com/jlbrooks/advent-2022/shared"
)

type direction uint8

const (
	up direction = iota
	down
	left
	right
)

var directions = []direction{up, down, left, right}

type forest struct {
	trees         [][]int
	width, height int
}

func fromLines(lines []string) forest {
	width := len(lines[0])
	height := len(lines)
	var trees [][]int
	for _, l := range lines {
		var treeLine []int
		for i := 0; i < len(l); i++ {
			c := l[i]
			d, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			treeLine = append(treeLine, d)
		}

		trees = append(trees, treeLine)
	}

	return forest{
		trees:  trees,
		width:  width,
		height: height,
	}
}

func (f forest) at(x, y int) int {
	return f.trees[y][x]
}

func (f forest) isVisible(x, y int) bool {
	cur := f.at(x, y)
	if x == 0 || y == 0 || x == f.width-1 || y == f.height-1 {
		return true
	}

	allShorter := true
	// left
	for l := 0; l < x; l++ {
		if f.at(l, y) >= cur {
			allShorter = false
			break
		}
	}
	if allShorter {
		return true
	}

	// right
	allShorter = true
	for r := x + 1; r < f.width; r++ {
		if f.at(r, y) >= cur {
			allShorter = false
			break
		}
	}
	if allShorter {
		return true
	}

	// up
	allShorter = true
	for u := 0; u < y; u++ {
		if f.at(x, u) >= cur {
			allShorter = false
			break
		}
	}
	if allShorter {
		return true
	}

	// down
	allShorter = true
	for d := y + 1; d < f.height; d++ {
		if f.at(x, d) >= cur {
			allShorter = false
			break
		}
	}
	if allShorter {
		return true
	}

	return false
}

func (f forest) distInDir(x, y int, d direction) int {
	cur := f.at(x, y)

	if x == 0 || y == 0 || x == f.width-1 || y == f.height-1 {
		return 0
	}

	if d == up {
		tot := 0
		for u := y - 1; u > -1; u-- {
			tot += 1
			if f.at(x, u) >= cur {
				break
			}
		}
		return tot
	}

	if d == down {
		tot := 0
		for d := y + 1; d < f.height; d++ {
			tot += 1
			if f.at(x, d) >= cur {
				break
			}
		}
		return tot
	}

	if d == left {
		tot := 0
		for l := x - 1; l > -1; l-- {
			tot += 1
			if f.at(l, y) >= cur {
				break
			}
		}
		return tot
	}

	if d == right {
		tot := 0
		for r := x + 1; r < f.width; r++ {
			tot += 1
			if f.at(r, y) >= cur {
				break
			}
		}
		return tot
	}

	panic(d)
}

func (f forest) score(x, y int) int {
	total := 1
	for _, d := range directions {
		total *= f.distInDir(x, y, d)
	}

	return total
}

func (f forest) countVisible() int {
	sum := 0
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			if f.isVisible(x, y) {
				if !(x == 0 || y == 0 || x == f.width-1 || y == f.height-1) {
					//fmt.Printf("(%d,%d)\n", x, y)
				}
				sum += 1
			}
		}
	}

	return sum
}

func (f forest) bestScore() int {
	best := 0
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			new := f.score(x, y)
			//fmt.Printf("%d ", new)
			if new > best {
				best = new
			}
		}
		//fmt.Printf("\n")
	}

	return best
}

func main() {
	lines := shared.ReadLines("day08/input.txt")
	forest := fromLines(lines)
	fmt.Printf("%d\n", forest.countVisible())
	fmt.Printf("%d\n", forest.bestScore())
}
