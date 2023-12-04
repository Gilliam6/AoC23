package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cardCounter = make(map[int]int)

func main() {
	input := os.Args[1]
	games := strings.Split(input, "\n")
	var res int
	cardCounter[0] = 1
	for i := range games {
		cardCounter[i] = 1
	}
	for i, game := range games {
		res = processGame(game)
		for index := 0; index < res; index++ {
			cardCounter[i+index+1] += 1 * cardCounter[i]
		}
	}
	res = 0
	for _, v := range cardCounter {
		res += v
	}
	fmt.Println(res)
}

func processGame(game string) (res int) {
	var winMap = make(map[int]bool)
	scores := strings.TrimSpace(strings.Split(game, ":")[1])
	splittedScores := strings.Split(scores, " | ")
	actualScore := strings.Split(splittedScores[0], " ")
	winScores := strings.Split(splittedScores[1], " ")
	for _, score := range winScores {
		num, err := strconv.Atoi(score)
		if err != nil {
			continue
		}
		winMap[num] = true
	}
	for _, actual := range actualScore {
		num, err := strconv.Atoi(actual)
		if err != nil {
			continue
		}
		if winMap[num] {
			res++
		}
	}
	fmt.Println(res)
	return
}
