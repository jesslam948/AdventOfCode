// Description: Advent of Code 2021 - Problem 5.1
// 	- determine the number of points that have at least 2 or more lines running through them
//	- for this, consider only horizontal or vertical lines

// Thoughts:
// 	- go through each line, if not x1 = x2 or y1 = y2 continue
//		- for each of the points that the line covers, update map
//			- if the value in the map == 2 (it just became 2, update counter)
// UPDATED to be cleaner

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

// since the math package seems to do these functions on floats only :(
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func absInt(a int) int {
	if a < 0 {
		return -1 * a
	} else {
		return a
	}
}

// cleaned up the func :D
func updateMap(PointMap map[Point]int, p1, p2 Point, count int) int {
	deltaY := 1
	deltaX := 1
	if p1.y > p2.y {
		deltaY = -1
	} else if p1.y == p2.y {
		deltaY = 0
	}

	if p1.x > p2.x {
		deltaX = -1
	} else if p1.x == p2.x {
		deltaX = 0
	}

	steps := maxInt(absInt(p1.x-p2.x), absInt(p1.y-p2.y))

	if !(deltaX == 0 || deltaY == 0) {
		return count
	}
	for i := 0; i <= steps; i++ {
		var p Point
		p.x = deltaX*i + p1.x
		p.y = deltaY*i + p1.y

		PointMap[p] += 1
		if PointMap[p] == 2 {
			count++
		}
	}

	return count
}

func main() {
	// open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// parse the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // set to scan lines

	PointMap := make(map[Point]int)
	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		// make points
		var p1 Point
		p1s := strings.Split(line[0], ",")
		p1.x, _ = strconv.Atoi(p1s[0])
		p1.y, _ = strconv.Atoi(p1s[1])

		var p2 Point
		p2s := strings.Split(line[1], ",")
		p2.x, _ = strconv.Atoi(p2s[0])
		p2.y, _ = strconv.Atoi(p2s[1])

		count = updateMap(PointMap, p1, p2, count)
	}

	fmt.Printf("Found: %d", count)
}
