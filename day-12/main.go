package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"

	stack "github.com/thejus-r/aoc-2024/utils/ds"
)

type Point struct {
	r, c int
}

type FloatPoint struct {
	r, c float64
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

	fmt.Printf("Total cost in Part 1: %d\n", findTotalCostPart1())
	fmt.Printf("Total cost in Part 2: %d\n", findTotalCostPart2())
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
				for _, next := range getNextPossiblePoints(curr) {
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

func findTotalCostPart1() int {
	cost := 0
	for _, region := range regions {
		cost = cost + len(region)*findPerimeter(region)
	}
	return cost
}

func findTotalCostPart2() int {
	cost := 0
	for _, region := range regions {
		cost = cost + len(region)*findSides(region)
	}
	return cost
}

// Part 1
func findPerimeter(region []Point) int {
	output := 0
	for _, curr := range region {
		output += 4
		for _, next := range getNextPossiblePoints(curr) {
			if slices.Contains(region, next) {
				output -= 1
			}
		}
	}
	return output
}

func getNextPossiblePoints(curr Point) []Point {
	return []Point{{curr.r - 1, curr.c}, {curr.r + 1, curr.c}, {curr.r, curr.c - 1}, {curr.r, curr.c + 1}}
}

// Make a sub coordinate system with direction
// Edge {[1.5,2] , [-0.5, 0]}
func findSides(region []Point) int {

	edges := make(map[FloatPoint]FloatPoint)
	for _, curr := range region {
		for _, next := range getNextPossiblePoints(curr) {
			if slices.Contains(region, next) {
				continue
			}
			er := float64(curr.r+next.r) / 2
			ec := float64(curr.c+next.c) / 2
			edges[FloatPoint{r: er, c: ec}] = FloatPoint{r: er - float64(curr.r), c: ec - float64(curr.c)}
		}
	}
	fmt.Println(edges)
	seen := make(map[FloatPoint]bool)
	sideCount := 0
	for edge, dir := range edges {
		if seen[edge] {
			continue
		}
		seen[edge] = true
		sideCount += 1
		if math.Mod(edge.r, 1) == 0 {
			for _, dr := range []float64{-1, 1} {
				cr := edge.r + dr

				for edges[FloatPoint{r: cr, c: edge.c}] == dir {
					seen[FloatPoint{r: cr, c: edge.c}] = true
					cr += dr
				} // --> Goes into loop ??
			}
		} else {
			for _, dc := range []float64{-1, 1} {
				cc := edge.c + dc

				for edges[FloatPoint{r: edge.r, c: cc}] == dir {
					seen[FloatPoint{r: edge.r, c: cc}] = true
					cc += dc
				}
			}
		}
	}
	return sideCount
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
