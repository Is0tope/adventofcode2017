package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFirstKey(m map[int]bool) (int, error) {
	for k := range m {
		return k, nil
	}
	return -1, fmt.Errorf("No more keys")
}

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
	network := make(map[int]bool)
	// which nodes are left
	remainingNodes := make(map[int]bool)
	for n := range mapping {
		remainingNodes[n] = true
	}
	// number of groups
	numGroups := 0

	queue := make(chan int, 99999)

	for 0 < len(remainingNodes) {
		firstNode, _ := getFirstKey(remainingNodes)

		queue <- firstNode
		numGroups++

		for 0 < len(queue) {
			// get item
			node := <-queue
			// fmt.Println(node)
			// Is it in the network?
			if _, ok := network[node]; ok {
				continue
			}
			// Add it to the network
			network[node] = true
			// remove it from remaining nodes
			delete(remainingNodes, node)
			// get the children
			children, _ := mapping[node]
			// Push children to the queue
			for _, c := range children {
				queue <- c
			}
		}
		fmt.Printf("Finished with group %d\n", numGroups)
	}
	fmt.Printf("DONE: %d\n", numGroups)
}
