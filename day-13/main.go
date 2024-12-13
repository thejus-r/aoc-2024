package main

import (
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/thejus-r/aoc-2024/utils"
)

type Machine struct {
	Ax, Ay, Bx, By int
	Px, Py         int
}

func main() {

	// for part 2, pass "true"
	machines := parseInput("input.txt", true)

	totalCost := 0

	// From equation
	// (Ax * S) + (Bx * T) = Px and (Ay * S) + (By * T) = Py :- This will have unique solution
	// we can solve S = (Px * By - Py * Bx) / (Ax * By - Ay * Bx) and T = (Px - Ax * S) / Bx
	for _, m := range machines {
		ca := float64(m.Px*m.By-m.Py*m.Bx) / float64(m.Ax*m.By-m.Ay*m.Bx)
		cb := (float64(m.Px) - float64(m.Ax)*ca) / float64(m.Bx)

		// Check for Int -> Button cannot be pressed in a fraction
		if math.Mod(ca, 1) == 0 && math.Mod(cb, 1) == 0 {
			totalCost += int(ca)*3 + int(cb)
		}
	}
	fmt.Printf("No of Tokens: %d\n", totalCost)
}

func parseInput(fileName string, part2 bool) []Machine {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)
	str := string(dat)

	re := regexp.MustCompile(`-?\d+`)

	matches := re.FindAllString(str, -1)

	machines := []Machine{}

	for i := 0; i < len(matches); i += 6 {
		m := Machine{
			Ax: utils.MustAtoi(matches[0+i]),
			Ay: utils.MustAtoi(matches[1+i]),
			Bx: utils.MustAtoi(matches[2+i]),
			By: utils.MustAtoi(matches[3+i]),
			Px: utils.MustAtoi(matches[4+i]),
			Py: utils.MustAtoi(matches[5+i]),
		}
		if part2 {
			m.Px += 10000000000000
			m.Py += 10000000000000
		}
		machines = append(machines, m)
	}
	return machines
}
