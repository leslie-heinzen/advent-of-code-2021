package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	value       int
	checked     bool
	startCoords string
}

func main() {
	input, _ := os.ReadFile("input.txt")
	var points [][]Point
	for _, line := range strings.Fields(string(input)) {
		conv := []Point{}
		for _, num := range strings.Split(line, "") {
			convNum, _ := strconv.Atoi(num)
			conv = append(conv, Point{value: convNum, checked: false})
		}

		points = append(points, conv)
	}

	sol1 := part1(points)
	sol2 := part2(points)

	fmt.Printf("Solution 1: %d\n", sol1)
	fmt.Printf("Solution 2: %d", sol2)
}

func part1(points [][]Point) int {
	var risk int

	for i, l := range points {
		upIdx := i - 1
		downIdx := i + 1
		for j, p := range l {
			leftIdx := j - 1
			rightIdx := j + 1
			var down int = 10
			var up int = 10
			var left int = 10
			var right int = 10

			if leftIdx >= 0 {
				left = l[j-1].value
			}

			if rightIdx < len(l) {
				right = l[j+1].value
			}

			if upIdx >= 0 {
				up = points[i-1][j].value
			}

			if downIdx < len(points) {
				down = points[i+1][j].value
			}

			if p.value < left && p.value < right && p.value < up && p.value < down {
				risk += p.value + 1
			}
		}
	}

	return risk
}

func part2(points [][]Point) int {
	for i, l := range points {
		for j := range l {
			coords := fmt.Sprint(i, "-", j)
			recurseMap(i, j, points, coords)
		}
	}

	pointMap := make(map[string]int)
	for _, l := range points {
		for _, p := range l {
			if p.startCoords != "" {
				pointMap[p.startCoords]++
			}
		}
	}

	slice := []int{}

	for _, v := range pointMap {
		slice = append(slice, v)
	}

	sort.Ints(slice)

	slice = slice[len(slice)-3:]

	var total int

	for _, v := range slice {
		if total > 0 {
			total *= v
		} else {
			total = v
		}

	}

	return total
}

func isInBasin(a Point, b Point) bool {
	isValid := !a.checked && a.value != 9

	return isValid
}

func recurseMap(i int, j int, points [][]Point, startCoords string) {
	upIdx := i - 1
	downIdx := i + 1
	leftIdx := j - 1
	rightIdx := j + 1

	if points[i][j].checked {
		return
	}

	points[i][j].checked = true
	if points[i][j].startCoords == "" && points[i][j].value != 9 {
		points[i][j].startCoords = startCoords
	}

	val := points[i][j]

	if upIdx >= 0 {
		if isInBasin(points[upIdx][j], val) {
			recurseMap(upIdx, j, points, startCoords)
		}
	}

	if downIdx < len(points) {
		if isInBasin(points[downIdx][j], val) {
			recurseMap(downIdx, j, points, startCoords)
		}
	}

	if leftIdx >= 0 {
		if isInBasin(points[i][leftIdx], val) {
			recurseMap(i, leftIdx, points, startCoords)
		}
	}

	if rightIdx < len(points[i]) {
		if isInBasin(points[i][rightIdx], val) {
			recurseMap(i, rightIdx, points, startCoords)
		}
	}
}
