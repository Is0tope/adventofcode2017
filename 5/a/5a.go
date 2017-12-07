package main

import (
	"fmt"
	// "strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	file,_ := os.Open("../input.txt")
	defer file.Close()
	code := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num,_ := strconv.Atoi(scanner.Text())
		code = append(code,num)
	}
	pos := 0
	counter := 0
	for pos < len(code) {
		offset := code[pos]
		code[pos]++
		pos += offset
		counter++
	}
	fmt.Println("DONE: %d",counter)
}
