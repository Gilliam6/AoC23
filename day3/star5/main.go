package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var arrayNumPos = make(map[int]bool)
var arraySymbolRange = make(map[int]bool)

func main() {
	var res int
	field := os.Args[1]
	twoDimArr := strings.Split(field, "\n")
	size := len(twoDimArr[0])

	for y, line := range twoDimArr {
		for x, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				checkField(twoDimArr, x, y, len(line))
			}
		}
	}
	for k, v := range arraySymbolRange {
		if v == true && unicode.IsDigit(rune(twoDimArr[k/size][k%size])) {
			res += findNum(twoDimArr, k, size)
		}
	}

	fmt.Println(res)
}

func checkField(field []string, x int, y int, size int) {
	for i := -1; i < 2; i++ {
		for n := 1; n > -2; n-- {
			if x+n < size && x+n > -1 && y+i > -1 && y+i < len(field) {
				if n == 0 && i == 0 {
					continue
				}
				arraySymbolRange[size*(y+i)+x+n] = true
			}
		}
	}
}

func findNum(twoDimArr []string, pos int, size int) (number int) {
	for ; pos%size > 0 && unicode.IsDigit(rune(twoDimArr[pos/size][pos%size])); pos-- {
	}
	if !unicode.IsDigit(rune(twoDimArr[pos/size][pos%size])) {
		pos++
	}
	line := pos / size
	for pos/size == line && unicode.IsDigit(rune(twoDimArr[pos/size][pos%size])) {
		number *= 10
		number += int(twoDimArr[pos/size][pos%size]) - 48

		arraySymbolRange[pos] = false
		pos++
	}
	return number
}
