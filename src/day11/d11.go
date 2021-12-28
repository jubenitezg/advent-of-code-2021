package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day11/d11.txt"

func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input [][]int
	for scanner.Scan() {
		var line []int
		for _, val := range strings.Split(scanner.Text(), "") {
			atoi, _ := strconv.Atoi(val)
			line = append(line, atoi)
		}
		input = append(input, line)
	}

	return input
}

var di = [...]int{-1, -1, -1, 0, 1, 1, 1, 0}
var dj = [...]int{-1, 0, 1, 1, 1, 0, -1, -1}

type Pair struct {
	i int
	j int
}

func tick(i, j int, mat [][]int, flashed map[Pair]bool) int {
	if _, ok := flashed[Pair{i, j}]; !(0 <= i && i < len(mat)) || !(0 <= j && j < len(mat)) || ok {
		return 0
	}
	mat[i][j]++
	flashes := 0
	if mat[i][j] > 9 {
		flashes++
		flashed[Pair{i, j}] = true
		mat[i][j] = 0
		for k := 0; k < 8; k++ {
			ni := i + di[k]
			nj := j + dj[k]
			flashes += tick(ni, nj, mat, flashed)
		}
	}

	return flashes
}


func solve1(mat [][]int, steps int) int {
	total := 0
	for step := 0; step < steps; step++ {
		flashed := make(map[Pair]bool)
		for i := range mat {
			for j := range mat {
				total += tick(i, j, mat, flashed)
			}
		}
	}

	return total
}

func solve2(mat [][]int) int {
	sum := -1
	step := 0
	for sum != 0 {
		flashed := make(map[Pair]bool)
		for i := range mat {
			for j := range mat {
				tick(i, j, mat, flashed)
			}
		}
		sum = 0
		for _, row := range mat {
			for _, value := range row {
				sum += value
			}
		}

		step++
	}

	return step
}

func main() {
	input1 := readFile(File)
	input2 := readFile(File)
	fmt.Println(solve1(input1, 100))
	fmt.Println(solve2(input2))
}
