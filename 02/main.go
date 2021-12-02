package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(commands []string) int {
	var x, y int = 0, 0

	for _, c := range commands {
		c = strings.TrimSpace(c)
		parts := strings.Split(c, " ")
		direction := parts[0]
		change, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			x += change
		case "down":
			y += change
		case "up":
			y -= change
		}
	}

	return x * y
}

func Part2(commands []string) int {
	var x, y, z int = 0, 0, 0

	for _, c := range commands {
		c = strings.TrimSpace(c)
		parts := strings.Split(c, " ")
		direction := parts[0]
		change, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			x += change
			y += (change * z)
		case "down":
			z += change
		case "up":
			z -= change
		}
	}

	return x * y
}

func main() {
	input, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(input), "\n")
	solution1 := Part1(commands)
	solution2 := Part2(commands)
	fmt.Printf("Part1: %d\nPart2: %d", solution1, solution2)
}
