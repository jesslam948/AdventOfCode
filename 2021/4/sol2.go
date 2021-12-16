// Description: Advent of Code 2021 - Problem 4.2
// 	- determine the last winning bingo board
// 	- updated sol1 to instead look for last winning as opposed to first winning

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// apparently Go doesn't have a built-in indexOf function like python :')
func indexOf(elem string, data []string) int {
	for i, val := range data {
		if val == elem {
			return i
		}
	}
	// not found
	return -1
}

func calculateScore(board []string, winningNum string) int64 {
	var score int64 = 0
	for _, spot := range board {
		if spot != "x" {
			spotNum, _ := strconv.ParseInt(spot, 10, 64)
			score += spotNum
		}
	}
	winNumNum, _ := strconv.ParseInt(winningNum, 10, 64)
	return score * winNumNum
}

func playBingo(bingoNums, board []string) (int, int64) {
	// play bingo!
	var cols [5]int
	var rows [5]int
	for i, num := range bingoNums {
		index := indexOf(num, board) // normally bingo boards don't have repeat numbers bc it makes it unfair
		if index != -1 {
			board[index] = "x" // mark as found
			// keep track of nums found
			cols[index%5] += 1 // a spot's column is the index % 5
			rows[index/5] += 1 // a spot's row is the index // 5
			// win bingo!
			if cols[index%5] == 5 || rows[index/5] == 5 {
				// count
				return i + 1, calculateScore(board, num)
			}
		}
	}
	// return # moves, final score
	return -1, -1 // did not win
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

	scanner.Scan()
	var bingoNums []string = strings.Split(scanner.Text(), ",")

	worstMoves := 0
	var finalScore int64 = 0
	for scanner.Scan() { // scan empty line
		// read the bingo board
		var board []string
		for i := 0; i < 5; i++ {
			scanner.Scan()
			line := strings.Fields(scanner.Text()) // so we split on whitespace rather than " "
			board = append(board, line...)
		}

		// play bingo
		moves, score := playBingo(bingoNums, board)

		if moves == -1 {
			continue
		}
		if moves > worstMoves {
			worstMoves = moves
			finalScore = score
		}
	}

	fmt.Printf("worstMoves: %d, finalScore: %d", worstMoves, finalScore)
}
