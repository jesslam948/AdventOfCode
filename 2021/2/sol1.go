// Description: Advent of Code 2021 - Problem 2.1
// 	- determine final horizontal position and depth after following an inputted path return the product

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var hpos int64 = 0
	var depth int64 = 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		num, _ := strconv.ParseInt(line[1], 10, 64)
		switch dir {
		case "forward":
			hpos += num
		case "up":
			depth -= num
		case "down":
			depth += num
		}
	}

	fmt.Printf("hpos: %d, depth: %d, product = %d", hpos, depth, hpos*depth)
}
