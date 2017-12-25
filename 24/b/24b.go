package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type component struct {
	a, b int
}

type state struct {
	end       int
	order     []int
	remaining map[int]bool
	score     int
}

func newState(initial int, components map[int]component) state {
	// set up initial order with first component
	order := make([]int, 1)
	order[0] = initial
	// set up remaining
	remaining := make(map[int]bool)
	for k := range components {
		// ignore initial
		if initial == k {
			continue
		}
		remaining[k] = true
	}
	var end int
	if components[initial].a == 0 {
		end = components[initial].b
	} else {
		end = components[initial].a
	}
	// score
	score := components[initial].a + components[initial].b
	return state{end, order, remaining, score}
}

func loadComponents(fname string) map[int]component {
	file, _ := os.Open(fname)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	components := make(map[int]component)

	cnt := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "/")
		a, _ := strconv.Atoi(tokens[0])
		b, _ := strconv.Atoi(tokens[1])
		components[cnt] = component{a, b}
		cnt++
	}
	return components
}

func cloneMap(m map[int]bool) map[int]bool {
	ret := make(map[int]bool)
	for k := range m {
		ret[k] = true
	}
	return ret
}

func nextStates(s state, components map[int]component) []state {
	possible := make([]state, 0)
	// check all components for possibility
	for k := range s.remaining {
		cmp := components[k]
		if s.end == cmp.a || s.end == cmp.b {
			// ending
			var end int
			if s.end == cmp.a {
				end = cmp.b
			} else {
				end = cmp.a
			}
			// order (need to clone)
			order := make([]int, len(s.order))
			copy(order, s.order)
			order = append(order, k)
			// remaining
			remaining := cloneMap(s.remaining)
			delete(remaining, k)
			// score
			score := s.score + cmp.a + cmp.b
			// add to posibilities
			possible = append(possible, state{end, order, remaining, score})
		}
	}
	return possible
}

func main() {
	// Load the components
	components := loadComponents("../input.txt")

	initialStates := make([]state, 0)
	for k, v := range components {
		if v.a == 0 || v.b == 0 {
			is := newState(k, components)
			initialStates = append(initialStates, is)
		}
	}

	// max score & length
	maxLength := 0
	maxScore := 0
	var maxState state

	// stack
	pile := stack.New()

	// insert initial items
	pile.Push(initialStates[1])
	pile.Push(initialStates[0])

	for pile.Len() > 0 {
		// get the state
		s := pile.Pop().(state)
		// fmt.Println(s)
		if len(s.order) > maxLength {
			// reset list length
			maxState = s
			maxLength = len(s.order)
			maxScore = s.score
		} else if len(s.order) == maxLength {
			// see if score is bigger
			if s.score > maxScore {
				maxScore = s.score
				maxState = s
			}
		}
		// get posibilities
		possible := nextStates(s, components)
		// TODO
		// Add them to the queue
		for _, p := range possible {
			pile.Push(p)
		}
	}

	fmt.Printf("DONE: max length = %d\n", maxLength)
	fmt.Println(maxState)

}
