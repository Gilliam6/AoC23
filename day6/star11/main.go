package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	splitInp := strings.Split(input, "\n")
	times := strings.Fields(strings.Split(splitInp[0], ":")[1])
	records := strings.Fields(strings.Split(splitInp[1], ":")[1])
	res := 1
	for i := range times {
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])
		res *= checkBest(time, record)
	}
	fmt.Println(res)
}

func checkBest(time int, record int) int {
	hold := 0
	cnt := 0
	for hold < time {
		if (1*hold)*(time-hold) > record {
			cnt++
		}
		hold++
	}
	fmt.Println(cnt)
	return cnt
}
