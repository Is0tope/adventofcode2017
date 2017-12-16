package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPeriod(n int) int {
	return (2 * n) - 2
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	severity := 0

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ": ")
		time, _ := strconv.Atoi(tokens[0])
		dist, _ := strconv.Atoi(tokens[1])
		period := getPeriod(dist)

		if time%period == 0 {
			severity += time * dist
			fmt.Printf("COLLISION: time: %d, dist: %d\n", time, dist)
		}
	}

	fmt.Printf("DONE: %d\n", severity)
}
