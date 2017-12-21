package main

import (
	"bufio"
	"fmt"
	"math"
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

type collision struct {
	p1, p2, t int
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

func quadraticSolver(p1, v1, a1, p2, v2, a2 int) (float64, float64) {
	a := float64(a1-a2) / 2
	b := float64(v1-v2) + float64(a1-a2)/2
	c := float64(p1 - p2)

	square := math.Pow(b, 2) - (4 * a * c)
	if square < 0 {
		// must be imaginary, hence return NaN
		return math.NaN(), math.NaN()
	}
	root := math.Sqrt(square)
	x1 := (-b + root) / (2 * a)
	x2 := (-b - root) / (2 * a)
	return x1, x2
}

// Boolean indicates singularity which is always valid
func linearSolver(p1, v1, p2, v2 int) (float64, bool) {
	// redundant case
	if p1 == p2 && v1 == v2 {
		return -1, true
	}
	return float64(p1-p2) / float64(v2-v1), false
}

func isFloat(x float64) bool {
	return x-math.Floor(x) > 0
}

func isValidTime(t float64) bool {
	return !math.IsNaN(t) && !math.IsInf(t, 0) && t >= 0 && !isFloat(t)
}

func sortPair(x []float64) []float64 {
	ret := make([]float64, 2)
	if x[0] > x[1] {
		ret[0] = x[1]
		ret[1] = x[0]
		return ret
	}
	return x
}

// NOTE: This is totally awful
func whenColide(p1, p2 particle) int {
	var tx, ty, tz []float64
	var rx, ry, rz bool
	if p1.a.x == p2.a.x {
		tmp, tmpr := linearSolver(p1.p.x, p1.v.x, p2.p.x, p2.v.x)
		rx = tmpr
		tx = append(tx, tmp, tmp)
	} else {
		t1, t2 := quadraticSolver(p1.p.x, p1.v.x, p1.a.x, p2.p.x, p2.v.x, p2.a.x)
		if isValidTime(t1) && isValidTime(t2) {
			tx = append(tx, t1, t2)
		} else if isValidTime(t1) {
			tx = append(tx, t1, t1)
		} else if isValidTime(t2) {
			tx = append(tx, t2, t2)
		} else {
			return -1
		}
	}
	if p1.a.y == p2.a.y {
		tmp, tmpr := linearSolver(p1.p.y, p1.v.y, p2.p.y, p2.v.y)
		ry = tmpr
		ty = append(ty, tmp, tmp)
	} else {
		t1, t2 := quadraticSolver(p1.p.y, p1.v.y, p1.a.y, p2.p.y, p2.v.y, p2.a.y)
		if isValidTime(t1) && isValidTime(t2) {
			ty = append(ty, t1, t2)
		} else if isValidTime(t1) {
			ty = append(ty, t1, t1)
		} else if isValidTime(t2) {
			ty = append(ty, t2, t2)
		} else {
			return -1
		}
	}
	if p1.a.z == p2.a.z {
		tmp, tmpr := linearSolver(p1.p.z, p1.v.z, p2.p.z, p2.v.z)
		rz = tmpr
		tz = append(tz, tmp, tmp)
	} else {
		t1, t2 := quadraticSolver(p1.p.z, p1.v.z, p1.a.z, p2.p.z, p2.v.z, p2.a.z)
		if isValidTime(t1) && isValidTime(t2) {
			tz = append(tz, t1, t2)
		} else if isValidTime(t1) {
			tz = append(tz, t1, t1)
		} else if isValidTime(t2) {
			tz = append(tz, t2, t2)
		} else {
			return -1
		}
	}
	tx = sortPair(tx)
	ty = sortPair(ty)
	tz = sortPair(tz)
	for i := 0; i < 2; i++ {
		res := make([]float64, 0)
		ret := true
		if !rx {
			res = append(res, tx[i])
		}
		if !ry {
			res = append(res, ty[i])
		}
		if !rz {
			res = append(res, tz[i])
		}
		if len(res) == 0 {
			return 0
		}
		currNum := res[0]
		for _, c := range res {
			ret = ret && currNum == c && isValidTime(c)
		}
		if ret {
			return int(currNum)
		}
	}
	return -1
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

	collisions := make([]collision, 0)
	for i := 0; i < len(particles)-1; i++ {
		for j := i + 1; j < len(particles); j++ {
			ret := whenColide(particles[i], particles[j])
			if ret != -1 {
				collisions = append(collisions, collision{particles[i].label, particles[j].label, ret})
			}
		}
	}

	sort.Slice(collisions, func(i, j int) bool {
		return collisions[i].t < collisions[j].t
	})

	removed := make(map[int]bool)
	currentGroup := make([]int, 0)
	currentLabel := -1
	for _, c := range collisions {
		if currentLabel != c.t {
			fmt.Println("Adding", currentGroup, "new label", c.t)
			for _, r := range currentGroup {
				removed[r] = true
			}
			currentGroup = make([]int, 0)
			currentLabel = c.t
		}
		_, ok1 := removed[c.p1]
		_, ok2 := removed[c.p2]
		if ok1 || ok2 {
			continue
		}
		currentGroup = append(currentGroup, c.p1, c.p2)
	}
	// Finish off any that remain
	for _, r := range currentGroup {
		removed[r] = true
	}

	fmt.Printf("DONE: %d particles remaining\n", len(particles)-len(removed))

}
