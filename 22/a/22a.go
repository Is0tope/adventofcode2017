package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

const target = 10000

func hashStr(x, y int) string {
	return "(" + strconv.Itoa(x) + "," + strconv.Itoa(y) + ")"
}

// Need to do this to get cyclical modulo which go doesn't do apparently...
func mod(x, d int) int {
	if x < 0 {
		return d + (x % d)
	}
	return x % d
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	initial := make([][]rune, 0)

	for scanner.Scan() {
		text := scanner.Text()
		initial = append(initial, []rune(text))
	}
	width, height := len(initial[0]), len(initial)
	offsetw, offseth := width/2, height/2

	grid := make(map[string]rune)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid[hashStr(j-offsetw, i-offseth)] = initial[i][j]
		}
	}

	bursts := 0
	x, y := 0, 0
	direction := north

	// how many were infected
	infections := 0

	// main loop
	for bursts < target {
		curr, ok := grid[hashStr(x, y)]
		if !ok {
			// This is a new square, and it is blank
			curr = '.'
		}
		fmt.Printf("burst: %d, direction: %d, infected: %d, curr: %q\n", bursts, direction, infections, curr)

		// turning & settings
		var newnode rune
		if curr == '#' {
			// turn right if infected
			direction = mod(direction+1, 4)
			newnode = '.'
		} else {
			// left otherwise
			direction = mod(direction-1, 4)
			newnode = '#'
			// Increase infection count
			infections++
		}

		// set the current node
		grid[hashStr(x, y)] = newnode

		// move forward
		switch direction {
		case north:
			y--
		case east:
			x++
		case south:
			y++
		case west:
			x--
		}

		// increment bursts
		bursts++
	}

	fmt.Printf("DONE: %d\n", infections)
}
