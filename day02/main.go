package main

import (
	"fmt"
	"strings"

	"github.com/jlbrooks/advent-2022/shared"
)

type Action uint8
type Outcome uint8

const (
	Rock Action = iota
	Paper
	Scissors
)

const (
	Win Outcome = iota
	Lose
	Draw
)

var winners = map[Action]Action{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var losers = map[Action]Action{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func oppInputToAction(letter string) Action {
	switch letter {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		panic(letter)
	}
}

func playerInputToAction(letter string) Action {
	switch letter {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic(letter)
	}
}

func playerInputToOutcome(letter string) Outcome {
	switch letter {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		panic(letter)
	}
}

func actionForOutcome(opp Action, desired Outcome) Action {
	switch desired {
	case Win:
		return losers[opp]
	case Draw:
		return opp
	case Lose:
		return winners[opp]
	default:
		panic(opp)
	}
}

func (a Action) points() int {
	switch a {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic(a)
	}
}

func (o Outcome) points() int {
	switch o {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	default:
		panic(o)
	}
}

func playOutcome(player, opponent Action) Outcome {
	if winners[player] == opponent {
		return Win
	} else if player == opponent {
		return Draw
	} else {
		return Lose
	}
}

func main() {
	lines := shared.ReadLines("day02/input.txt")
	sum := 0
	sum2 := 0
	for _, l := range lines {
		plays := strings.Split(l, " ")
		if plays[0] == "" {
			continue
		}
		opp := oppInputToAction(plays[0])

		// part 1
		player := playerInputToAction(plays[1])
		outcome := playOutcome(player, opp)
		sum += outcome.points() + player.points()

		// part 2
		desired := playerInputToOutcome(plays[1])
		player = actionForOutcome(opp, desired)
		sum2 += desired.points() + player.points()
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
