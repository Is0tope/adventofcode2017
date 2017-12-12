package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseChildren(str string) []string {
	return strings.Split(strings.Replace(str, " ", "", -1), ",")
}

type node struct {
	name     string
	val      int
	children []*node
}

// returns: value, should, index, ok
func consensus(arr []int) (int, int, int, bool) {
	m := make(map[int][]int)
	for i, a := range arr {
		if _, ok := m[a]; ok {
			m[a] = append(m[a], i)
		} else {
			m[a] = []int{}
			m[a] = append(m[a], i)
		}
	}
	// find the odd one out (only one)
	if len(m) > 1 {
		for k, v := range m {
			if len(v) == 1 {
				otherval := 0
				for _, a := range arr {
					if k != a {
						otherval = a
						break
					}
				}
				return k, v[0], otherval, false
			}
		}
	}
	// all the same
	return arr[0], 0, 0, true
}

func imbalance(n *node) (int, *node, int, bool) {
	if len(n.children) == 0 {
		return n.val, nil, 0, true
	}
	sum := 0
	vals := []int{}
	for _, c := range n.children {
		ret, nd, ot, ok := imbalance(c)
		// if there is a node returned, pass it back up
		if !ok {
			return ret, nd, ot, false
		}
		sum += ret
		vals = append(vals, ret)
	}
	// check for consensus
	mv, mi, ot, ok := consensus(vals)
	fmt.Printf("mv: %d, mi: %d, ot: %d, ok: %t\n", mv, mi, ot, ok)
	if !ok {
		return mv, n.children[mi], n.children[mi].val + ot - vals[mi], false
	}
	return sum + n.val, nil, 0, true
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// regex
	rparent := regexp.MustCompile(`(?P<name>\w+) \((?P<value>\d+)\) -> (?P<children>.*)`)
	rleaf := regexp.MustCompile(`(?P<name>\w+) \((?P<value>\d+)\)`)

	// tree map
	nodes := make(map[string]*node)

	for scanner.Scan() {
		text := scanner.Text()
		// parent node
		if strings.ContainsAny(text, "->") {
			match := rparent.FindStringSubmatch(text)
			children := parseChildren(match[3])
			v, ok := nodes[match[1]]
			// Node does not exist
			if !ok {
				v = new(node)
				v.name = match[1]
				nodes[v.name] = v
			}
			val, _ := strconv.Atoi(match[2])
			v.val = val
			for _, c := range children {
				cv, ok2 := nodes[c]
				if !ok2 {
					cv = new(node)
					cv.name = c
					nodes[c] = cv
				}
				v.children = append(v.children, cv)
			}
		} else {
			// child node
			match := rleaf.FindStringSubmatch(text)
			v, ok := nodes[match[1]]
			// Node does not exist
			if !ok {
				v = new(node)
				v.name = match[1]
				nodes[v.name] = v
			}
			val, _ := strconv.Atoi(match[2])
			v.val = val
		}
	}

	// Can't be bothered to copy to root finding code
	_, nd, ot, _ := imbalance(nodes["rqwgj"])
	fmt.Printf("DONE: %s should be: %d, but is %d\n", nd.name, ot, nd.val)
}
