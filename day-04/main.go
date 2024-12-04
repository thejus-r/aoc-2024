package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type index struct {
	r, c int
}

var directionsV1 = []index{
	{r: 0, c: 1},
	{r: 0, c: -1},
	{r: 1, c: 0},
	{r: -1, c: 0},
	{r: -1, c: -1},
	{r: 1, c: 1},
	{r: 1, c: -1},
	{r: -1, c: 1},
}

func main() {
	input := parseInput()
	fmt.Printf("Part One: %d\n", partOne(input))

}

func parseInput() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func partOne(grid [][]rune) (wordCount int) {
	const word = "XMAS"
	maxRow, maxCol := len(grid), len(grid[0])

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == rune(word[0]) {
				for _, dir := range directionsV1 {
					findXMAS(grid, index{r: r, c: c}, dir, 0, word, &wordCount, maxRow, maxCol)
				}
			}
		}
	}

	return wordCount
}

func findXMAS(grid [][]rune, idx index, dir index, at int, word string, wordCount *int, maxRow, maxCol int) {
	if at+1 == len(word) {
		(*wordCount)++
		return
	}
	neigbourIndex := index{r: idx.r + dir.r, c: idx.c + dir.c}
	if isValidIndex(neigbourIndex, maxRow, maxCol) {
		neighbourRune := grid[neigbourIndex.r][neigbourIndex.c]
		if neighbourRune == rune(word[at+1]) {
			findXMAS(grid, neigbourIndex, dir, at+1, word, wordCount, maxRow, maxCol)
		}
	}
}

func isValidIndex(idx index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}
