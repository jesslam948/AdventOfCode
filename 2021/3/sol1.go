// Description: Advent of Code 2021 - Problem 3.1
// 	- determine the power consumption by multiplying the gamma rate by the epsilon rate
//	- gamma rate = bin most common bit in 1st, 2nd, 3rd, etc. place of all of the bin inputs
//	- epsilon rate = bin gamma rate inverted

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

	var ones [12]int
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, bit := range line {
			if bit == '1' {
				ones[i] += 1
			}
		}
		total += 1
	}

	// construct gamma and epsilon
	gamma := ""
	epsilon := ""
	for _, num := range ones {
		if num > total-num { // most common for this position is one
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	// convert to bin to nums
	gNum, _ := strconv.ParseInt(gamma, 2, 0)
	eNum, _ := strconv.ParseInt(epsilon, 2, 0)

	fmt.Printf("gamma: %d, epsilon: %d, power = %d", gNum, eNum, gNum*eNum)
}
