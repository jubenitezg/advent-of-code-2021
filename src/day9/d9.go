package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day9/d9.txt"

var di = [4]int{-1, 0, 1, 0}
var dj = [4]int{0, 1, 0, -1}

type Pair struct {
	i int
	j int
}

func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input [][]int
	for scanner.Scan() {
		var row []int
		for _, str := range strings.Split(scanner.Text(), "") {
			atoi, _ := strconv.Atoi(str)
			row = append(row, atoi)
		}
		input = append(input, row)
	}
	return input
}

func lows(cavern [][]int) []*Pair {
	n := len(cavern)
	m := len(cavern[0])
	var lowPoints []*Pair
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			smaller := true
			for k := 0; k < 4; k++ {
				ni := i + di[k]
				nj := j + dj[k]
				if 0 <= ni && ni < n && 0 <= nj && nj < m {
					smaller = smaller && cavern[i][j] < cavern[ni][nj]
				}
			}
			if smaller {
				lowPoints = append(lowPoints, &Pair{i, j})
			}
		}
	}
	return lowPoints
}

func floodBasin(i int, j int, cavern [][]int, seen map[Pair]bool) int {
	n := len(cavern)
	m := len(cavern[0])
	seen[Pair{i, j}] = true
	area := 1
	for k := 0; k < 4; k++ {
		ni := i + di[k]
		nj := j + dj[k]
		if _, ok := seen[Pair{ni, nj}]; !ok && 0 <= ni && ni < n && 0 <= nj && nj < m && cavern[ni][nj] != 9 {
			area += floodBasin(ni, nj, cavern, seen)
		}
	}
	return area
}

func solve1(cavern [][]int) int {
	riskPoints := 0
	for _, coord := range lows(cavern) {
		i, j := coord.i, coord.j
		riskPoints += cavern[i][j] + 1
	}

	return riskPoints
}

func solve2(cavern [][]int) int {
	var basins []int
	seen := make(map[Pair]bool)
	for _, coord := range lows(cavern) {
		i, j := coord.i, coord.j
		basins = append(basins, floodBasin(i, j, cavern, seen))
		seen[*coord] = true
	}
	top1, top2, top3 := math.MinInt, math.MinInt, math.MinInt
	for _, basin := range basins {
		if basin > top1 {
			top3 = top2
			top2 = top1
			top1 = basin
		} else if basin > top2 {
			top3 = top2
			top2 = basin
		} else if basin > top3{
			top3 = basin
		}
	}

	return top1 * top2 * top3
}

func main() {
	input := readFile(File)
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
