package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const genAFactor = 16807
const genBFactor = 48271
const divisor = 2147483647
const hex16b = 65535 // 0xFFFF (16 bits, all one)

func generate(start int, factor int) int {
	return (start * factor) % divisor
}

func getCount(text string) int {
	tokens := strings.Split(text, " ")
	ret, _ := strconv.Atoi(tokens[len(tokens)-1])
	return ret
}

func checkEquality16b(a int, b int) bool {
	// And both with 0xFFFF and check for equality
	return (hex16b & a) == (hex16b & b)
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	AVal := getCount(scanner.Text())
	scanner.Scan()
	BVal := getCount(scanner.Text())

	pairCount := 0

	target := 40000000
	for counter := 0; counter < target; counter++ {
		// Check for equality
		if checkEquality16b(AVal, BVal) {
			pairCount++
		}
		AVal = generate(AVal, genAFactor)
		BVal = generate(BVal, genBFactor)
	}

	fmt.Printf("DONE: %d\n", pairCount)
}
