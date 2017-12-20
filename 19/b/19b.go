package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

func isValidField(c rune) bool {
	matched, _ := regexp.MatchString("[A-Z|\\-+]", string(c))
	return matched
}

func isLetter(c rune) bool {
	matched, _ := regexp.MatchString("[A-Z]", string(c))
	return matched
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make([][]rune, 0)
	letterOrder := ""

	// Set up the grid
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	// Find the starting point
	startPos := -1
	for i, c := range grid[0] {
		if c == '|' {
			startPos = i
			break
		}
	}

	// Start searching
	currentDir := south
	x, y := startPos, 0
	steps := 0

loop:
	for {
		char := grid[y][x]
		steps++

		fmt.Printf("(%d,%d) [%q] %d\n", x, y, char, currentDir)

		// Is this a letter?
		if isLetter(char) {
			letterOrder += string(char)
			// Is this the final one?
			switch currentDir {
			case north:
				if y > 0 {
					if !isValidField(grid[y-1][x]) {
						break loop
					}
				} else {
					break loop
				}
			case south:
				if y < len(grid)-1 {
					if !isValidField(grid[y+1][x]) {
						break loop
					}
				} else {
					break loop
				}
			case east:
				if x < len(grid[0])-1 {
					if !isValidField(grid[y][x+1]) {
						break loop
					}
				} else {
					break loop
				}
			case west:
				if x > 0 {
					if !isValidField(grid[y][x-1]) {
						break loop
					}
				} else {
					break loop
				}
			}
		}

		if char == '+' {
			if currentDir == north || currentDir == south {
				// check east
				if x < len(grid[0])-1 {
					if isValidField(grid[y][x+1]) {
						currentDir = east
					}
				}

				// check west
				if x > 0 {
					if isValidField(grid[y][x-1]) {
						currentDir = west
					}
				}
			} else {
				// check north
				if y > 0 {
					if isValidField(grid[y-1][x]) {
						currentDir = north
					}
				}

				// check south
				if y < len(grid)-1 {
					if isValidField(grid[y+1][x]) {
						currentDir = south
					}
				}
			}
		}
		// keep going
		switch currentDir {
		case north:
			y--
		case south:
			y++
		case east:
			x++
		case west:
			x--
		}

	}
	fmt.Printf("DONE: %s (%d)\n", letterOrder, steps)
}
