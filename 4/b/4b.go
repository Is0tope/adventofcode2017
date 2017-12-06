package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func anagram(str []rune) []string {
	if len(str) == 1 {
		ret := make([]string,1)
		ret[0] = string(str)
		return ret
	}
	if len(str) < 3 {
		ret := make([]string,2)
		ret[0] = string(str)
		ret[1] = string(str[1])+string(str[0])
		return ret
	}
	ret := make([]string,0)
	for pos,char := range str {
		cpy := make([]rune, len(str))
		copy(cpy,str)
		data := anagram(append(cpy[:pos], cpy[pos+1:]...))
		for _,d := range data {
			ret = append(ret,string(char)+string(d))
		}
	}
	return ret
}

func main() {
	file,_ := os.Open("../input.txt")
	defer file.Close()
	
	counter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		tokens := strings.Split(scanner.Text()," ")
		m := make(map[string]bool)
		flag := false
		for _,t := range tokens {
			anas := anagram([]rune(t))
			// fmt.Println(anas)
			for _,a := range anas {
				_,ok := m[a]
				if ok {
					flag = true
					break
				}
				
			}
			if flag {
				break
			}
			// otherwise insert
			for _,a := range anas {
				m[a] = true
			}
		}
		if !flag {
			counter++
			fmt.Println("FLAGGED")
		}
	}
	fmt.Printf("DONE: %d\n",counter)
}
