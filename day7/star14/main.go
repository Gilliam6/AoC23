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
		'J': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
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
	pairs := make([]struct {
		Key   rune
		Value int
	}, 0)

	for _, char := range hand {
		counter[char]++
	}

	for k, v := range counter {
		if k != 'J' {
			pairs = append(pairs, struct {
				Key   rune
				Value int
			}{Key: k, Value: v})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	rate := 1

	if counter['J'] == 5 {
		return 7
	}
	for _, pair := range pairs {
		if pair.Value+counter['J'] == 5 {
			return 7 // five of kind
		} else if pair.Value+counter['J'] == 4 {
			return 6 // four of kind
		} else if pair.Value+counter['J'] == 3 && rate == 2 {
			return 5 // full house
		} else if pair.Value+counter['J'] == 3 {
			counter['J'] = 0
			rate = 4 // three of kind
		} else if pair.Value+counter['J'] == 2 && rate == 4 {
			return 5 // full house
		} else if pair.Value+counter['J'] == 2 && rate == 2 {
			return 3 // two pairs
		} else if pair.Value+counter['J'] == 2 {
			counter['J'] = 0
			rate = 2 // one pair
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
