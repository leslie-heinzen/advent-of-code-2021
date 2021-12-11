package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	positions := []int{}

	for _, p := range strings.Split(string(input), ",") {
		position, _ := strconv.Atoi(p)
		positions = append(positions, position)
	}

	solution1 := part1(positions)
	solution2 := part2(positions)

	fmt.Printf("Part 1 solution: %d\n", solution1)
	fmt.Printf("Part 2 solution: %d", solution2)
}

func part1(positions []int) int {
	var lowestFuelCost int
	for i := range positions {
		var fuelCost int
		for _, p := range positions {
			cost := p - i
			if cost < 0 {
				cost = -cost
			}
			fuelCost += cost
		}

		if i == 0 || fuelCost < lowestFuelCost {
			lowestFuelCost = fuelCost
		}
	}

	return lowestFuelCost
}

func part2(positions []int) int {
	var lowestFuelCost int
	for i := range positions {
		var fuelCost int
		for _, p := range positions {
			diff := p - i
			if diff < 0 {
				diff = -diff
			}

			series := make([]int, diff)
			var seriesTotal int
			for idx := range series {
				seriesTotal += (idx + 1)
			}

			fuelCost += seriesTotal
		}

		if i == 0 || fuelCost < lowestFuelCost {
			lowestFuelCost = fuelCost
		}
	}

	return lowestFuelCost
}
