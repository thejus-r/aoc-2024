package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

type Vec struct {
	x int
	y int
}

func convertSeqToVec(s rune) Vec {
	switch s {
	case '<':
		return Vec{x: -1, y: 0}
	case '^':
		return Vec{x: 0, y: -1}
	case '>':
		return Vec{x: 1, y: 0}
	case 'v':
		return Vec{x: 0, y: 1}
	default:
		return Vec{x: 0, y: 0}
	}
}

var floorMap [][]rune
var sequence []rune
var robot Vec

func main() {
	parseInput("input.txt")
	simulate()
	showState()
	findPart1()
}

func findPart1() {

	sum := 0

	for i := 0; i < len(floorMap); i++ {
		for j := 0; j < len(floorMap[0]); j++ {
			if floorMap[i][j] == 'O' {
				sum = sum + (100 * i) + j
			}
		}
	}

	fmt.Println(sum)
}

func simulate() {
	for _, seq := range sequence {
		m := convertSeqToVec(seq)
		n := Vec{robot.x + m.x, robot.y + m.y}
		if floorMap[n.y][n.x] == rune('#') {
			continue
		}
		if floorMap[n.y][n.x] == rune('O') {
			if !moveBlock(n, m) {
				continue
			}
		}
		robot = n

	}
}

func moveBlock(n, m Vec) bool {
	nn := Vec{x: n.x + m.x, y: n.y + m.y}
	if floorMap[nn.y][nn.x] == rune('#') {
		return false
	}
	if floorMap[nn.y][nn.x] == rune('.') {
		floorMap[n.y][n.x] = '.'
		floorMap[nn.y][nn.x] = 'O'
		return true
	}
	if floorMap[nn.y][nn.x] == rune('O') {
		if moveBlock(nn, m) {
			floorMap[n.y][n.x] = '.'
			floorMap[nn.y][nn.x] = 'O'
		} else {
			return false
		}
	}
	return true
}

func showState() {
	for i := 0; i < len(floorMap); i++ {
		for j := 0; j < len(floorMap[0]); j++ {
			if i == robot.y && j == robot.x {

				fmt.Printf("%c ", '@')
			} else {

				fmt.Printf("%c ", floorMap[i][j])
			}
		}
		fmt.Println()
	}
}

func parseInput(fileName string) {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)

	rawStr := string(dat)

	parts := strings.Split(rawStr, "\n\n")

	floorMap = [][]rune{}

	for y, row := range strings.Split(parts[0], "\n") {
		temp := []rune{}
		for x, el := range row {
			if el == rune('@') {
				robot = Vec{x: x, y: y}
				el = rune('.')
			}
			temp = append(temp, el)
		}
		floorMap = append(floorMap, temp)
	}

	sequence = []rune{}
	for _, seq := range parts[1] {
		sequence = append(sequence, seq)
	}
}
