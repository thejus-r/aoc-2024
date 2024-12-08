package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

type Point struct {
	x int
	y int
}

var maxRow int
var maxCol int
var noOfAntenna int

var cityMap [][]string
var antennas map[string][]Point
var antiNodes map[Point]bool
var resonantNodes map[Point]bool

func main() {

	parseInput("input.txt")
	maxRow, maxCol = len(cityMap), len(cityMap[0])
	antiNodes = make(map[Point]bool)
	resonantNodes = make(map[Point]bool)
	temp := cityMap
	findAllAntiNode()
	cityMap = temp
	findAllResonantNode()
	fmt.Printf("No of Antinodes: %d\n", len(antiNodes))
	fmt.Printf("No of Resonant Nodes: %d\n", len(resonantNodes)+noOfAntenna)
	showMap()
}

func isValidPoint(point Point) bool {
	return point.y >= 0 && point.y < maxRow && point.x >= 0 && point.x < maxCol
}

func isValidPair(p1 Point, p2 Point) bool {
	return 2 <= utils.Abs(p1.y-p2.y)+utils.Abs(p1.x-p2.x)
}

// Part 1
func findAntiNodes(p1 Point, p2 Point) {

	dx, dy := p1.x-p2.x, p1.y-p2.y
	a1, a2 := Point{x: p1.x + dx, y: p1.y + dy}, Point{x: p2.x - dx, y: p2.y - dy}

	if !antiNodes[a1] && isValidPoint(a1) {
		antiNodes[a1] = true
	}

	if !antiNodes[a2] && isValidPoint(a2) {
		antiNodes[a2] = true
	}
}

func findAllAntiNode() {
	for _, nodes := range antennas {
		for i := 0; i < len(nodes); i++ {
			for j := 0; j < i; j++ {
				p1 := Point{x: nodes[j].x, y: nodes[j].y}
				p2 := Point{x: nodes[i].x, y: nodes[i].y}
				if isValidPair(p1, p2) {
					findAntiNodes(p1, p2)
				}
			}
		}
	}
}

// Part 2
func findResonantNodes(p1 Point, p2 Point) {

	dx, dy := p1.x-p2.x, p1.y-p2.y
	var a1 Point

	for isValidPoint(p1) {
		a1 = Point{x: p1.x + dx, y: p1.y + dy}
		if !resonantNodes[a1] && isValidPoint(a1) && cityMap[a1.y][a1.x] == "." {
			cityMap[a1.y][a1.x] = "#"
			resonantNodes[a1] = true
		}
		p1 = a1
	}
	p1 = p2
	for isValidPoint(p1) {
		a1 = Point{x: p1.x - dx, y: p1.y - dy}
		if !resonantNodes[a1] && isValidPoint(a1) && cityMap[a1.y][a1.x] == "." {
			cityMap[a1.y][a1.x] = "#"
			resonantNodes[a1] = true
		}
		p1 = a1
	}
}

func findAllResonantNode() {
	for _, nodes := range antennas {
		for i := 0; i < len(nodes); i++ {
			for j := 0; j < i; j++ {
				p1 := Point{x: nodes[j].x, y: nodes[j].y}
				p2 := Point{x: nodes[i].x, y: nodes[i].y}
				if isValidPair(p1, p2) {
					findResonantNodes(p1, p2)
				}
			}
		}
	}
}

func showMap() {
	for y, row := range cityMap {
		for x := range row {
			fmt.Printf("%s ", cityMap[y][x])
		}
		fmt.Println()
	}
}

func parseInput(fileName string) {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)
	// String manupulations

	antennas = make(map[string][]Point)

	input := string(dat)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		els := strings.Split(line, "")
		cityMap = append(cityMap, els)
		for j, el := range els {
			if el != "." {
				noOfAntenna++
				antennas[el] = append(antennas[el], Point{x: j, y: i})
			}
		}
	}
}
