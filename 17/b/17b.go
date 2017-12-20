package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const year = 50000000

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	skip, _ := strconv.Atoi(scanner.Text())
	counter := 1
	offset := 0
	length := 1
	lastZero := -1
	for counter < year {
		offset = (offset + skip + 1) % length
		if offset == 0 {
			lastZero = counter
			fmt.Printf("0 insert found at %d\n", counter)
		}
		length++
		counter++
	}

	fmt.Printf("DONE: %d\n", lastZero)
}
