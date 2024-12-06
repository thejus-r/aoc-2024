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

type position struct {
	x int
	y int
}

type dir struct {
	dx int
	dy int
}

type Directions struct {
	all  []dir
	curr dir
}

func (d *Directions) changeDirection() {
	currIdx := slices.Index(d.all, d.curr)
	d.curr = d.all[(currIdx+1)%4]
}
func main() {

	walkerMap := parseInput()
	maxRow, maxCol := len(walkerMap), len(walkerMap[0])

	// Directions in order UP, RIGHT, DOWN, LEFT (Right Rotation)
	// curr Initialised to {0, -1}
	directions := Directions{
		all:  []dir{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
		curr: dir{0, -1},
	}

	walks := 0
	curr := position{}
	objs := []position{}

	for y, rowMap := range walkerMap {
		if slices.Contains(rowMap, rune('^')) {
			curr.y = y
			curr.x = slices.Index(rowMap, rune('^'))
		}
		if slices.Contains(rowMap, rune('#')) {
			for x, rowItem := range rowMap {
				if rowItem == rune('#') {
					objs = append(objs, position{x, y})
				}
			}
		}
	}

	walker(maxRow, maxCol, &walkerMap, curr, directions, objs, &walks)

	fmt.Println(walks)
}

func walker(maxRow int, maxCol int, walkerMap *[][]rune, curr position, directions Directions, objs []position, walks *int) {

	next := position{
		x: curr.x + directions.curr.dx,
		y: curr.y + directions.curr.dy,
	}

	if maxRow <= next.x || next.x < 0 || maxCol <= next.y || next.y < 0 {
		return
	}

	if slices.Contains(objs, next) {
		directions.changeDirection()
	} else {
		curr = next
		if (*walkerMap)[next.x][next.y] != rune('X') {
			*walks++
		}
		(*walkerMap)[curr.x][curr.y] = rune('X')

	}
	walker(maxRow, maxCol, walkerMap, curr, directions, objs, walks)
}
