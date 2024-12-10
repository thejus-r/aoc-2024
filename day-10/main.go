package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
	stack "github.com/thejus-r/aoc-2024/utils/ds"
)

type Point struct {
	r int
	c int
}

var topographicMap [][]int
var trailHeads []Point
var max Point

func main() {
	topographicMap = parseInput("input.txt")
	trailHeads = findAllTrailHeads(topographicMap)
	max.r, max.c = len(topographicMap), len(topographicMap[0])

	sumOfScore := 0
	for _, trailHead := range trailHeads {
		sumOfScore += findTrailScore(trailHead, false)
	}
	fmt.Println(sumOfScore)
}

func findTrailScore(h Point, partOne bool) int {
	s := stack.New[Point]()

	score := 0
	visited := make(map[Point]bool)
	s.Push(h)
	for {
		if s.IsEmpty() {
			break
		}
		curr, _ := s.Pop()
		for _, next := range []Point{{curr.r - 1, curr.c}, {curr.r, curr.c - 1}, {curr.r + 1, curr.c}, {curr.r, curr.c + 1}} {
			if isValid(next) {
				if topographicMap[next.r][next.c] != topographicMap[curr.r][curr.c]+1 {
					continue
				}
				if visited[next] && partOne {
					continue
				}
				visited[next] = true
				if topographicMap[next.r][next.c] == 9 {
					score += 1
				} else {
					s.Push(next)
				}
			}
		}
	}
	return score
}

func isValid(p Point) bool {
	return p.r >= 0 && p.c >= 0 && p.r < max.r && p.c < max.c
}

// Gets all trail locations
func findAllTrailHeads(topographicMap [][]int) []Point {
	trailHeads := []Point{}
	for r := 0; r < len(topographicMap); r++ {
		for c := 0; c < len(topographicMap[0]); c++ {
			if topographicMap[r][c] == 0 {
				trailHeads = append(trailHeads, Point{r, c})
			}
		}
	}
	return trailHeads
}

func parseInput(fileName string) [][]int {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)

	lines := strings.Split(string(dat), "\n")

	result := make([][]int, len(lines))
	for y, line := range lines {
		charArr := strings.Split(line, "")
		temp := make([]int, len(charArr))
		for x, char := range charArr {
			temp[x] = utils.MustAtoi(char)
		}
		result[y] = temp
	}

	return result
}
