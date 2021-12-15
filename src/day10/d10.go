package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const File = "./src/day10/d10.txt"

var illegal = map[int32]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func mismatch(input []string) (int, [][]int32) {
	var incomplete [][]int32
	mismatches := 0
	for _, line := range input {
		curr := 0
		var stack []int32
		for _, bracket := range line {
			switch bracket {
			case '(', '{', '[', '<':
				stack = append(stack, bracket)
				continue
			}
			top := len(stack) - 1
			if stack[top] == '(' && bracket != ')' {
				curr += illegal[bracket]
			} else if stack[top] == '{' && bracket != '}' {
				curr += illegal[bracket]
			} else if stack[top] == '[' && bracket != ']' {
				curr += illegal[bracket]
			} else if stack[top] == '<' && bracket != '>' {
				curr += illegal[bracket]
			}
			stack = stack[:top]
		}
		mismatches += curr
		if curr == 0 {
			incomplete = append(incomplete, stack)
		}
	}

	return mismatches, incomplete
}

func solve1(input []string) int {
	mismatches, _ := mismatch(input)

	return mismatches
}

func solve2(input []string) int {
	_, incomplete := mismatch(input)
	var scores []int
	for _, stack := range incomplete {
		score := 0
		for len(stack) > 0 {
			score *= 5
			top := len(stack) - 1
			if stack[top] == '(' {
				score += 1
			} else if stack[top] == '{' {
				score += 3
			} else if stack[top] == '[' {
				score += 2
			} else if stack[top] == '<' {
				score += 4
			}
			stack = stack[:top]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	input := readFile(File)
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
