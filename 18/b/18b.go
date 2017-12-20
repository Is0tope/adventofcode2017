package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

func program(instructions [][]string, label int, send chan int, recv chan int) {
	numSent := 0
	registers := make(map[string]int)
	// add label to p
	registers["p"] = label

	for i := 0; i >= 0 && i < len(instructions); i++ {
		inst := instructions[i]
		opcode := inst[0]
		switch opcode {
		case "snd":
			x, _ := resolve(registers, inst[1])
			send <- x
			numSent++
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
			val := <-recv
			registers[inst[1]] = val
		case "jgz":
			x, _ := resolve(registers, inst[1])
			y, _ := resolve(registers, inst[2])
			if x > 0 {
				i += (y - 1)
			}
		}
		fmt.Println(label, i, inst, registers, numSent)
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := make([][]string, 0)

	// channels
	// must be buffered to prevent block
	chan0 := make(chan int, 1000000)
	chan1 := make(chan int, 1000000)

	// waitgroup
	var wg sync.WaitGroup

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, tokens)
	}

	wg.Add(2)
	go program(instructions, 0, chan1, chan0)
	go program(instructions, 1, chan0, chan1)

	// Wait...
	wg.Wait()
	fmt.Println("DONE")

}
