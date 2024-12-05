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

func parseOrders(ordering string) (map[int][]int, [][]int) {
	inOrder := make(map[int][]int)
	rules := [][]int{}
	for _, order := range strings.Split(ordering, "\n") {
		numbers := strings.Split(order, "|")
		f, s := utils.MustAtoi(numbers[0]), utils.MustAtoi(numbers[1])
		inOrder[f] = append(inOrder[f], s)
		rule := []int{f, s}
		rules = append(rules, rule)
	}
	return inOrder, rules
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

func parseInput(fileName string) (map[int][]int, [][]int, [][]int) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// String manupulations
	input := string(dat)

	pattern := `(\r?\n){2,}`

	re := regexp.MustCompile(pattern)

	parts := re.Split(input, -1)

	ruleMap, rules := parseOrders(parts[0])
	sequences := parseSequence(parts[1])
	return ruleMap, rules, sequences
}

func main() {
	ruleMap, _, sequences := parseInput("input.txt")

	sumOfMiddlePages := 0
	sumOfMiddlePagesCorrected := 0

	for _, eachSeq := range sequences {
		checkOrder(ruleMap, eachSeq, &sumOfMiddlePages, &sumOfMiddlePagesCorrected)
	}

	fmt.Printf("the ans for Part 1 is :%d\n", sumOfMiddlePages)
	fmt.Printf("the ans for Part 2 is :%d\n", sumOfMiddlePagesCorrected)
}

func checkOrder(ruleMap map[int][]int, sequence []int, sumOfMiddlePages *int, sumOfMiddlePagesCorrected *int) {
	isOrdered := true
	visited := []int{}
	for _, el := range sequence {
		visited = append(visited, el)
		val := ruleMap[el]
		for _, v := range val {
			if slices.Contains(visited, v) {
				isOrdered = false
			}
		}
	}
	if isOrdered {
		*sumOfMiddlePages = *sumOfMiddlePages + sequence[len(sequence)/2]
	} else {
		result := repairOrder(ruleMap, sequence)
		*sumOfMiddlePagesCorrected = *sumOfMiddlePagesCorrected + result[len(result)/2]
	}
}

func repairOrder(ruleMap map[int][]int, sequence []int) []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, el := range sequence {
		if !visited[el] {
			processOrder(el, visited, &stack, ruleMap, sequence)
		}
	}
	return stack
}

func processOrder(el int, visited map[int]bool, stack *[]int, ruleMap map[int][]int, sequence []int) {
	visited[el] = true

	for _, u := range ruleMap[el] {
		if !visited[u] && slices.Contains(sequence, u) {
			processOrder(u, visited, stack, ruleMap, sequence)
		}
	}
	*stack = append([]int{el}, *stack...)
}
