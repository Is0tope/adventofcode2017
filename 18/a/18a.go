package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func resolve(m map[string]int, str string) (int, error) {
	// Is it a variable?
	matched, _ := regexp.MatchString("[a-z]+", str)
	if matched {
		v, ok := m[str]
		if !ok {
			return 0, fmt.Errorf("No such key")
		}
		return v, nil
	}
	return strconv.Atoi(str)
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := make([][]string, 0)

	lastSound := -1
	lastRcv := -1
	registers := make(map[string]int)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, tokens)
	}

	for i := 0; i >= 0 && i < len(instructions); i++ {
		inst := instructions[i]
		opcode := inst[0]
		switch opcode {
		case "snd":
			x, _ := resolve(registers, inst[1])
			lastSound = x
		case "set":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] = y
		case "add":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] += y
		case "mul":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] *= y
		case "mod":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] = registers[inst[1]] % y
		case "rcv":
			x, _ := resolve(registers, inst[1])
			if x > 0 {
				lastRcv = lastSound
				fmt.Printf("DONE: %d\n", lastRcv)
				return
			}
		case "jgz":
			x, _ := resolve(registers, inst[1])
			y, _ := resolve(registers, inst[2])
			if x > 0 {
				i += (y - 1)
			}
		}
		fmt.Println(i, inst, registers, lastSound, lastRcv)
	}
}
