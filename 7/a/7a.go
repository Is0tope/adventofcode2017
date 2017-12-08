package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseChildren(str string) []string {
	return strings.Split(strings.Replace(str, " ", "", -1), ",")
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// regex
	rparent := regexp.MustCompile(`(?P<name>\w+) \((?P<value>\d+)\) -> (?P<children>.*)`)
	// rleaf := regexp.MustCompile(`(?P<name>\w+) \((?P<value>\d+)\)`)

	// tree map
	child2parent := make(map[string]string)

	for scanner.Scan() {
		text := scanner.Text()
		// parent node
		if strings.ContainsAny(text, "->") {
			match := rparent.FindStringSubmatch(text)
			children := parseChildren(match[3])
			for _, c := range children {
				child2parent[c] = match[1]
			}
		} else {
			// match := rleaf.FindStringSubmatch(text)
			// fmt.Println(match)
		}
	}
	// This is dumb, but cant find a way to get keys
	test := ""
	for k := range child2parent {
		test = k
		break
	}
	fmt.Println(test)
	valid := true
	for valid {
		v, ok := child2parent[test]
		if ok {
			test = v
		} else {
			break
		}
	}
	fmt.Printf("DONE: %s\n", test)
}
