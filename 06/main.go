package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countFish(school []int, days int) {
	var day int
	schoolByTimer := [9]int{}

	for _, f := range school {
		schoolByTimer[f]++
	}

	for day < days {
		newSchoolByTimer := [9]int{}
		for idx, v := range schoolByTimer {
			if idx == 0 {
				newSchoolByTimer[6] += v
				newSchoolByTimer[8] += v
			} else {
				newSchoolByTimer[idx-1] += v
			}
		}
		schoolByTimer = newSchoolByTimer
		day++
	}

	var total int
	for _, f := range schoolByTimer {
		total += f
	}

	fmt.Printf("Fish after %d days: %d\n", days, total)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	school := []int{}

	for _, f := range strings.Split(string(input), ",") {
		fish, _ := strconv.Atoi(f)
		school = append(school, fish)
	}

	countFish(school, 80)
	countFish(school, 256)
}
