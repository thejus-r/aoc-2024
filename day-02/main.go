package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// convert to str -> int
func GenerateSliceOfInts(s []string) []int {

	arr := make([]int, len(s))
	for index, c := range s {
		n, _ := strconv.Atoi(c)
		arr[index] = n
	}
	return arr
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)

	// String manupulations
	input := string(dat)
	strarr := strings.Split(input, "\n")

	noOfLines := len(strarr)
	noOfSafeReports := 0
	noOfDampedReports := 0

	arr := make([][]int, noOfLines)

	for index, value := range strarr {
		value := strings.Split(value, " ")
		arr[index] = GenerateSliceOfInts(value)
	}

	// PART 1

	// PART 2

	fmt.Printf("No of safe reports: %d\n", noOfSafeReports)
	fmt.Printf("No of damped reports: %d\n", noOfDampedReports)
}
