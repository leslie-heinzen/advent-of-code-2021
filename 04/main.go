package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	value  int
	marked bool
}

type Positions [5][5]Position

func (positions *Positions) Mark(draw int) {
	for rIdx, r := range positions {
		for cIdx, c := range r {
			if c.value == draw && !c.marked {
				positions[rIdx][cIdx].marked = true
			}
		}
	}
}

func (positions *Positions) Score(draw int) int {
	sum := 0
	for _, p := range positions {
		for _, c := range p {
			if !c.marked {
				sum += c.value
			}
		}
	}

	return sum * draw
}

func (positions *Positions) Reset() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			positions[i][j].marked = false
		}
	}
}

func (positions *Positions) CheckBingo() bool {
	rowCount := 0

	for i := 0; i < 5; i++ {
		rowCount = 0
		for j := 0; j < 5; j++ {
			if positions[i][j].marked {
				rowCount++
			}
		}
		if rowCount == 5 {
			break
		}
	}

	flipped := transpose(positions)
	colCount := 0
	for i := 0; i < 5; i++ {
		colCount = 0
		for j := 0; j < 5; j++ {
			if flipped[i][j].marked {
				colCount++
			}
		}
		if colCount == 5 {
			break
		}
	}

	if rowCount == 5 || colCount == 5 {
		return true
	}

	return false
}

func main() {
	drawsInput, _ := os.ReadFile("draws.txt")
	boardsInput, _ := os.ReadFile("boards.txt")
	drawsInputStrs := strings.Split(string(drawsInput), ",")
	draws := []int{}

	for _, d := range drawsInputStrs {
		res, _ := strconv.Atoi(d)
		draws = append(draws, res)
	}

	var boards = make([]Positions, 100)
	boardsIdx := 0
	rowIdx := 0

	for i := range boards {
		boards[i] = [5][5]Position{}
	}

	for _, b := range strings.Split(string(boardsInput), "\r\n") {
		if b == "" {
			boardsIdx++
			rowIdx = 0
			continue
		}

		b = strings.TrimSpace(b)
		numStrs := strings.Fields(b)

		for i, n := range numStrs {
			converted, _ := strconv.Atoi(n)
			boards[boardsIdx][rowIdx][i].value = converted
			boards[boardsIdx][rowIdx][i].marked = false
		}

		rowIdx++
	}

	solution1 := Part1(draws, boards)
	for i := range boards {
		boards[i].Reset()
	}
	solution2 := Part2(draws, boards)
	fmt.Printf("Solution1: %d", solution1)
	fmt.Printf("Solution2: %d", solution2)

}

func play(draw int, boards []Positions, part1 bool) ([]int, int, bool, []Positions) {
	var lastScore int
	var lastHasBingo bool

	for i := range boards {
		boards[i].Mark(draw)
	}

	for i := len(boards) - 1; i >= 0; i-- {
		hasBingo := boards[i].CheckBingo()

		if hasBingo {
			if part1 {
				return nil, boards[i].Score(draw), hasBingo, boards
			} else {
				lastHasBingo = hasBingo
				lastScore = boards[i].Score(draw)
				boards = append(boards[:i], boards[i+1:]...)
			}
		}
	}

	if lastHasBingo {
		return nil, lastScore, lastHasBingo, boards
	}

	return nil, -1, false, boards
}

func Part1(draws []int, boards []Positions) int {
	var winningScore int

	for _, d := range draws {
		_, score, hasBingo, _ := play(d, boards, true)

		if hasBingo {
			winningScore = score
			break
		}
	}

	return winningScore
}

func Part2(draws []int, boards []Positions) int {
	var finalScore int
	for len(boards) > 0 {
		for _, d := range draws {
			_, score, hasBingo, updated := play(d, boards, false)
			if hasBingo {
				finalScore = score
			}
			boards = updated
		}

		for i := range boards {
			boards[i].Reset()
		}
	}

	return finalScore
}

func transpose(slice *Positions) [][]Position {
	transposed := make([][]Position, len(slice[0]))
	for i := range transposed {
		transposed[i] = make([]Position, len(slice))
	}
	for i, row := range slice {
		for j, val := range row {
			transposed[j][i] = val
		}
	}
	return transposed
}
