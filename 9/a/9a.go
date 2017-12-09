package main

import (
	"bufio"
	"fmt"
	"os"
)

func numGroups(str string) int {
	// gCount := 0
	gScore := 0
	gValue := 0
	isGarbage := false

	for i := 0; i < len(str); i++ {
		c := str[i]
		fmt.Printf("%c", c)
		// ignore !
		if c == '!' && isGarbage {
			i++
			continue
		}
		// garbage
		if c == '<' && !isGarbage {
			isGarbage = true
			continue
		}
		if c == '>' && isGarbage {
			isGarbage = false
			continue
		}
		// if still garbage by this point, skip
		if isGarbage {
			continue
		}

		// Actual groups
		if c == '{' {
			gScore++
			continue
		}
		if c == '}' {
			gValue += gScore
			gScore--
			continue
		}
	}
	fmt.Printf(": %d\n", gValue)
	return 0
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		numGroups(text)
	}

	// fmt.Printf("DONE: %s\n", test)
}
