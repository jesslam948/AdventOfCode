// Description: Advent of Code 2021 - Problem 1.2
// 	- determine the number of sums of measurements in a window of 3 that are larger (deeper) than the previous

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

	count := 0 // result to return

	// scan everything into an array
	var measurements []int64
	for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		measurements = append(measurements, curr)
	}

	// sum the first 3 numbers
	var sum int64 = 0
	for i := 0; i < 3; i++ {
		sum += measurements[i]
	}

	// go through the rest of the numbers
	for i := 3; i < len(measurements); i++ {
		newsum := sum - measurements[i-3] + measurements[i]
		if sum < newsum {
			count++
		}
		sum = newsum
	}

	fmt.Printf("Result: %d", count)
}
