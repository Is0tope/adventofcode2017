package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type condition struct {
	write int
	move  int
	next  string
}

type state struct {
	conditions map[int]condition
}

func getInitialState(line string) string {
	re := regexp.MustCompile("Begin in state ([A-Z]+)")
	return re.FindStringSubmatch(line)[1]
}

func getTarget(line string) int {
	re := regexp.MustCompile("Perform a diagnostic checksum after (\\d+) steps")
	num, _ := strconv.Atoi(re.FindStringSubmatch(line)[1])
	return num
}

func getState(line string) string {
	re := regexp.MustCompile("In state ([A-Z]+):")
	return re.FindStringSubmatch(line)[1]
}

func getCondition(line string) int {
	re := regexp.MustCompile("If the current value is (\\d+):")
	num, _ := strconv.Atoi(re.FindStringSubmatch(line)[1])
	return num
}

func getWrite(line string) int {
	re := regexp.MustCompile("Write the value (\\d+)")
	num, _ := strconv.Atoi(re.FindStringSubmatch(line)[1])
	return num
}

func getMovement(line string) int {
	re := regexp.MustCompile("Move one slot to the (\\w+)")
	var num int
	if re.FindStringSubmatch(line)[1] == "right" {
		num = 1
	} else {
		num = -1
	}
	return num
}

func getNext(line string) string {
	re := regexp.MustCompile("Continue with state ([A-Z]+)")
	return re.FindStringSubmatch(line)[1]
}

func getTape(tape map[int]int, index int) int {
	ret, ok := tape[index]
	if !ok {
		tape[index] = 0
		ret = 0
	}
	return ret
}

func checksum(tape map[int]int) int {
	cnt := 0
	for _, v := range tape {
		if v == 1 {
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
	initialState := getInitialState(scanner.Text())
	scanner.Scan()
	target := getTarget(scanner.Text())

	// Map of all states & instructions
	program := make(map[string]state)

	// Get all of the instructions
	for scanner.Scan() {
		// skip newline
		scanner.Scan()

		conditions := make(map[int]condition)

		stateID := getState(scanner.Text())

		// 0
		scanner.Scan()
		cond := getCondition(scanner.Text())
		scanner.Scan()
		write := getWrite(scanner.Text())
		scanner.Scan()
		move := getMovement(scanner.Text())
		scanner.Scan()
		next := getNext(scanner.Text())
		conditions[cond] = condition{write, move, next}

		// 1
		scanner.Scan()
		cond = getCondition(scanner.Text())
		scanner.Scan()
		write = getWrite(scanner.Text())
		scanner.Scan()
		move = getMovement(scanner.Text())
		scanner.Scan()
		next = getNext(scanner.Text())
		conditions[cond] = condition{write, move, next}

		program[stateID] = state{conditions}
	}

	pointer := 0
	currentState := initialState
	tape := make(map[int]int)

	// Main loop
	for i := 0; i < target; i++ {
		if i%10000 == 0 {
			fmt.Println(pointer, currentState)
		}
		value := getTape(tape, pointer)
		// get the conditional
		cond := program[currentState].conditions[value]
		// set the value
		tape[pointer] = cond.write
		// move the pointer
		pointer += cond.move
		// set the new state
		currentState = cond.next
	}
	fmt.Printf("DONE: %d\n", checksum(tape))
}
