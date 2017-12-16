package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPeriod(n int) int {
	return (2 * n) - 2
}

type sentry struct {
	time   int
	period int
}

// This is more or less brute force as far as I can see. I had a plausible method using a sieve,
// but unable to get it from mind into code

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sentries := make([]sentry, 0)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ": ")
		time, _ := strconv.Atoi(tokens[0])
		dist, _ := strconv.Atoi(tokens[1])
		period := getPeriod(dist)

		sentries = append(sentries, sentry{time, period})
	}

	delay := 0
	for {
		failed := false
		for _, s := range sentries {
			if (s.time+delay)%s.period == 0 {
				failed = true
				break
			}
		}
		if !failed {
			break
		}
		delay++
	}

	fmt.Printf("DONE: %d\n", delay)
}
