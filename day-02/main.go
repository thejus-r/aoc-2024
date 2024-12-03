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
	for _, report := range arr {
		if isReportSafe(report) {
			noOfSafeReports++
		} else if isReportDampable(report) {
			noOfDampedReports++
		}

	}
	fmt.Printf("No of safe reports: %d\n", noOfSafeReports)

	// PART 2
	fmt.Printf("No of damped reports: %d\n", noOfDampedReports)
	fmt.Printf("Total Safe Reports: %d\n", noOfSafeReports+noOfDampedReports)
}

func isReportSafe(report []int) bool {
	isSafe := false
	isIncreasing := true
	isDecreasing := true

	for idx := 0; idx < len(report)-1; idx++ {

		if report[idx+1] < report[idx] {
			isIncreasing = false
		}
		if report[idx+1] > report[idx] {
			isDecreasing = false
		}
		if report[idx+1] == report[idx] {
			isIncreasing = false
			isDecreasing = false
		}

		if Abs(report[idx+1]-report[idx]) <= 3 && (isIncreasing != isDecreasing) {
			isSafe = true
		} else {
			return false
		}
	}
	return isSafe
}

func delOneEl(arr []int, idx int) []int {
	temp := []int{}
	temp = append(temp, arr[:idx]...)
	temp = append(temp, arr[idx+1:]...)
	return temp
}

func isReportDampable(report []int) bool {
	for idx := 0; idx < len(report); idx++ {
		temp := delOneEl(report, idx)
		fmt.Println(temp)
		if isReportSafe(temp) {
			return true
		}
	}
	return false
}
