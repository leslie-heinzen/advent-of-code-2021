package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type probe struct {
	x         int
	y         int
	xv        int
	yv        int
	archeight int
	success   bool
}

func (p *probe) Travel() {
	p.x = p.x + p.xv
	newHeight := p.y + p.yv
	if newHeight > p.archeight {
		p.archeight = newHeight
	}

	p.y = newHeight

	if p.xv > 0 {
		p.xv--
	}

	p.yv--
}

func (p *probe) IsInBounds(t target) bool {
	if p.x >= t.x.min && p.x <= t.x.max &&
		p.y >= t.y.min && p.y <= t.y.max {
		p.success = true
		return true
	}

	return false
}

func (p *probe) MissedTarget(t target) bool {
	if p.x > t.x.max || p.y < t.y.min {
		return true
	}

	return false
}

type minmax struct {
	min int
	max int
}
type target struct {
	x minmax
	y minmax
}

// calc lowest starting xv for probe based on cumulative drag
func (t *target) MinXV() int {
	vel := 0
	sum := 0
	if sum < t.x.min {
		vel++
		sum += vel
	}

	return vel
}

func (t *target) MaxYV() int {
	return -t.y.min - 1
}

func main() {
	target := readInput("input.txt")
	probes := simulateLaunches(target)
	highest := findHighestVertical(probes)
	successfulLandings := countSuccessfulAttempts(probes)
	fmt.Printf("Highest y velocity: %d\n", highest)
	fmt.Printf("Successful launches: %d", successfulLandings)
}

func readInput(fileName string) target {
	input, _ := os.ReadFile(fileName)
	var targetArea target

	re := regexp.MustCompile(`\-?[0-9]+`)
	for _, ta := range strings.Split(string(input), ", ") {
		axis := ta[:1]
		matches := re.FindAllString(ta[2:], -1)

		var coords []int
		for _, m := range matches {
			toInt, _ := strconv.Atoi(m)
			coords = append(coords, toInt)
		}

		if axis == "x" {
			targetArea.x = minmax{coords[0], coords[1]}
		} else {
			targetArea.y = minmax{coords[0], coords[1]}
		}
	}

	return targetArea
}

func simulateLaunches(t target) []probe {
	var probes []probe

	for xv := t.MinXV(); xv <= t.x.max; xv++ {
		for yv := t.y.min; yv <= t.MaxYV(); yv++ {
			probe := simulateLaunch(t, xv, yv)
			probes = append(probes, probe)
		}
	}

	return probes
}

func simulateLaunch(t target, xv int, yv int) probe {
	p := probe{0, 0, xv, yv, 0, false}

	for p.IsInBounds(t) || !p.MissedTarget(t) {
		if p.IsInBounds(t) {
			break
		}
		if !p.MissedTarget(t) {
			p.Travel()
		}
	}

	return p
}

func findHighestVertical(probes []probe) int {
	highest := -1
	for _, p := range probes {
		if highest == -1 || p.archeight > highest {
			highest = p.archeight
		}
	}

	return highest
}

func countSuccessfulAttempts(probes []probe) int {
	var attempts int
	for _, probe := range probes {
		if probe.success {
			attempts++
		}
	}

	return attempts
}
