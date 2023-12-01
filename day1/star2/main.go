package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var dictionary = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var res int
	file, err := os.OpenFile("./input", os.O_RDONLY, 0755)
	if err != nil {
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		res += digitsFromLine(string(line))
	}
	fmt.Println(res)
}

func digitsFromLine(line string) (res int) {
	var arr = make([]int, 0)

	for i, _ := range line {
		if unicode.IsDigit(rune(line[i])) {

			arr = append(arr, int(line[i])-48)
			continue
		}
		for k, v := range dictionary {
			if strings.HasPrefix(line[i:], k) {
				fmt.Println("Line ", line[i:], " has prefix ", k)
				arr = append(arr, v)
				break
			}
		}
	}
	fmt.Println("ARRAY: ", arr)

	fmt.Println(arr[0]*10 + arr[len(arr)-1])
	return arr[0]*10 + arr[len(arr)-1]
}
