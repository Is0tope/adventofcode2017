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
const criteriaA = 4
const criteriaB = 8
const target = 5000000

// Global!
var pairCount = 0
var checkedPairs = 0
var isFinished = false

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

func checkerLoop(a chan int, b chan int) {
	for {
		amsg := <-a
		bmsg := <-b
		// fmt.Println("Comparing", amsg, bmsg)
		checkedPairs++
		if checkEquality16b(amsg, bmsg) {
			pairCount++
			// fmt.Println("MATCH", checkedPairs)
		}
		// Quit if number of pairs equal to target has been checked
		if checkedPairs == target {
			// fmt.Println("FINISHED")
			isFinished = true
			break
		}
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	AVal := getCount(scanner.Text())
	scanner.Scan()
	BVal := getCount(scanner.Text())

	chanA := make(chan int, 10000000)
	chanB := make(chan int, 10000000)

	// target := 200000

	// start the thread
	go checkerLoop(chanA, chanB)

	for checkedPairs < target {
		AVal = generate(AVal, genAFactor)
		BVal = generate(BVal, genBFactor)
		if AVal%criteriaA == 0 {
			chanA <- AVal
		}
		if BVal%criteriaB == 0 {
			chanB <- BVal
		}
		// fmt.Println(AVal, BVal, AVal%criteriaA == 0, BVal%criteriaB == 0, pairCount, len(chanA), len(chanB))
	}
	// Need to stall here to let thread finish
	for !isFinished {
		// Do nothing
	}
	fmt.Printf("DONE: %d\n", pairCount)
}
