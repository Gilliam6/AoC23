package main

import (
	"fmt"
	"os"
	"sort"
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
	//counter := 0
	positions := make([]string, 0)
	for k, _ := range lefts {
		if strings.HasSuffix(k, "A") {
			positions = append(positions, k)
		}
	}
	loops := make([]int, 0)
	for _, position := range positions {
		cnt := 0
		for !strings.HasSuffix(position, "Z") {
			move := moves[cnt%len(moves)]
			if move == 'R' {
				position = rights[position]
			} else {
				position = lefts[position]
			}
			cnt++
		}
		loops = append(loops, cnt)
	}
	sort.Slice(loops, func(i, j int) bool {
		return loops[i] > loops[j]
	})
	result := lcmOfNumbers(loops)

	fmt.Println(result, "moves")
}

func lcmOfNumbers(numbers []int) int {
	lcmRes := numbers[0]

	for i := 1; i < len(numbers); i++ {
		lcmRes = lcm(lcmRes, numbers[i])
	}

	return lcmRes
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
