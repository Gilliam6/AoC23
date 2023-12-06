package main

import (
	"fmt"
	"math"
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
	time, _ := strconv.Atoi(strings.Join(times, ""))
	record, _ := strconv.Atoi(strings.Join(records, ""))
	res *= checkBest(float64(time), float64(record))

	fmt.Println(res)
}

func checkBest(time float64, record float64) int {
	/*
		time * x - x^2 > record
		-x^2 + time*x - record > 0
		x^2 - time*x + record < 0
		x1 := [7 + sqrt((7)^2 - 4*1*9)] / (2*1)
		x2 := [7 - sqrt((7)^2 - 4*1*9)] / (2*1)
	*/
	x1 := (time - math.Sqrt(math.Pow(time, 2)-4*1*record)) / 2.0
	x2 := (time + math.Sqrt(math.Pow(time, 2)-4*1*record)) / 2.0
	fmt.Printf("%.2f --> %.2f\n", x1, x2)

	if x1 == math.Ceil(x1) {
		x1++
	}
	x1 = math.Ceil(x1)
	if x2 == math.Floor(x2) {
		x2--
	}
	x2 = math.Floor(x2)

	fmt.Println("Solutions:", x2-x1+1)
	return int(x2 - x1 + 1)
}
