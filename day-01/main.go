package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// This function returns absolute value of an integer
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

	var inputArr1 []int
	var inputArr2 []int

	for _, value := range strarr {
		// splits the string around each instance of one or more consecutive white space characters
		temp := strings.Fields(value)
		arrEl1, _ := strconv.Atoi(temp[0])
		arrEl2, _ := strconv.Atoi(temp[1])
		inputArr1 = append(inputArr1, arrEl1)
		inputArr2 = append(inputArr2, arrEl2)
	}

	// PART 1
	// sorting the arrays
	slices.Sort(inputArr1)
	slices.Sort(inputArr2)

	length := len(inputArr1)

	sum := 0

	for i := 0; i < length; i++ {
		sum = sum + Abs(inputArr1[i]-inputArr2[i])
	}

	fmt.Printf("Total distance: %d\n", sum)

	// PART 2
	// Finding Frequency Map of second array
	freqMap := make(map[int]int)
	for i := 0; i < length; i++ {
		val, ok := freqMap[inputArr2[i]]
		if ok {
			freqMap[inputArr2[i]] = val + 1
		} else {
			freqMap[inputArr2[i]] = 1
		}
	}

	// Finding similarity score
	similarityScore := 0

	for i := 0; i < length; i++ {
		val, ok := freqMap[inputArr1[i]]
		if ok {
			similarityScore = similarityScore + inputArr1[i]*val
		}
	}

	fmt.Printf("Similarity score: %d\n", similarityScore)

}
