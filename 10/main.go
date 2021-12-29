package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := readInput("example.txt")
	_, illegalChars, corrupted := iterateLines(lines, map[int]bool{})
	lineCompletions, _, _ := iterateLines(lines, corrupted)
	part1Score := calcCorruptedScore(illegalChars)
	part2Score := calcAutocompleteScore(lineCompletions)
	fmt.Printf("Solution 1: %d\n", part1Score)
	fmt.Printf("Solution 2: %d", part2Score)
}

func readInput(fileName string) []string {
	file, _ := os.Open(fileName)
	lines := []string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func iterateLines(lines []string, toSkip map[int]bool) ([][]rune, []rune, map[int]bool) {
	charMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	lineCompletions := [][]rune{}
	illegalChars := []rune{}
	corrupted := map[int]bool{}

	for lIdx, line := range lines {
		stack := []rune{}

		for i, c := range line {
			if i == 0 {
				stack = append(stack, c)
				continue
			}

			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack = append(stack, c)
				continue
			}

			sLen := len(stack)
			prev := rune(stack[sLen-1])

			if charMap[prev] == c {
				stack = stack[:sLen-1]
			} else {
				illegalChars = append(illegalChars, c)
				corrupted[lIdx] = true
				break
			}
		}

		if corrupted[lIdx] {
			continue
		}

		lineCompletion := getCompletion(stack)
		lineCompletions = append(lineCompletions, lineCompletion)
	}

	return lineCompletions, illegalChars, corrupted
}

func getCompletion(stack []rune) []rune {
	lineCompletion := []rune{}
	reversed := reverse(stack)

	for _, c := range reversed {
		switch c {
		case '(':
			lineCompletion = append(lineCompletion, ')')
		case '[':
			lineCompletion = append(lineCompletion, ']')
		case '{':
			lineCompletion = append(lineCompletion, '}')
		case '<':
			lineCompletion = append(lineCompletion, '>')
		}
	}

	return lineCompletion
}

func calcAutocompleteScore(lineCompletions [][]rune) int {
	scoresMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scores := []int{}

	for _, lcs := range lineCompletions {
		total := 0
		for _, lc := range lcs {
			total = (total * 5) + scoresMap[lc]
		}

		scores = append(scores, total)
	}

	sort.Ints(scores)
	mIdx := len(scores) / 2

	return scores[mIdx]
}

func calcCorruptedScore(illegalChars []rune) int {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := 0

	for _, c := range illegalChars {
		score += scores[c]
	}

	return score
}

func reverse(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}
