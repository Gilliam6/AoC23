package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	games := strings.Split(input, "\n")
	var res int
	for _, game := range games {
		res += processGame(game)
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
			if res == 0 {
				res++
			} else {
				res *= 2
			}
		}
	}
	return
}
