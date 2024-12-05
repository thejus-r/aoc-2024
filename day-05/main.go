package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

func parseOrders(ordering string) map[int][]int {
	result := make(map[int][]int)
	for _, order := range strings.Split(ordering, "\n") {
		numbers := strings.Split(order, "|")
		f, s := utils.MustAtoi(numbers[0]), utils.MustAtoi(numbers[1])
		result[f] = append(result[f], s)
	}
	return result
}

func parseSequence(sequences string) [][]int {
	seq := strings.Split(sequences, "\n")
	result := make([][]int, len(seq))

	for row, line := range seq {
		sequences := strings.Split(line, ",")
		numbers := make([]int, len(sequences))
		for i, num := range sequences {
			numbers[i] = utils.MustAtoi(num)
		}
		result[row] = numbers
	}

	return result
}

func parseInput(fileName string) (map[int][]int, [][]int) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// String manupulations
	input := string(dat)

	pattern := `(\r?\n){2,}`

	re := regexp.MustCompile(pattern)

	parts := re.Split(input, -1)

	ruleMap := parseOrders(parts[0])
	sequences := parseSequence(parts[1])
	return ruleMap, sequences
}

func main() {
	ruleMap, sequences := parseInput("input.txt")

	sumOfMiddlePages := 0
	for _, eachSeq := range sequences {
		checkOrder(ruleMap, eachSeq, &sumOfMiddlePages)
	}

	fmt.Printf("the ans is :%d\n", sumOfMiddlePages)
}

func checkOrder(rules map[int][]int, sequence []int, sumOfMiddlePages *int) {
	isOrdered := true
	visited := []int{}
	for _, el := range sequence {
		visited = append(visited, el)
		val := rules[el]
		for _, v := range val {
			if slices.Contains(visited, v) {
				isOrdered = false
			}
		}

	}
	if isOrdered {
		*sumOfMiddlePages = *sumOfMiddlePages + sequence[len(sequence)/2]
	}
}
