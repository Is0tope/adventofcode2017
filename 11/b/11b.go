package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func cubeDist(x, y, z, x1, y1, z1 int) int {
	return int((math.Abs(float64(x1-x)) + math.Abs(float64(y1-y)) + math.Abs(float64(z1-z))) / 2)
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	tokens := strings.Split(text, ",")

	maxDist := 0

	x, y, z := 0, 0, 0
	for _, t := range tokens {
		switch t {
		case "n":
			y++
			z--
		case "ne":
			z--
			x++
		case "se":
			x++
			y--
		case "s":
			z++
			y--
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		}
		fmt.Println(x, y, z)
		dist := cubeDist(0, 0, 0, x, y, z)
		if dist > maxDist {
			maxDist = dist
		}
	}
	fmt.Printf("DONE: %d\n", maxDist)

	// fmt.Printf("DONE: %s\n", test)
}
