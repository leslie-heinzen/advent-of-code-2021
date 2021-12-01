package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput parses the input text file.
func ReadInput(fname string) []int {
	input, _ := os.ReadFile("input.txt")

	depths := []int{}

	for _, s := range strings.Fields(string(input)) {
		i, _ := strconv.Atoi(s)
		depths = append(depths, i)
	}

	return depths
}

// Part1 solves the first part of the Day 1 problem.
func Part1(depths []int) int {
	count := 0

	for i, v := range depths[1:] {
		if v > depths[i] {
			count++
		}
	}

	return count
}

// Part2 solves the second part of the Day 1 problem.
func Part2(depths []int) int {
	count := 0

	for i, v := range depths[3:] {
		if v > depths[i] {
			count++
		}
	}

	return count
}

func main() {
	var depths = ReadInput("input.txt")
	var solution1 = Part1(depths)
	var solution2 = Part2(depths)

	fmt.Printf("Part1: %d\nPart2: %d", solution1, solution2)
}
