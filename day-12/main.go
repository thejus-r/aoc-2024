package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	stack "github.com/thejus-r/aoc-2024/utils/ds"
)

type Point struct {
	r, c int
}

var seen map[Point]bool
var plot [][]rune
var maxRow, maxCol int
var regions [][]Point

func main() {
	plot = parseInput()
	maxRow, maxCol = len(plot), len(plot[0])

	// Initializing maps
	seen = make(map[Point]bool)

	findAllRegions()

	fmt.Printf("Total cost: %d\n", findTotalCost())
}

func findAllRegions() {
	for r := range plot {
		for c := range plot[0] {
			if seen[Point{r: r, c: c}] {
				continue
			}
			seen[Point{r: r, c: c}] = true
			stk := stack.New[Point]()
			stk.Push(Point{r: r, c: c})
			currCrop := plot[r][c]
			region := []Point{{r: r, c: c}}
			for !stk.IsEmpty() {
				curr, _ := stk.Pop()
				for _, next := range []Point{{curr.r - 1, curr.c}, {curr.r + 1, curr.c}, {curr.r, curr.c - 1}, {curr.r, curr.c + 1}} {
					if next.r < 0 || next.c < 0 || next.r >= maxRow || next.c >= maxCol {
						continue
					}
					if plot[next.r][next.c] != currCrop {
						continue
					}
					if slices.Contains(region, next) {
						continue
					}
					region = append(region, next)
					seen[next] = true
					stk.Push(next)
				}
			}
			regions = append(regions, region)
		}
	}
}

func findTotalCost() int {
	cost := 0
	for _, region := range regions {
		cost = cost + len(region)*findPerimeter(region)
	}
	return cost
}

func findPerimeter(region []Point) int {
	output := 0
	for _, curr := range region {
		output += 4
		for _, next := range []Point{{curr.r - 1, curr.c}, {curr.r + 1, curr.c}, {curr.r, curr.c - 1}, {curr.r, curr.c + 1}} {
			if slices.Contains(region, next) {
				output -= 1
			}
		}
	}
	return output

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
