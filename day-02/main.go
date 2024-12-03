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

func removeOneEl(arr []int, index int) []int {
	temp := []int{}
	temp = append(temp, arr[:index]...)
	temp = append(temp, arr[index+1:]...)
	return temp
}

func checkNeighbors(arr []int, index int) bool {
	temp := removeOneEl(arr, index)
	if checkArrayDirection(temp) && checkArrayDeviation(temp) {
		return true
	}
	temp = removeOneEl(arr, index+1)
	if checkArrayDirection(temp) && checkArrayDeviation(temp) {
		return true
	}
	return false
}

func checkForDampableDeviation(arr []int) (bool, int) {
	for i, v := range arr {
		if Abs(v) >= 3 {
			return true, i
		}
	}
	return false, -1
}

func checkArrayDeviation(arr []int) bool {
	allOk := true
	for _, v := range arr {
		if Abs(v) > 0 && Abs(v) <= 3 {
			allOk = true
		} else {
			return false
		}
	}
	return allOk
}

func checkArrayDirection(arr []int) bool {
	allPositive := true
	allNegative := true

	for _, v := range arr {
		if v > 0 {
			allNegative = false
		} else if v < 0 {
			allPositive = false
		}

		if !allPositive && !allNegative {
			return false
		}
	}

	if allPositive {
		return true
	}
	if allNegative {
		return true
	}
	return false
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
	diffArr := [][]int{}
	for i := 0; i < noOfLines; i++ {
		tempArr := []int{}
		for j := 0; j < len(arr[i])-1; j++ {
			tempArr = append(tempArr, arr[i][j+1]-arr[i][j])
		}
		diffArr = append(diffArr, tempArr)
	}

	for i := 0; i < noOfLines; i++ {
		uniform := checkArrayDirection(diffArr[i])
		deviation := checkArrayDeviation(diffArr[i])
		if uniform && deviation {
			noOfSafeReports++
		} else if checkIfEqualIsDampable(diffArr[i]) {
			noOfDampedReports++
		} else if deviation && uniform != false {
			// one uniform check
			// todo
			noOfDampedReports++
		} else if uniform && deviation != false {
			ok, index := checkForDampableDeviation(diffArr[i])
			if ok {
				if checkNeighbors(arr[i], index) {
					noOfDampedReports++
				}
			}
		}
	}

	// PART 2
	fmt.Println(arr)
	fmt.Println(diffArr)
	fmt.Printf("No of safe reports: %d\n", noOfSafeReports)
	fmt.Printf("No of damped reports: %d\n", noOfDampedReports)
}

func checkIfDeviationIsDampable(arr []int) bool {

	return false
}

func checkIfEqualIsDampable(arr []int) bool {
	for i, v := range arr {
		temp := []int{}
		if v == 0 {
			temp = append(temp, arr[:i]...)
			temp = append(temp, arr[i+1:]...)
			uniform := checkArrayDirection(temp)
			deviation := checkArrayDeviation(temp)
			if uniform && deviation {
				return true
			}
		}
	}
	return false
}
