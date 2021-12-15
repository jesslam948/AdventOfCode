// Description: Advent of Code 2021 - Problem 2.2
// 	- determine final horizontal position and depth after following an inputted path return the product
// 	- this time, up and down don't directly change the depth, but changes the aim, which affects depth

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
	var aim int64 = 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0]
		num, _ := strconv.ParseInt(line[1], 10, 64)
		switch dir {
		case "forward":
			hpos += num
			depth += aim * num
		case "up":
			aim -= num
		case "down":
			aim += num
		}
	}

	fmt.Printf("hpos: %d, depth: %d, product = %d", hpos, depth, hpos*depth)
}
