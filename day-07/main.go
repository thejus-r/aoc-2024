// AoC 2024, Day 07
// Bridge Repair

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

var Results []int
var Operators [][]int

func parseInput(fileName string) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// String manupulations
	input := string(dat)

	eachLine := strings.Split(input, "\n")

	for _, line := range eachLine {
		parts := strings.Split(line, ": ")
		Results = append(Results, utils.MustAtoi(parts[0]))
		ops := strings.Split(parts[1], " ")
		var temp []int
		for _, op := range ops {
			temp = append(temp, utils.MustAtoi(op))
		}
		Operators = append(Operators, temp)
	}
}

func main() {
	parseInput("input.txt")

	sumOfResults1 := 0
	sumOfResults2 := 0
	for idx, result := range Results {
		if isPossible(result, Operators[idx], false) {
			sumOfResults1 = sumOfResults1 + result
		}
	}

	for idx, result := range Results {
		if isPossible(result, Operators[idx], true) {
			sumOfResults2 = sumOfResults2 + result
		}
	}

	fmt.Printf("sum of results (Part 1): %d\n", sumOfResults1)
	fmt.Printf("sum of results (Part 2): %d\n", sumOfResults2)
}

func isPossible(target int, ops []int, useConcat bool) bool {
	var f func(int, int) bool
	f = func(idx, total int) bool {
		if total > target {
			return false
		}
		if idx == len(ops) {
			return total == target
		}
		c := false
		if useConcat {
			c = f(idx+1, concat(total, ops[idx]))
		}

		return f(idx+1, total+ops[idx]) || f(idx+1, total*ops[idx]) || c
	}

	return f(1, ops[0])
}

func concat(a, b int) int {
	return utils.MustAtoi(fmt.Sprintf("%v%v", a, b))
}
