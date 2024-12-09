// AoC 2024, Day 09
// Disk Fragmenter

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thejus-r/aoc-2024/utils"
)

//2333133121414131402121243

var diskMap []int
var fileMap map[int]int
var freeMap map[int]int
var blocks []int
var checkSum int

func main() {
	blocks = make([]int, 0)
	fileMap = make(map[int]int)
	freeMap = make(map[int]int)
	diskMap = readDiskMap("input.txt")
	checkSum = 0
	// Select Part
	part := 2

	createDisk()
	showDisk()

	// Part 1
	if part == 1 {
		defragDisk()
	}
	if part == 2 {
		betterDefragDisk()
	}
	// Part 2

	showDisk()

	findCheckSum()

	fmt.Printf("The checksum of the disk is: %d\n", checkSum)
}

// Find till index, isAvailable, Index
func getFreeSpace(sizeNeeded, till int) (bool, int) {
	size := 0
	left := 0
	for {
		if left > till {
			return false, -1
		}
		if blocks[left] == -1 {
			i := left
			for {
				if i > till {
					break
				}
				if blocks[i] != -1 {
					break
				}
				size++
				i++
			}
			if size < sizeNeeded {
				left = left + size
				size = 0
			} else {
				return true, left
			}
		}
		left++
	}
}

func swapFile(freeHead, fileHead, size int) {
	for idx := 0; idx < size; idx++ {
		blocks[idx+freeHead] = blocks[idx+fileHead+1]
		blocks[idx+fileHead+1] = -1
	}
}

func betterDefragDisk() {
	right := len(blocks) - 1

	for {
		if right < 1 {
			break
		}
		for {
			if blocks[right] != -1 {
				break
			}
			right--
		}
		sizeOfFile := fileMap[blocks[right]]

		ok, idx := getFreeSpace(sizeOfFile, right)
		if ok {
			swapFile(idx, right-fileMap[blocks[right]], sizeOfFile)
		}
		right--
	}
}

func defragDisk() {
	left := 0
	right := len(blocks) - 1
	for {
		if left < right {
			for left < len(blocks) && blocks[left] != -1 {
				left++
			}
			for right >= 0 && blocks[right] == -1 {
				right--
			}
			if left < right {
				blocks[left], blocks[right] = blocks[right], -1
				left++
				right--
			}
		} else {
			break
		}
	}

}

func createDisk() {
	fileId := 0
	for idx, el := range diskMap {
		if idx%2 == 0 {
			for j := 0; j < el; j++ {
				blocks = append(blocks, fileId)
			}
			fileMap[fileId] = el
			fileId++
		} else {
			for j := 0; j < el; j++ {
				blocks = append(blocks, -1)
			}
			freeMap[fileId] = el
		}
	}
}

func findCheckSum() {
	for idx, block := range blocks {
		if block != -1 {
			checkSum = checkSum + idx*block
		}
	}
}

func showDisk() {
	for _, block := range blocks {
		if block != -1 {
			fmt.Printf("%d", block)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}

func readDiskMap(fileName string) []int {
	dat, err := os.ReadFile(fileName)
	utils.Check(err)

	result := []int{}
	inputString := string(dat)
	inputCharArr := strings.Split(inputString, "")
	for _, char := range inputCharArr {
		result = append(result, utils.MustAtoi(char))
	}
	return result
}
