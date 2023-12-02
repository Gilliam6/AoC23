package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var winsIdSum int
	input := os.Args[1]
	splittedLines := strings.Split(input, "\n")

	for _, line := range splittedLines {
		winsIdSum += checkGame(line)
	}
	fmt.Println(winsIdSum)
}

func checkGame(game string) int {
	var mins = map[string]int{
		"green": 0,
		"blue":  0,
		"red":   0,
	}
	splittGame := strings.Split(game, ":")
	gameScore := 1
	sets := strings.Split(splittGame[1], ";")
	for _, set := range sets {
		checkSet(set, mins)
	}
	for _, v := range mins {
		gameScore *= v
	}
	return gameScore
}

func checkSet(set string, mins map[string]int) {
	splittSet := strings.Split(set, ",")
	for _, colorRes := range splittSet {
		splittColor := strings.Split(colorRes, " ")
		score, _ := strconv.Atoi(splittColor[1])
		if mins[splittColor[2]] < score {
			mins[splittColor[2]] = score
		}
	}
}
