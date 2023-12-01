package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

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
	end := len(line) - 1
	start := 0
	var first = -1
	var last = -1
	for i, _ := range line {
		if first >= 0 && last >= 0 {
			break
		}
		if first == -1 && unicode.IsDigit(rune(line[start+i])) {
			first = int(line[start+i]) - 48
		}
		if last == -1 && unicode.IsDigit(rune(line[end-i])) {
			last = int(line[end-i]) - 48
		}
	}
	fmt.Println(first*10 + last)
	return first*10 + last
}
