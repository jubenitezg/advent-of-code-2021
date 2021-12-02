package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day 1 challenge from Advent of Code 2021

const File = "./input/sonar_sweep_d1.txt"

func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, num)
	}

	return lines
}

func solve(input []int) int {
	n := len(input)
	if n < 0 {
		return -1
	}
	answer := 0
	curr := input[0]
	for i := 1; i < n; i++ {
		next := input[i]
		if curr < next {
			answer++
		}
		curr = next
	}
	return answer
}

func main() {
	input := readFile(File)
	//input := []int{199,200,208,210,200,207,240,269,260,263}
	fmt.Println(solve(input))
}