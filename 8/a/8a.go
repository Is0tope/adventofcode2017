package main

import (

	// "strings"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// nicked from SO lol
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	registers := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		register := tokens[0]
		inst := tokens[1]
		val, _ := strconv.Atoi(tokens[2])
		target := tokens[4]
		comp := tokens[5]
		compval, _ := strconv.Atoi(tokens[6])

		// Add non seen registers
		if _, ok := registers[register]; !ok {
			registers[register] = 0
		}
		if _, ok := registers[target]; !ok {
			registers[target] = 0
		}

		isCond := false
		switch comp {
		case ">":
			isCond = registers[target] > compval
		case "<":
			isCond = registers[target] < compval
		case ">=":
			isCond = registers[target] >= compval
		case "<=":
			isCond = registers[target] <= compval
		case "==":
			isCond = registers[target] == compval
		case "!=":
			isCond = registers[target] != compval
		}

		if isCond {
			switch inst {
			case "inc":
				registers[register] += val
			case "dec":
				registers[register] -= val
			}
		}
		fmt.Println(registers)
	}
	mxv := MinInt
	for _, v := range registers {
		if v > mxv {
			mxv = v
		}
	}
	fmt.Printf("DONE: %d\n", mxv)
}
