package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func main() {
	input, _ := os.ReadFile("input.txt")
	directions := strings.Split(string(input), "\r\n")
	lines := []line{}

	for _, d := range directions {
		var x1, y1, x2, y2 int

		fmt.Sscanf(d, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		lines = append(lines, line{point{x1, y1}, point{x2, y2}})
	}

	solution1 := countOverlaps(lines, false)
	solution2 := countOverlaps(lines, true)

	fmt.Printf("Part1 Solution: %d\n", solution1)
	fmt.Printf("Part2 Solution: %d", solution2)
}

func countOverlaps(lines []line, countDiagonal bool) int {
	touches := map[point]int{}

	for _, l := range lines {
		startX, targetX := l.start.x, l.end.x
		startY, targetY := l.start.y, l.end.y

		if startX == targetX {
			y0, y1 := order(startY, targetY)

			for y := y0; y <= y1; y++ {
				touches[point{targetX, y}]++
			}
		} else if startY == targetY {
			x0, x1 := order(startX, targetX)

			for x := x0; x <= x1; x++ {
				touches[point{x, targetY}]++
			}
		} else if countDiagonal {
			x0, y0, x1, y1 := orderPoints(startX, startY, targetX, targetY)
			dy := getVerticalDirection(x0, y0, x1, y1)

			for x, y := x0, y0; x <= x1; x, y = x+1, y+dy {
				touches[point{x, y}]++
			}
		}

	}

	var touchesTotal int
	for _, g := range touches {
		if g >= 2 {
			touchesTotal++
		}
	}

	return touchesTotal
}

func getVerticalDirection(x0 int, y0 int, x1 int, y1 int) int {
	var dy int

	if y0 < y1 {
		dy = 1
	} else {
		dy = -1
	}

	return dy
}

func orderPoints(x0 int, y0 int, x1 int, y1 int) (int, int, int, int) {
	if x1 < x0 {
		x1, x0 = x0, x1
		y1, y0 = y0, y1
	}

	return x0, y0, x1, y1
}

func order(a int, b int) (int, int) {
	if a <= b {
		return a, b
	}

	return b, a
}
