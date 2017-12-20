package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

const year = 2018

func printRing(r *ring.Ring) {
	buf := r
	for i := 0; i < r.Len(); i++ {
		fmt.Printf("%d -> ", buf.Value)
		buf = buf.Next()
	}
	fmt.Printf("\n")
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	skip, _ := strconv.Atoi(scanner.Text())

	buffer := ring.New(1)
	// initial value
	buffer.Value = 0
	counter := 1
	fmt.Printf("counter: %d, buffer: %d, len: %d\n", counter, buffer.Value, buffer.Len())
	printRing(buffer)

	for counter < year {
		// skip forward
		buffer = buffer.Move(skip)
		// insert new element
		elem := ring.New(1)
		elem.Value = counter
		buffer.Link(elem)
		buffer = buffer.Next()
		fmt.Printf("counter: %d, buffer: %d, len: %d\n", counter, buffer.Value, buffer.Len())
		// printRing(buffer)
		// increment counter
		counter++
	}

	fmt.Printf("DONE: %d\n", buffer.Next().Value)
}
