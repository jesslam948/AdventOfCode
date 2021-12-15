// Description: Advent of Code 2021 - Problem 1.1
// 	- determine the number of measurements that are larger (deeper) than the previous

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	count := 0     // result to return
	scanner.Scan() // read first input since it's a special case
	prev, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if prev < curr {
			count++
		}
		prev = curr
	}

	fmt.Printf("Result: %d", count)
}
