package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

type vector3 struct {
	x, y, z int
}

type particle struct {
	p, v, a vector3
	label   int
}

func (v vector3) magnitude() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func parseParticle(str string, label int) particle {
	re := regexp.MustCompile("<(.?\\d+),(.?\\d+),(.?\\d+)>, v=<(.?\\d+),(.?\\d+),(.?\\d+)>, a=<(.?\\d+),(.?\\d+),(.?\\d+)>")
	groups := re.FindStringSubmatch(str)
	vals := make([]int, len(groups))
	for i, g := range groups[1:] {
		v, _ := strconv.Atoi(strings.Trim(g, " "))
		vals[i] = v
	}
	position := vector3{vals[0], vals[1], vals[2]}
	velocity := vector3{vals[3], vals[4], vals[5]}
	acceleration := vector3{vals[6], vals[7], vals[8]}
	return particle{position, velocity, acceleration, label}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	particles := make([]particle, 0)

	label := 0
	for scanner.Scan() {
		text := scanner.Text()
		particles = append(particles, parseParticle(text, label))
		label++
	}

	// Sort in order of presedence: lowest a, lowest v, lowest p
	sort.Slice(particles, func(i, j int) bool {
		p1 := particles[i]
		p2 := particles[j]
		p1am := p1.a.magnitude()
		p2am := p2.a.magnitude()
		if p1am != p2am {
			return p1am < p2am
		}
		p1vm := p1.v.magnitude()
		p2vm := p2.v.magnitude()
		if p1vm != p2vm {
			return p1vm < p2vm
		}
		p1pm := p1.p.magnitude()
		p2pm := p2.p.magnitude()
		if p1pm != p2pm {
			return p1pm < p2pm
		}
		return false
	})

	// fmt.Println(particles)

	// get the closest particle (the one with lowest magnitude)
	closestParticle := particles[0]
	fmt.Printf("DONE: %d is closest with aMag: %d, vMag: %d, pMag: %d\n",
		closestParticle.label,
		closestParticle.a.magnitude(),
		closestParticle.v.magnitude(),
		closestParticle.p.magnitude())
}
