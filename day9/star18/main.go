package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	splittedInp := strings.Split(input, "\n")
	res := 0

	for _, line := range splittedInp {
		res += processLine(line)
	}
	fmt.Println("RES:", res)
}

func processLine(line string) int {
	fmt.Println(line)

	var res = make([]int, 0)
	var arr = make([]int, 0)

	fields := strings.Fields(line)
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		arr = append(arr, num)
	}
	for !zeros(arr) {
		var newArr = make([]int, 0)

		res = append([]int{arr[0]}, res...)
		for i := 1; i < len(arr); i++ {
			newArr = append(newArr, arr[i]-arr[i-1])
		}
		arr = newArr
	}
	for i := 1; i < len(res); i++ {
		res[i] = res[i] - res[i-1]
	}
	//fmt.Println(line)
	//fmt.Println(res)
	fmt.Println("LINE RESULT:", res[len(res)-1])
	fmt.Println("================")
	return res[len(res)-1]
}

func zeros(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}
