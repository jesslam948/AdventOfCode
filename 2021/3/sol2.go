// Description: Advent of Code 2021 - Problem 3.1
// 	- determine the life support rating by getting the product of the oxygen and co2 ratings
//	- determine most common bit in a position and remove all other numbers, go right and keep filtering -- last number is the oxygen rating
//	- similar process for co2 except use the least common bit

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getOnes(place int, lines []string) int {
	ones := 0
	for _, line := range lines {
		if line[place] == '1' {
			ones++
		}
	}
	return ones
}

func filterBit(place int, bitKeep byte, lines []string) []string {
	var filtered []string
	for _, line := range lines {
		if line[place] == bitKeep {
			filtered = append(filtered, line)
		}
	}
	return filtered
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

	var lines []string // scan all lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// construct oxygen
	oxyLines := append([]string(nil), lines...) // get deepcopy of lines
	place := 0
	for len(oxyLines) > 1 {
		ones := getOnes(place, oxyLines)
		var bitKeep byte
		if ones >= len(oxyLines)-ones { // ones are more common
			bitKeep = '1'
		} else {
			bitKeep = '0'
		}
		oxyLines = filterBit(place, bitKeep, oxyLines)
		place++
	}

	// construct co2
	co2Lines := append([]string(nil), lines...) // get deepcopy of lines
	place = 0
	for len(co2Lines) > 1 {
		ones := getOnes(place, co2Lines)
		var bitKeep byte
		if ones >= len(co2Lines)-ones { // ones are more common, thus zeros are less common
			bitKeep = '0'
		} else {
			bitKeep = '1'
		}
		co2Lines = filterBit(place, bitKeep, co2Lines)
		place++
	}

	// construct oxy and co2
	oxygen := oxyLines[0]
	co2 := co2Lines[0]

	// convert to bin to nums
	oNum, _ := strconv.ParseInt(oxygen, 2, 0)
	cNum, _ := strconv.ParseInt(co2, 2, 0)

	fmt.Printf("oxygen: %d, co2: %d, life support = %d", oNum, cNum, oNum*cNum)
}
