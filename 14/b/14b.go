package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/spakin/disjoint"
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

func normaliseMappings(m map[int]int) {
	for k, v := range m {
		fmt.Printf("[%d] %d", k, v)
		iv := v
		for iv != m[iv] {
			tmp, ok := m[iv]
			iv = tmp
			fmt.Printf(" -> %d", iv)
			if !ok {
				fmt.Println("failed lookup", iv)
			}
		}
		m[k] = iv
		fmt.Printf("\n")
	}
}

func findLowestMapping(m map[int]int, key int) int {
	val := m[key]
	for val != m[val] {
		val = m[val]
	}
	return val
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()

	groups := [128][128]int{}
	grid := [128][128]rune{}

	// Initialise groups to -1
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			groups[i][j] = -1
		}
	}

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("%s-%d", text, i)
		// fmt.Printf("KEY: %s\n", key)
		hash := knotHash(key, 64)
		// printhash(hash)
		binary := ""
		for _, b := range hash {
			binary = binary + fmt.Sprintf("%08b", b)
		}
		// fmt.Println(binary)
		for j, c := range binary {
			grid[i][j] = c
		}
	}

	currentGroup := 0
	groupMapping := make(map[int]*disjoint.Element)

	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if grid[i][j] == '0' {
				continue
			}
			// ignore top if we are at the top
			top := -1
			if i == 0 {
				top = -1
			} else {
				top = groups[i-1][j]
			}
			// ignore top if we are at the top
			left := -1
			if j == 0 {
				left = -1
			} else {
				left = groups[i][j-1]
			}
			// Only left has group
			if top == -1 && left != -1 {
				groups[i][j] = left
				continue
			}
			// Only top has group
			if top != -1 && left == -1 {
				groups[i][j] = top
				continue
			}
			// Neither has group, so start a new one
			if top == -1 && left == -1 {
				groups[i][j] = currentGroup
				// groupMapping[currentGroup] = currentGroup
				groupMapping[currentGroup] = disjoint.NewElement()
				groupMapping[currentGroup].Data = currentGroup
				currentGroup++
				continue
			}
			// Both have different group, choose lowest & remap
			if top != -1 && left != -1 && top != left {
				disjoint.Union(groupMapping[top], groupMapping[left])
				if top > left {
					groups[i][j] = left
				} else {
					groups[i][j] = top
				}
				continue
			}
			// Both are the same
			if top != -1 && left == top {
				groups[i][j] = left
				continue
			}
		}
	}

	// Follow the links all the way through
	// fmt.Println(groupMapping)
	distinctGroups := make(map[int]bool)
	for _, v := range groupMapping {
		distinctGroups[v.Find().Data.(int)] = true
	}
	fmt.Printf("DONE: %d\n", len(distinctGroups))
}
