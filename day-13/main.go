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
	machines := parseInput("input.txt")

	totalPrizes := 0
	totalCost := 0

	for _, machine := range machines {
		minCost := math.MaxInt32
		foundSolution := false

		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				if a*machine.Ax+b*machine.Bx == machine.Px && a*machine.Ay+b*machine.By == machine.Py {
					foundSolution = true
					cost := a*3 + b*1
					if cost < minCost {
						minCost = cost
					}
				}
			}
		}

		if foundSolution {
			totalPrizes++
			totalCost += minCost
		}

	}
	fmt.Println(totalCost)
}

func parseInput(fileName string) []Machine {
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
		machines = append(machines, m)
	}
	return machines
}
