package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

var cache map[[2]int]int // Used to memoize the results of recursive function
var stones []int

func main() {
	noOfBlinks := 75 // For Part 1 noOfBlinks := 25
	cache = make(map[[2]int]int)
	var noOfStones int = 0
	parseInput("input.txt")

	// iterating through all the initial stones
	for _, stone := range stones {
		noOfStones = noOfStones + int(countStones(stone, noOfBlinks))
	}

	fmt.Printf("Number of stones: %d\n", noOfStones)
}

func countStones(stone, blinks int) int {
	var result int

	// Check if cache already has the value
	val, ok := cache[[2]int{stone, blinks}]
	if ok {
		return val // if yes, return that
	}

	if blinks == 0 {
		result = 1 // Last Blink so 1 stone
	} else if stone == 0 {
		result = countStones(1, blinks-1)
	} else {
		str := fmt.Sprintf("%d", stone)
		len := len(str)
		if len%2 == 0 {
			result = 0
			result += countStones(utils.MustAtoi(str[:len/2]), blinks-1)
			result += countStones(utils.MustAtoi(str[len/2:]), blinks-1)
		} else {
			result = countStones(stone*2024, blinks-1)
		}
	}
	cache[[2]int{stone, blinks}] = result // store value in cache first
	return cache[[2]int{stone, blinks}]
}

// Parsing input to a Array of Int
func parseInput(fileName string) {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)
	str := string(dat)
	for _, el := range strings.Split(str, " ") {
		stones = append(stones, utils.MustAtoi(el))
	}
}
