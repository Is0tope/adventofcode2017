package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func max(arr []int) (int,int) {
	mx := 0
	pos:= 0
	for i,v := range arr {
		if v > mx {
			mx = v
			pos = i
		}
	}
	return pos,mx
}

func blocks2str(arr []int) string {
	ret := []string{}
	for _,a := range arr {
		ret = append(ret,strconv.Itoa(a))
	}
	return strings.Join(ret," ")
}

func main() {
	file,_ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// only one line
	scanner.Scan()
	tokens := strings.Split(scanner.Text(),"\t")

	blocks := []int{}
	seen := make(map[string]bool)
	counter := 0
	targetstate := ""
	cyclestart := 0
	incycle := false 
	// Add the first one
	seen[blocks2str(blocks)] = true

	for _,t := range tokens {
		num,_ := strconv.Atoi(t)
		blocks = append(blocks,num)
	}
	for {
		counter++
		pos,mx := max(blocks)
		blocks[pos] = 0
		for i:=1;i<mx+1;i++ {
			blocks[(pos+i)%len(blocks)]++
		}
		// fmt.Println(blocks)
		strv := blocks2str(blocks)
		if strv == targetstate {
			break
		}
		if _,ok := seen[strv]; ok {
			if !incycle {
				incycle = true
				targetstate = strv
				cyclestart = counter
				fmt.Println("Target state reached at ",counter,strv)
			}
		}
		seen[strv] = true
	}
	fmt.Printf("DONE: %d\n",counter-cyclestart)
}
