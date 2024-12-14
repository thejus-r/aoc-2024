package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

type Robot struct {
	px, py, vx, vy int
}

var HEIGHT, WIDTH int
var robots []Robot

func main() {

	part1 := false

	parseInput("input.txt")

	// HEIGHT, WIDTH = 7, 11
	HEIGHT, WIDTH = 103, 101
	NO_OF_SECONDS := 5273

	if !part1 {
		for itr := 1; itr < 103*101; itr++ {
			for id, r := range robots {
				robots[id].px = utils.SimpleMod(r.px+r.vx*itr, WIDTH)
				robots[id].py = utils.SimpleMod(r.py+r.vy*itr, HEIGHT)
			}

			if checkForTree() {
				fmt.Printf("Found at : %d\n", itr)
			}
		}
	}
	if part1 {
		for id, r := range robots {
			robots[id].px = utils.SimpleMod(r.px+r.vx*NO_OF_SECONDS, WIDTH)
			robots[id].py = utils.SimpleMod(r.py+r.vy*NO_OF_SECONDS, HEIGHT)
		}

		sf := findSafetyFactor()
		generateMap()

		fmt.Printf("Safety Factor: %d\n", sf)
	}

}

func checkForTree() bool {

	var sb strings.Builder

	pos := make(map[[2]int]bool)
	for _, r := range robots {
		pos[[2]int{r.px, r.py}] = true
	}

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			if pos[[2]int{j, i}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('-')
			}
		}
		sb.WriteRune('\n')
	}

	if strings.Contains(sb.String(), "##########") {
		fmt.Println(sb.String())
		return true
	}
	return false
}

func findSafetyFactor() int {
	y := HEIGHT / 2
	x := WIDTH / 2

	var q1, q2, q3, q4 int

	for _, r := range robots {
		if r.px == x || r.py == y {
			continue
		}
		if r.px < x && r.py < y {
			q1++
		}
		if r.px > x && r.py < y {
			q2++
		}
		if r.px < x && r.py > y {
			q3++
		}
		if r.px > x && r.py > y {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func generateMap() {

	pos := make(map[[2]int]int)
	for _, r := range robots {
		pos[[2]int{r.px, r.py}]++
	}

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			val, ok := pos[[2]int{j, i}]
			if ok {
				fmt.Printf("%d", val)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func parseInput(fileName string) {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)
	str := string(dat)

	lines := strings.Split(str, "\n")

	re := regexp.MustCompile(`-?\d+`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		newRobot := Robot{
			px: utils.MustAtoi(matches[0]),
			py: utils.MustAtoi(matches[1]),
			vx: utils.MustAtoi(matches[2]),
			vy: utils.MustAtoi(matches[3]),
		}
		robots = append(robots, newRobot)
	}
}
