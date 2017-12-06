package main

import (
	"fmt"
	// "strconv"
	"strings"
	"bufio"
	"os"
)

func main() {
	file,_ := os.Open("../input.txt")
	defer file.Close()
	
	counter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text()," ")
		m := make(map[string]bool)
		flag := false
		for _,t := range tokens {
			_,ok := m[t]
			if ok {
				flag = true
				break
			}
			m[t] = true
		}
		if !flag {
			counter++
		}
	}
	fmt.Printf("DONE: %d\n",counter)
}
