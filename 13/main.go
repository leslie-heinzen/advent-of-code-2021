package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	axis      int
}

func main() {
	paper, instructions := readInput("input-dots.txt", "input-instructions.txt")
	fold(paper, instructions)
}

func readInput(dotsFileName string, instrFileName string) ([1500][1500]int, []instruction) {
	dotsInput, _ := os.ReadFile(dotsFileName)
	instrInput, _ := os.ReadFile(instrFileName)

	paper := [1500][1500]int{} // lazy and dumb
	instructions := []instruction{}

	for _, line := range strings.Fields(string(dotsInput)) {
		parts := strings.Split(line, ",")
		convArr := []int{}

		for _, p := range parts {
			conv, _ := strconv.Atoi(p)
			convArr = append(convArr, conv)
		}

		paper[convArr[1]][convArr[0]] = 1
	}

	re := regexp.MustCompile(`([xy]=[0-9]+)`)
	for _, field := range strings.Fields(string(instrInput)) {
		found := re.FindString(field)

		if found != "" {
			parts := strings.Split(found, "=")
			axis, _ := strconv.Atoi(parts[1])
			instructions = append(instructions, instruction{parts[0], axis})
		}
	}

	return paper, instructions
}

func fold(paper [1500][1500]int, instructions []instruction) {
	p1DotCount := 0
	for i, instruction := range instructions {
		if instruction.direction == "y" {
			for x := instruction.axis; x < len(paper); x++ {
				for y := range paper[x] {
					if paper[x][y] == 1 {
						translated := instruction.axis - (x - instruction.axis)
						paper[x][y] = 0
						paper[translated][y] = 1
					}
				}
			}
		} else {
			for x := range paper {
				for y := instruction.axis; y < len(paper[x]); y++ {
					if paper[x][y] == 1 {
						translated := instruction.axis - (y - instruction.axis)
						paper[x][y] = 0
						paper[x][translated] = 1
					}
				}
			}
		}

		if i == 0 {
			// part1 solution
			for x, row := range paper {
				for y := range row {
					if paper[x][y] == 1 {
						p1DotCount++
					}
				}
			}
		}
	}

	fmt.Printf("Solution 1: %d", p1DotCount)
	writeP2Output(paper)
}

func writeP2Output(paper [1500][1500]int) {
	var sb strings.Builder
	for x, row := range paper {
		for y := range row {
			if paper[x][y] == 1 {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	f, _ := os.Create("./output.txt")
	f.WriteString(sb.String())
}
