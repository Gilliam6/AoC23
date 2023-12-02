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
	splittGame := strings.Split(game, ":")
	gameScore, _ := strconv.Atoi(strings.Split(splittGame[0], " ")[1])
	sets := strings.Split(splittGame[1], ";")
	for _, set := range sets {
		if checkSet(set) {
			return 0
		}
	}
	return gameScore
}

func checkSet(set string) bool {
	var limits = map[string]int{
		"green": 13,
		"blue":  14,
		"red":   12,
	}
	splittSet := strings.Split(set, ",")
	for _, colorRes := range splittSet {
		splittColor := strings.Split(colorRes, " ")
		score, _ := strconv.Atoi(splittColor[1])
		limits[splittColor[2]] -= score
		if limits[splittColor[2]] < 0 {
			//fmt.Println("LOSS")
			return true
		}
	}
	return false
}
