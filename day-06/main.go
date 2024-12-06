// AoC 2024, Day 06
// Guard Gallivant
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func parseInput() [][]rune {
	file, err := os.Open("example.txt")
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

type Coordinates struct {
	r int
	c int
}

type Direction struct {
	dr int
	dc int
}

type DirectionController struct {
	all  []Direction
	curr Direction
}

func (d *DirectionController) changeDirection() {
	currIdx := slices.Index(d.all, d.curr)
	d.curr = d.all[(currIdx+1)%4]
}

// global variables
var maxRow, maxCol int
var walkerMap [][]rune

func main() {
	walkerMap = parseInput()
	maxRow, maxCol = len(walkerMap), len(walkerMap[0])

	// Directions in order UP, RIGHT, DOWN, LEFT (Right Rotation)
	// curr Initialised to {0, -1}
	directionController := DirectionController{
		all:  []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
		curr: Direction{-1, 0},
	}

	curr := Coordinates{}

	for y, rowMap := range walkerMap {
		if slices.Contains(rowMap, rune('^')) {
			curr.r, curr.c = y, slices.Index(rowMap, rune('^'))
		}
	}

	walked := make(map[Coordinates]int)
	walker(curr, &walked, directionController)
	fmt.Println(len(walked))
}

func walker(curr Coordinates, walked *map[Coordinates]int, directionController DirectionController) {

	(*walked)[curr] = (*walked)[curr] + 1
	next := Coordinates{
		r: curr.r + directionController.curr.dr,
		c: curr.c + directionController.curr.dc,
	}
	if maxRow <= next.r || next.r < 0 || maxCol <= next.c || next.c < 0 {
		return
	}

	if walkerMap[next.r][next.c] == rune('#') {
		directionController.changeDirection()
	} else {
		curr = next
	}
	walker(curr, walked, directionController)
}
