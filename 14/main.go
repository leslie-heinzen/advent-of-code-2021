package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	template, rules := readInput("input.txt")

	sol1 := run(template, rules, 10)
	sol2 := run(template, rules, 40)

	fmt.Printf("Solution 1: %d\n", sol1)
	fmt.Printf("Solution 2: %d\n", sol2)
}

func readInput(fileName string) (string, map[string]string) {
	input, _ := os.ReadFile(fileName)
	lines := strings.Split(string(input), "\r\n")
	template := lines[0]
	rules := map[string]string{}

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	return template, rules
}

func run(template string, rules map[string]string, steps int) int {
	pairCounts := countPairs(template, rules, steps)
	elementCounts := countElements(template, pairCounts)
	result := subMinMax(elementCounts)
	return result
}

func countPairs(template string, rules map[string]string, steps int) map[string]int {
	pairCounts := map[string]int{}

	for i := 1; i < len(template); i++ {
		pair := string(template[i-1]) + string(template[i])
		pairCounts[pair]++
	}

	for i := 0; i < steps; i++ {
		updatedCounts := map[string]int{}
		for k, v := range pairCounts {
			updatedCounts[string(k[0])+rules[k]] += v
			updatedCounts[rules[k]+string(k[1])] += v
		}

		pairCounts = updatedCounts
	}

	return pairCounts
}

func countElements(template string, pairCounts map[string]int) map[string]int {
	elementCounts := map[string]int{}
	for k, v := range pairCounts {
		char := string(k[1])
		elementCounts[char] += v
	}

	elementCounts[string(template[0])]++

	return elementCounts
}

func subMinMax(elementCounts map[string]int) int {
	countsSlice := []int{}
	for _, v := range elementCounts {
		countsSlice = append(countsSlice, v)
	}

	sort.Ints(countsSlice)
	diff := countsSlice[len(countsSlice)-1] - countsSlice[0]
	return diff
}
