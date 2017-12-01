package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	line,_ := ioutil.ReadFile("../input1.txt")
	// add last character to end
	line = append(line,line[len(line)-1])	
	
	sum := 0
	lastVal := string(line[0])

	for i := 1; i < len(line); i++ {
		ch := string(line[i])
		if ch == lastVal {
			num,_ := strconv.Atoi(ch)
			sum += num
		}
		lastVal = ch
	}
	fmt.Println(sum)
}
