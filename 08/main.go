package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type IOPair struct {
	input  []string
	output []string
}

func main() {
	file, _ := os.Open("input.txt")
	lines := []string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	ioPairs := []IOPair{}
	re := regexp.MustCompile(`([a-g]+)`)

	for i, l := range lines {
		if i%2 == 1 {
			input := re.FindAllString(lines[i-1], -1)
			output := re.FindAllString(l, -1)
			pair := IOPair{input, output}
			ioPairs = append(ioPairs, pair)
		}
	}

	sol1 := part1(ioPairs)
	sol2 := part2(ioPairs)
	fmt.Printf("Solution1: %d", sol1)
	fmt.Printf("Solution2: %d", sol2)
}

func part1(pairs []IOPair) int {
	count := 0

	for _, p := range pairs {
		for _, o := range p.output {
			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	return count
}

func part2(pairs []IOPair) int {
	var sum int

	for _, p := range pairs {
		var signals = make(map[string]string)
		var twoFiveOrThree []string
		var zeroSixOrNine []string

		for _, i := range p.input {
			switch len(i) {
			case 2:
				signals["1"] = i
			case 3:
				signals["7"] = i
			case 4:
				signals["4"] = i
			case 5:
				twoFiveOrThree = append(twoFiveOrThree, i)
			case 6:
				zeroSixOrNine = append(zeroSixOrNine, i)
			case 7:
				signals["8"] = i
			}
		}

		for _, zsn := range zeroSixOrNine {
			if compare(zsn, signals["4"]) {
				signals["9"] = zsn
			} else if compare(zsn, signals["1"]) {
				signals["0"] = zsn
			} else {
				signals["6"] = zsn
			}
		}

		for _, tft := range twoFiveOrThree {
			if compare(tft, signals["1"]) {
				signals["3"] = tft
			} else if compare(signals["9"], tft) {
				signals["5"] = tft
			} else {
				signals["2"] = tft
			}
		}

		var display string
		for _, o := range p.output {
			for k, v := range signals {
				oLen := len(o)
				vLen := len(v)
				var contains bool

				if oLen == vLen {
					contains = compare(o, v)
					if contains {
						display += k
						break
					}
				}
			}
		}

		res, _ := strconv.Atoi(display)
		sum += res
	}

	return sum
}

func compare(a string, b string) bool {
	var contains bool

	for _, v := range b {
		contains = strings.Contains(a, string(v))

		if !contains {
			break
		}
	}

	return contains
}
