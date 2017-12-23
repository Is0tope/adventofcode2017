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
	registers := make(map[string]int)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, tokens)
	}

	// initialise all registers (can do via cache miss but this is more formal)
	for _, c := range "abcdefgh" {
		registers[string(c)] = 0
	}

	// number of muls
	mulCount := 0

	pointer := 0
	for pointer >= 0 && pointer < len(instructions) {
		inst := instructions[pointer]
		opcode := inst[0]
		switch opcode {
		case "set":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] = y
		case "sub":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] -= y
		case "mul":
			y, _ := resolve(registers, inst[2])
			registers[inst[1]] *= y
			mulCount++
		case "jnz":
			x, _ := resolve(registers, inst[1])
			y, _ := resolve(registers, inst[2])
			if x != 0 {
				pointer += (y - 1)
			}
		}
		pointer++
		fmt.Println(pointer, inst, registers, mulCount)
	}
}
