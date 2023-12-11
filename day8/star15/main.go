package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	splittedInp := strings.Split(input, "\n\n")

	moves := splittedInp[0]
	moveMap := strings.Split(splittedInp[1], "\n")

	lefts := make(map[string]string)
	rights := make(map[string]string)

	for _, line := range moveMap {
		point, leftRight, _ := strings.Cut(line, " = ")
		splitLR := strings.Split(leftRight, ", ")
		L := strings.Trim(splitLR[0], "()")
		R := strings.Trim(splitLR[1], "()")
		lefts[point] = L
		rights[point] = R
	}
	counter := 0
	position := "AAA"
	for position != "ZZZ" {
		move := moves[counter%len(moves)]
		if move == 'R' {
			position = rights[position]
		} else {
			position = lefts[position]
		}
		counter++
	}
	fmt.Println(counter, "moves")

}
