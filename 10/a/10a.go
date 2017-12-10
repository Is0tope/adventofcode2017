package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const ringSize int = 256

func reverse(list []int, start int, off int) {
	ll := len(list)
	span := 0
	end := (start + off - 1) % ll
	span = int(math.Ceil(float64(off) / 2))
	// fmt.Println(span)
	for i := 0; i < span; i++ {
		// fmt.Println(i)
		// get end
		ecoord := (end - i) % ll
		if ecoord < 0 {
			ecoord = ll + ecoord
		}
		scoord := (start + i) % ll
		fmt.Printf("s: %d, e: %d\n", scoord, ecoord)
		tmp := list[ecoord]
		list[ecoord] = list[scoord]
		list[scoord] = tmp
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")
	lengths := make([]int, len(tokens))
	for i, t := range tokens {
		lengths[i], _ = strconv.Atoi(t)
	}
	ring := make([]int, ringSize)
	for i := range ring {
		ring[i] = i
	}

	curPos := 0
	skip := 0

	for _, l := range lengths {
		// reverse
		fmt.Println(curPos, l)
		reverse(ring, curPos, l)
		fmt.Println(ring)
		// jump forward
		offset := l + skip
		curPos = (curPos + offset) % ringSize
		skip++
	}
	// fmt.Printf("DONE: %s\n", test)
}
