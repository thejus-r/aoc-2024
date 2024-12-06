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

// Directions in order UP, RIGHT, DOWN, LEFT (Right Rotation)
// curr Initialised to {0, -1}
func (d *DirectionController) init() {
	d.all = []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	d.curr = Direction{-1, 0}
}

func (d *DirectionController) changeDirection() {
	currIdx := slices.Index(d.all, d.curr)
	d.curr = d.all[(currIdx+1)%4]
}

// global variables
var maxRow, maxCol int
var walkerMap [][]rune
var startCoordinate Coordinates
var directionController DirectionController

func main() {
	walkerMap = parseInput()
	maxRow, maxCol = len(walkerMap), len(walkerMap[0])

	directionController.init()

	for y, rowMap := range walkerMap {
		if slices.Contains(rowMap, rune('^')) {
			startCoordinate.r, startCoordinate.c = y, slices.Index(rowMap, rune('^'))
		}
	}

	curr := startCoordinate

	walked := make(map[Coordinates]int)
	walker(curr, &walked, directionController)

	diffrentCoordinatesForObstacle()

	fmt.Printf("Distinct positions will the guard visits: %d\n", len(walked))
}

func walker(curr Coordinates, walked *map[Coordinates]int, directionController DirectionController) {

	// add current to the map
	(*walked)[curr] = (*walked)[curr] + 1
	next := Coordinates{
		r: curr.r + directionController.curr.dr,
		c: curr.c + directionController.curr.dc,
	}
	if maxRow <= next.r || next.r < 0 || maxCol <= next.c || next.c < 0 {
		return // out of bound
	}

	// obstacle found
	if walkerMap[next.r][next.c] == rune('#') {
		directionController.changeDirection()
	} else {
		curr = next
	}
	walker(curr, walked, directionController)
}

func diffrentCoordinatesForObstacle() {

	validPositions := 0

	// iterate through the map to get all the possible obstacle positions
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if walkerMap[r][c] == rune('#') {
				continue
			}
			if checkForLoop(Coordinates{r: r, c: c}) {
				validPositions++
			}
		}
	}

	fmt.Printf("Different positions could you choose for this obstruction: %d\n", validPositions)
}

func checkForLoop(obstacle Coordinates) bool {
	guard := Coordinates{r: startCoordinate.r, c: startCoordinate.c}
	directionController.init()

	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d, %d, %d, %d", guard.r, guard.c, directionController.curr.dc, directionController.curr.dr)] = true

	for {
		next := Coordinates{
			r: guard.r + directionController.curr.dr,
			c: guard.c + directionController.curr.dc,
		}
		if maxRow <= next.r || next.r < 0 || maxCol <= next.c || next.c < 0 {
			return false // out of bound
		}
		nextCell := rune('#')
		if !(next.r == obstacle.r && next.c == obstacle.c) {
			nextCell = walkerMap[next.r][next.c]
		}

		if nextCell == rune('#') {
			directionController.changeDirection() // obstacle detected
		} else {
			guard = next
		}

		state := fmt.Sprintf("%d, %d, %d, %d", guard.r, guard.c, directionController.curr.dc, directionController.curr.dr)
		if visited[state] {
			return true // it loops
		}
		visited[state] = true
	}
}
