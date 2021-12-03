package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// I'm pretty embarrassed of this solution.
// I am not familiar enough with Go's built-in
// functions yet!

func Part1(input []string) int64 {
	gamma := []string{}
	epsilon := []string{}

	for i := 0; i < 12; i++ {
		oneCount := 0
		zeroCount := 0

		for _, v := range input {
			v = strings.TrimSpace(v)
			digit, _ := strconv.Atoi(string(v[i]))
			if digit == 1 {
				oneCount++
			} else {
				zeroCount++
			}
		}

		if oneCount > zeroCount {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	fmt.Print(gamma)
	gammaStr := strings.Join(gamma, "")
	fmt.Print(gammaStr)
	epsilonStr := strings.Join(epsilon, "")
	g, _ := strconv.ParseInt(gammaStr, 2, 64)
	e, _ := strconv.ParseInt(epsilonStr, 2, 64)
	result := g * e
	return result
}

func GetOxygenRating(input []string) int64 {
	var oxStr string
	oxArr := make([]string, len(input))
	copy(oxArr, input)

	for i := 0; i < 12; i++ {
		var oneIndices []int
		var zeroIndices []int

		for j, v := range oxArr {
			trimmed := strings.TrimSpace(v)
			digit, _ := strconv.ParseInt(string(trimmed[i]), 0, 64)
			if digit == 1 {
				oneIndices = append([]int{j}, oneIndices...)
			} else {
				zeroIndices = append([]int{j}, zeroIndices...)
			}
		}

		if len(oneIndices) >= len(zeroIndices) {
			for _, v := range zeroIndices {
				oxArr = append(oxArr[:v], oxArr[v+1:]...)
			}
		} else {
			for _, v := range oneIndices {
				oxArr = append(oxArr[:v], oxArr[v+1:]...)
			}
		}

		fmt.Printf("ox len: %d\n", len(oxArr))

		if len(oxArr) == 1 {
			oxStr = strings.TrimSpace(oxArr[0])
		}
	}

	o, _ := strconv.ParseInt(oxStr, 2, 64)

	return o
}

func GetCo2Rating(input []string) int64 {
	var scrubberStr string
	scrubberArr := make([]string, len(input))
	copy(scrubberArr, input)

	for i := 0; i < 12; i++ {
		var oneIndices []int
		var zeroIndices []int

		for j, v := range scrubberArr {
			trimmed := strings.TrimSpace(v)
			digit, _ := strconv.ParseInt(string(trimmed[i]), 0, 64)
			if digit == 1 {
				oneIndices = append([]int{j}, oneIndices...)
			} else {
				zeroIndices = append([]int{j}, zeroIndices...)
			}
		}

		if len(oneIndices) >= len(zeroIndices) {
			for _, v := range oneIndices {
				scrubberArr = append(scrubberArr[:v], scrubberArr[v+1:]...)
			}
		} else {
			for _, v := range zeroIndices {
				scrubberArr = append(scrubberArr[:v], scrubberArr[v+1:]...)
			}
		}

		fmt.Printf("len: %d\n", len(scrubberArr))
		if len(scrubberArr) == 1 {
			scrubberStr = strings.TrimSpace(scrubberArr[0])
		}
	}

	s, _ := strconv.ParseInt(scrubberStr, 2, 64)

	return s
}

func Part2(input []string) int64 {
	o := GetOxygenRating(input)
	s := GetCo2Rating(input)
	return o * s
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	solution1 := Part1(lines)
	solution2 := Part2(lines)
	fmt.Printf("Part1: %d", solution1)
	fmt.Printf("Part2: %d", solution2)
}
