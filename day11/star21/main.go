package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	splittedInp := strings.Split(input, "\n")
	fmt.Println(splittedInp)
}
