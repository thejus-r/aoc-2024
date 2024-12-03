package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// For Part 1
// RegEx to match 'mul(num,num)'
// (?i)mul\([0-9]+,[0-9]+\)
func useRegex(s string) []string {
	re := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")
	return re.FindAllString(s, -1)
}

// For Part 2
// don\'t\(\).*?(do\(\))|$)
func useRegex2(s string) string {
	re := regexp.MustCompile("(?is)don\\'t\\(\\).*?(?:do\\(\\)|$)")
	return re.ReplaceAllString(s, "")
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)

	// String manupulations
	input := string(dat)

	// Part 1
	afterRegEx := useRegex(input)
	sumOfProducts := 0

	for _, val := range afterRegEx {
		val = strings.Trim(val, "mul()")
		n := strings.Split(val, ",")
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		sumOfProducts = sumOfProducts + n1*n2
	}

	fmt.Printf("Part 1: %d\n", sumOfProducts)

	// Part 2

	sumOfProducts2 := 0
	afterRegEx2 := useRegex2(input)
	newStrs := useRegex(afterRegEx2)
	for _, val := range newStrs {
		val = strings.Trim(val, "mul()")
		n := strings.Split(val, ",")
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		sumOfProducts2 = sumOfProducts2 + n1*n2
	}
	fmt.Printf("Part 2: %d\n", sumOfProducts2)

}
