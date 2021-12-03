package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const File = "./src/day3/d3.txt"

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func count(input []string, n int) ([]int, []int) {
	cnt1 := make([]int, n)
	cnt0 := make([]int, n)
	for _, v := range input {
		for j, ch := range v {
			if ch == '0' {
				cnt0[j]++
			} else {
				cnt1[j]++
			}
		}
	}
	return cnt1, cnt0
}

func find(sl []string, col int, val uint8) []string {
	var match []string
	for _, v := range sl {
		if v[col] == val {
			match = append(match, v)
		}
	}
	return match
}

func solve1(input []string) int64 {
	n := len(input[0])
	cnt1, cnt0 := count(input, n)
	gamma := ""
	eps := ""
	for i, cont := range cnt0 {
		cont1 := cnt1[i]
		if cont > cont1 {
			eps += "1"
			gamma += "0"
		} else {
			gamma += "1"
			eps += "0"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsInt, _ := strconv.ParseInt(eps, 2, 64)

	return gammaInt * epsInt
}

func solve2(input []string) int64 {
	n := len(input[0])
	greatNums := make([]string, len(input))
	copy(greatNums, input)
	curr := 0
	for len(greatNums) > 1 {
		cnt1, cnt0 := count(greatNums, n)
		var max uint8
		if cnt1[curr] >= cnt0[curr] {
			max = '1'
		} else {
			max = '0'
		}
		ngn := find(greatNums, curr, max)
		curr++
		greatNums = ngn
	}
	smallNums := make([]string, len(input))
	copy(smallNums, input)
	curr = 0
	for len(smallNums) > 1 {
		cnt1, cnt0 := count(smallNums, n)
		var min uint8
		if cnt1[curr] >= cnt0[curr] {
			min = '0'
		} else {
			min = '1'
		}
		nsm := find(smallNums, curr, min)
		curr++
		smallNums = nsm
	}
	og, _ := strconv.ParseInt(greatNums[0], 2, 64)
	co2, _ := strconv.ParseInt(smallNums[0], 2, 64)
	return og * co2
}

func main() {
	fmt.Println(solve1(readFile(File)))
	fmt.Println(solve2(readFile(File)))
}
