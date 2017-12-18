package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spin(arr []*rune, n int) []*rune {
	l := len(arr)
	return append(arr[l-n:], arr[:l-n]...)
}

func exchange(arr []*rune, a int, b int) []*rune {
	ret := make([]*rune, len(arr))
	copy(ret, arr)
	tmp := ret[a]
	ret[a] = ret[b]
	ret[b] = tmp
	return ret
}

func partner(mp map[rune]*rune, a rune, b rune) {
	tmp := *mp[a]
	*mp[a] = *mp[b]
	*mp[b] = tmp
	tmpk := mp[a]
	mp[a] = mp[b]
	mp[b] = tmpk
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")

	letters := []rune("abcdefghijklmnop")
	// letters := []rune("abcde")

	programs := make([]*rune, 0)
	programMap := make(map[rune]*rune)
	for _, x := range letters {
		val := x
		ptr := &val
		programs = append(programs, ptr)
		programMap[x] = ptr
	}

	for _, inst := range tokens {
		typ := inst[0]
		switch typ {
		case 's':
			n, _ := strconv.Atoi(inst[1:])
			programs = spin(programs, n)
		case 'x':
			sp := strings.Split(inst[1:], "/")
			a, _ := strconv.Atoi(sp[0])
			b, _ := strconv.Atoi(sp[1])
			programs = exchange(programs, a, b)
		case 'p':
			sp := strings.Split(inst[1:], "/")
			partner(programMap, rune(sp[0][0]), rune(sp[1][0]))
		}
		fmt.Printf("%s : ", inst)
		for _, x := range programs {
			fmt.Printf("%c", *x)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("DONE\n")
}
