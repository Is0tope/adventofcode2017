package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	line,_ := ioutil.ReadFile("../input1.txt")
	
	sum := 0
	length := len(line)
	loffset := length / 2

	for i,c := range line {
		if c == line[(i+loffset) % length] {
			num,_ := strconv.Atoi(string(c))
			sum += num
		}
	}
	fmt.Println(sum)
}
