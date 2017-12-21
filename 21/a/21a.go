package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const numIterations = 5

func parseGrid(str string) [][]rune {
	grid := make([][]rune, 0)
	tokens := strings.Split(str, "/")
	for _, t := range tokens {
		grid = append(grid, []rune(t))
	}
	return grid
}

// Transforms

func flipH(arr [][]rune) [][]rune {
	ret := make([][]rune, 0)
	for i := 0; i < len(arr); i++ {
		line := make([]rune, 0)
		for j := len(arr[i]) - 1; j >= 0; j-- {
			line = append(line, arr[i][j])
		}
		ret = append(ret, line)
	}
	return ret
}

func flipV(arr [][]rune) [][]rune {
	ret := make([][]rune, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}
	return ret
}

func rotate90(arr [][]rune) [][]rune {
	ret := make([][]rune, 0)
	// transpose
	for i := 0; i < len(arr); i++ {
		line := make([]rune, 0)
		for j := 0; j < len(arr[0]); j++ {
			line = append(line, arr[j][i])
		}
		ret = append(ret, line)
	}
	// reverse each row (flip horizontally)
	return flipH(ret)
}

func gridToString(arr [][]rune) string {
	ret := ""
	for i, l := range arr {
		ret += string(l)
		if i < len(arr)-1 {
			ret += "/"
		}
	}
	return ret
}

func prettyGrid(arr [][]rune) string {
	ret := ""
	for _, l := range arr {
		ret += string(l) + "\n"
	}
	return ret
}

// This assumes that height of these grids are the same
func mergeGridHorizontaly(a, b [][]rune) [][]rune {
	ret := make([][]rune, 0)
	for i := 0; i < len(a); i++ {
		line := make([]rune, 0)
		line = append(line, a[i]...)
		line = append(line, b[i]...)
		ret = append(ret, line)
	}
	return ret
}

func subSlice(arr [][]rune, i, j, i2, j2 int) [][]rune {
	ret := make([][]rune, 0)
	for _, c := range arr[i:i2] {
		ret = append(ret, c[j:j2])
	}
	return ret
}

func countGrid(arr [][]rune) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == '#' {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	image := [][]rune{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'}}

	enhancements := make(map[string][][]rune)

	for scanner.Scan() {
		text := scanner.Text()
		var gridi, grido [][]rune
		var inputstr, outputstr string

		// 2x2 have 20 characters, 3x3 have 34 characters in a line of text
		if len(text) == 20 {
			inputstr = text[:5]
			outputstr = text[9:]
		}
		if len(text) == 34 {
			inputstr = text[:11]
			outputstr = text[15:]
		}

		gridi = parseGrid(inputstr)
		grido = parseGrid(outputstr)

		// grid itself
		enhancements[gridToString(gridi)] = grido
		// flips
		enhancements[gridToString(flipV(gridi))] = grido
		enhancements[gridToString(flipH(gridi))] = grido

		// Do the rotations
		for i := 0; i < 3; i++ {
			gridi = rotate90(gridi)
			// rotate
			enhancements[gridToString(gridi)] = grido
			enhancements[gridToString(flipV(gridi))] = grido
			enhancements[gridToString(flipH(gridi))] = grido
		}
	}

	fmt.Println(prettyGrid(image))

	// for k := range enhancements {
	// 	fmt.Println(k)
	// }

	for iter := 0; iter < numIterations; iter++ {
		// get the image
		size := len(image)
		var offset int
		if size%2 == 0 {
			offset = 2
		} else {
			offset = 3
		}

		fmt.Println("iteration", iter, "offset", offset)

		newImage := make([][]rune, 0)
		for i := 0; i < len(image); i += offset {
			row := make([][]rune, 0)
			for j := 0; j < len(image[i]); j += offset {
				ss := subSlice(image, i, j, i+offset, j+offset)
				// fmt.Println(prettyGrid(ss))
				enhanced := enhancements[gridToString(ss)]
				// fmt.Println(prettyGrid(enhanced))
				if j == 0 {
					row = append(row, enhanced...)
				} else {
					row = mergeGridHorizontaly(row, enhanced)
				}
				// fmt.Println("row\n", prettyGrid(row))

			}
			newImage = append(newImage, row...)
		}
		image = newImage

		fmt.Println(prettyGrid(image))

	}
	fmt.Printf("\nDONE: %d\n", countGrid(image))
}
