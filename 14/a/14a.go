package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const ringSize int = 256

var suffix = []int{17, 31, 73, 47, 23}

func reverse(list []int, start int, off int) {
	ll := len(list)
	span := 0
	end := (start + off - 1) % ll
	span = int(math.Ceil(float64(off) / 2))
	for i := 0; i < span; i++ {
		// get end
		ecoord := (end - i) % ll
		if ecoord < 0 {
			ecoord = ll + ecoord
		}
		scoord := (start + i) % ll
		tmp := list[ecoord]
		list[ecoord] = list[scoord]
		list[scoord] = tmp
	}
}

func dense(list []int) []int {
	l := 16
	off := 16
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		char := list[i*16]
		for j := 1; j < off; j++ {
			char = char ^ list[(i*16)+j]
		}
		ret[i] = char
	}
	return ret
}

func knotHash(text string, repeat int) []int {
	lengths := make([]int, len(text))
	for i, t := range text {
		lengths[i] = int(t)
	}
	for _, v := range suffix {
		lengths = append(lengths, v)
	}
	ring := make([]int, ringSize)
	for i := range ring {
		ring[i] = i
	}

	curPos := 0
	skip := 0
	for i := 0; i < repeat; i++ {
		for _, l := range lengths {
			// reverse
			reverse(ring, curPos, l)
			// jump forward
			offset := l + skip
			curPos = (curPos + offset) % ringSize
			skip++
		}
	}
	return dense(ring)
}

func printhash(hash []int) {
	for _, v := range hash {
		char := strconv.FormatInt(int64(v), 16)
		if len(char) < 2 {
			char = "0" + char
		}
		fmt.Printf("%s", char)
	}
	fmt.Printf("\n")
}

func countOnes(text string) int {
	cnt := 0
	for _, s := range text {
		if s == '1' {
			cnt++
		}
	}
	return cnt
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()
	total := 0

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("%s-%d", text, i)
		// fmt.Printf("KEY: %s\n", key)
		hash := knotHash(key, 64)
		// printhash(hash)
		binary := ""
		for _, b := range hash {
			binary = binary + fmt.Sprintf("%08b", b)
		}
		fmt.Println(binary)
		total += countOnes(binary)
	}
	fmt.Printf("DONE: %d\n", total)
}
