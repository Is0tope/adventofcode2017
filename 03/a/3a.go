package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func getlayer(x int) int {
	val:= math.Sqrt(float64(x))
	mn := int(math.Floor(val))
	layer := -1
	// its odd, so layer is same
	if (val - float64(mn) == 0) && (mn % 2) == 1 {
		layer = int(val)
	}else {
		if mn % 2 == 0 {
			layer = mn + 1
		}else{
			layer = mn + 2
		}
	}
	return int((layer + 1)/2)
}

func getsize(layer int) int {
	return int(math.Pow(float64((2*(layer-1))-1),2))
}

func main() {
	// NUM := 368078
	NUM,_ := strconv.Atoi(os.Args[1])

	layer := getlayer(NUM)
	fmt.Printf("layer: %d\n",layer)
	// Break out early for exact match
	if NUM == int(math.Pow(float64((2*layer)-1),2)) {
		fmt.Printf("DONE: %d\n",(2*layer)-2)
		return
	}
	size := getsize(layer)
	remain := NUM - size
	side := int(math.Sqrt(float64(getsize(layer+1))))
	middle := side/2
	fmt.Printf("size: %d\n",size)
	fmt.Printf("remain: %d\n",remain)
	fmt.Printf("side: %d\n",side)
	fmt.Printf("middle:%d\n",middle)

	offset := remain % (side-1)
	fmt.Printf("offset: %d\n",offset)
	fmt.Printf("DONE: %d\n",(layer-1)+int(math.Abs(float64(offset-middle))))

}
