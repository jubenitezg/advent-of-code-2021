package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day4/d4.txt"

func readFile(fileName string) ([]int, [][][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	scanner.Scan()
	split := strings.Split(scanner.Text(), ",")
	for _, v := range split {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, atoi+1) // adjust by adding 1
	}
	var boards [][][]int
	scanner.Scan() //white space
	for scanner.Scan() {
		var board [][]int
		for i := 0; i < 5; i++ {
			split := strings.Fields(scanner.Text())
			//fmt.Println(split)
			board = append(board, make([]int, 5))
			for j, v := range split {
				atoi, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				board[i][j] = atoi + 1
			}
			scanner.Scan()
		}
		boards = append(boards, board)
	}
	return numbers, boards
}

func getSumUnmarked(board [][]int) int {
	sum := 0
	for i := range board {
		for j := range board {
			if board[i][j] > 0 {
				sum += board[i][j] - 1 //re-adjust to original
			}
		}
	}
	return sum
}

func win(board [][]int) bool {
	for i := range board {
		row, col := 0, 0
		for j := range board {
			if board[i][j] < 0 {
				row++
			}
			if board[j][i] < 0 {
				col++
			}
		}
		if row == 5 || col == 5 {
			return true
		}
	}

	return false
}

func solve1(numbers []int, boards [][][]int) int {
	for _, number := range numbers {
		for _, board := range boards {
			for i := range board {
				for j := range board {
					if board[i][j] == number {
						board[i][j] *= -1 //mark as seen
					}
				}
			}
			if win(board) {
				return getSumUnmarked(board) * (number - 1)
			}
		}
	}

	return -1
}

func solve2(numbers []int, boards [][][]int) int {
	won := make(map[int]bool)
	lastBoardNumber := -1
	lastNumber := -1
	for _, number := range numbers {
		for boardNumber, board := range boards {
			if _, ok := won[boardNumber]; !ok {
				won[boardNumber] = false
			}
			if hasWon, _ := won[boardNumber]; !hasWon {
				for i := range board {
					for j := range board {
						if board[i][j] == number {
							board[i][j] *= -1 //mark as seen
						}
					}
				}
			}
			if hasWon, _ := won[boardNumber]; !hasWon && win(board) {
				won[boardNumber] = true
				lastBoardNumber = boardNumber
				lastNumber = number
			}
		}
	}
	lastBoard := boards[lastBoardNumber]

	return getSumUnmarked(lastBoard) * (lastNumber - 1)
}

func main() {
	numbers, boards := readFile(File)
	fmt.Println(solve1(numbers, boards))
	fmt.Println(solve2(numbers, boards))
}
