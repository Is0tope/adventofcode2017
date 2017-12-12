package main

import (
	"fmt"
	// "math"
	"os"
	"strconv"
)

type Point struct {
	x,y,val int
}

const (
	EAST = iota
	NORTH = iota
	WEST = iota
	SOUTH = iota
)

func coord2str(x,y int) string {
	return fmt.Sprintf("%d-%d",x,y)
}

func main() {
	target,err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Enter a target number")
		return
	}
	fmt.Printf("TARGET: %d\n",target)
	
	// map
	MAP := make(map[string]int)

	// // Initial coords
	X := 0
	Y := 0
	dir := EAST

	// others
	layer := 1
	ordinal := 1
	layercnt := 0
	layersz := 1
	sidel := 1
	val := 0
	
	// add first value
	MAP[coord2str(X,Y)] = 1
	
	for val < target {
		val = 0
		// increment
		ordinal++
		layercnt++
		// Change coordinates
		switch dir {
			case EAST:
				X++
			case NORTH:
				Y++
			case WEST:
				X--
			case SOUTH:
				Y--
		}
		if layercnt == layersz {
			layercnt = 0
			// such a hack
			if layer == 1 {
				layersz += 7
				sidel = 2
			}else {
				layersz += 8
				sidel += 2
			}
			layer++
			// change dir
			dir++
			dir = dir % 4
			fmt.Printf("\n")
		}
		// Change direction
		if layercnt > 0 && layercnt < layersz-1 && (sidel-1) == layercnt % sidel {
			dir++
			dir = dir % 4
		}
		// LOGIC
		// calculate value by checking neighbours
		// top left
		v,ok := MAP[coord2str(X-1,Y+1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// top
		v,ok = MAP[coord2str(X,Y+1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// top right
		v,ok = MAP[coord2str(X+1,Y+1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// right
		v,ok = MAP[coord2str(X+1,Y)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// bottom right
		v,ok = MAP[coord2str(X+1,Y-1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// bottom
		v,ok = MAP[coord2str(X,Y-1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// bottom left
		v,ok = MAP[coord2str(X-1,Y-1)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		// left
		v,ok = MAP[coord2str(X-1,Y)]
		if ok {
			val += v
			// fmt.Printf("got: %d\n",v)
		}
		
		// add to map
		MAP[coord2str(X,Y)] = val
		fmt.Printf("ord: %d, coords: (%d,%d), layer: %d, layercnt: %d, layersz: %d, sidel: %d, dir: %d, val: %d\n",ordinal,X,Y,layer,layercnt,layersz,sidel,dir,val)
		
	}
	fmt.Printf("DONE: %d\n",val)
}