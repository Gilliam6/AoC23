package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	cardScore = map[rune]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

func main() {
	input := os.Args[1]
	splitInput := strings.Split(input, "\n")
	offset := 1
	res := 0
	var rewards = make(map[string]int)
	var scoreboard = make(map[int][]string)
	for _, line := range splitInput {
		splitLine := strings.Fields(line)
		reward, _ := strconv.Atoi(splitLine[1])
		rewards[splitLine[0]] = reward
	}
	for k, _ := range rewards {
		rate := rating(k)
		scoreboard[rate] = append(scoreboard[rate], k)
	}
	for _, v := range scoreboard {
		less := func(i, j int) bool {
			return compareLess(v[i], v[j])
		}
		sort.Slice(v, less)
	}
	for i := 1; i < 8; i++ {
		for _, hand := range scoreboard[i] {
			res += offset * rewards[hand]
			offset++
		}
	}
	fmt.Println(res)
}

func rating(hand string) int {
	var counter = make(map[rune]int)

	for _, char := range hand {
		counter[char]++
	}
	rate := 1
	for _, v := range counter {
		if v == 5 {
			return 7 // five of kind
		} else if v == 4 {
			return 6 // four of kind
		} else if v == 3 && rate == 2 {
			return 5
		} else if v == 3 {
			rate = 4 // three of kind
		} else if v == 2 && rate == 4 {
			return 5 // full house
		} else if v == 2 && rate == 2 {
			return 3
		} else if v == 2 {
			rate = 2
		}
	}
	return rate
}

func compareLess(s1, s2 string) bool {
	for i := range s1 {
		if cardScore[rune(s1[i])] > cardScore[rune(s2[i])] {
			return false
		} else if cardScore[rune(s1[i])] < cardScore[rune(s2[i])] {
			return true
		}
	}
	return false
}
