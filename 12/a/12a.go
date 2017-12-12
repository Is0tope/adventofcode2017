package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const startNode = 0

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	mapping := make(map[int][]int)
	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " <-> ")
		root, _ := strconv.Atoi(tokens[0])
		children := make([]int, 0)
		childTokens := strings.Split(tokens[1], ",")
		for _, c := range childTokens {
			child, _ := strconv.Atoi(strings.Trim(c, " "))
			children = append(children, child)
		}
		mapping[root] = children
	}

	// Find the 0 network
	zeroNetwork := make(map[int]bool)
	queue := make(chan int, 99999)

	// Add initial item
	queue <- startNode
	for 0 < len(queue) {
		// get item
		node := <-queue
		fmt.Println(node)
		// Is it in the network?
		if _, ok := zeroNetwork[node]; ok {
			continue
		}
		// Add it to the network
		zeroNetwork[node] = true
		// get the children
		children, _ := mapping[node]
		// Push children to the queue
		for _, c := range children {
			queue <- c
		}
	}
	// fmt.Println(zeroNetwork)
	fmt.Printf("DONE: %d\n", len(zeroNetwork))
}
