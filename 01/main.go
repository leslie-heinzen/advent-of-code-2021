package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(fname string) []int {
	input, _ := os.ReadFile("input.txt")

	depths := []int{}

	for _, s := range strings.Fields(string(input)) {
		i, _ := strconv.Atoi(s)
		depths = append(depths, i)
	}

	return depths
}

func Accumulate(arr []int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func CountIncrements(arr []int, window int) int {
	count := 0
	prev := Accumulate(arr[0:window])

	slice := arr[window:]

	for i := range slice {
		curr := Accumulate(slice[i : i+window])
		if curr > prev {
			count++
		}

		prev = curr
	}

	return count
}

func main() {
	var depths = ReadInput("input.txt")
	var solution1 = CountIncrements(depths, 1)
	var solution2 = CountIncrements(depths, 3)

	fmt.Printf("Part1: %d\nPart2: %d", solution1, solution2)
}
