package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalcMovement(commands []string, calcAim bool) int {
	var x, y, z int = 0, 0, 0

	for _, c := range commands {
		c = strings.TrimSpace(c)
		parts := strings.Split(c, " ")
		direction := parts[0]
		change, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			x += change
			if calcAim {
				y += (change * z)
			}
		case "down":
			if calcAim {
				z += change
			} else {
				y += change
			}

		case "up":
			if calcAim {
				z -= change
			} else {
				y -= change
			}
		}
	}

	return x * y
}

func main() {
	input, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(input), "\n")
	solution1 := CalcMovement(commands, false)
	solution2 := CalcMovement(commands, true)
	fmt.Printf("Part1: %d\nPart2: %d", solution1, solution2)
}
