package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	grid := make(map[point]int)

	lines := strings.Fields(string(input))

	for y, line := range lines {
		nums := getLineLevels(line)
		for x, level := range nums {
			grid[point{x, y}] = level
		}
	}

	flashCount := 0
	step := 1

	for {
		flashed := map[point]bool{}

		var flash func(p point)
		flash = func(p point) {
			if flashed[p] {
				return
			}

			flashed[p] = true
			flashCount++

			adj := []point{
				{p.x - 1, p.y}, {p.x - 1, p.y - 1}, {p.x - 1, p.y + 1},
				{p.x + 1, p.y}, {p.x + 1, p.y - 1}, {p.x + 1, p.y + 1},
				{p.x, p.y - 1}, {p.x, p.y + 1},
			}

			for _, ap := range adj {
				if _, ok := grid[ap]; !ok {
					continue
				}

				if flashed[ap] {
					continue
				}

				grid[ap] = grid[ap] + 1
				if grid[ap] > 9 {
					flash(ap)
				}
			}
		}

		for point, val := range grid {
			grid[point] = val + 1

			if grid[point] > 9 {
				flash(point)
			}
		}

		for point := range flashed {
			grid[point] = 0
		}

		if step == 100 {
			fmt.Printf("Solution 1: %d\n", flashCount)
		}

		if len(flashed) == len(grid) {
			fmt.Printf("Solution 2: %d\n", step)
			break
		}

		step++
	}
}

func getLineLevels(line string) []int {
	lineLevels := []int{}

	for _, num := range strings.Split(line, "") {
		cast, _ := strconv.Atoi(num)
		lineLevels = append(lineLevels, cast)
	}

	return lineLevels
}
