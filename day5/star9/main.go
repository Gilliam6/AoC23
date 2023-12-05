package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MinMaxDst struct {
	min int
	max int
	dst int
}

func (s *MinMaxDst) MinMaxDst(seed int) int {
	if seed < s.max && seed >= s.min {
		return s.dst + (seed - s.min)
	}
	return seed
}

func main() {
	var seedMaps = make([][]MinMaxDst, 0)
	input := os.Args[1]
	splittedInput := strings.Split(input, "\n\n")
	for _, routes := range splittedInput[1:] {
		splitRoute := strings.Split(routes, ":")
		mapsArr := make([]MinMaxDst, 0)
		values := strings.Split(splitRoute[1], "\n")
		generateRanges(&mapsArr, values)
		seedMaps = append(seedMaps, mapsArr)
	}
	seeds := strings.Fields(splittedInput[0])[1:]
	var minimal = math.MaxInt
	for _, seed := range seeds {
		seedNum, _ := strconv.Atoi(seed)
		loc := findPlaceForSeed(seedNum, seedMaps)
		if loc < minimal {
			minimal = loc
		}
	}
	fmt.Println(minimal)
}

func generateRanges(ranges *[]MinMaxDst, values []string) {
	for _, line := range values {
		if len(line) == 0 {
			continue
		}
		nums := strings.Fields(line)
		size, _ := strconv.Atoi(nums[2])
		startDst, _ := strconv.Atoi(nums[0])
		startSrc, _ := strconv.Atoi(nums[1])
		*ranges = append(*ranges, MinMaxDst{
			min: startSrc,
			max: startSrc + size,
			dst: startDst,
		})
	}
}

func findPlaceForSeed(seed int, routes [][]MinMaxDst) int {
	for _, route := range routes {
		for _, ranges := range route {
			newVal := ranges.MinMaxDst(seed)
			if newVal != seed {
				seed = newVal
				break
			}
		}
	}
	return seed
}
