package main

import "testing"

func TestFindStart(t *testing.T) {
	inputs := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for seq, expected := range inputs {
		res := sequence(seq).findStart()
		if res != expected {
			t.Errorf("Got %d, expected %d for seq %s", res, expected, seq)
		}
	}
}
