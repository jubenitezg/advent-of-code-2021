package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const File = "./src/day7/d7.txt"

func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []int
	scanner.Scan()
	for _, str := range strings.Split(scanner.Text(), ",") {
		atoi, _ := strconv.Atoi(str)
		input = append(input, atoi)
	}
	return input
}

func solve1(input []int) int {
	sort.Ints(input)
	n := len(input)
	var median int
	if n%2 == 0 {
		median = (input[n/2] + input[n/2-1]) / 2
	} else {
		median = input[n/2]
	}
	fuel := 0
	for _, position := range input {
		fuel += int(math.Abs(float64(position - median)))
	}

	return fuel
}

func solve2(input []int) int {
	sum := 0
	for _, position := range input {
		sum += position
	}
	avg := sum / len(input)
	fuel := 0
	for _, position := range input {
		delta := int(math.Abs(float64(position - avg)))
		fuel += (delta * (delta + 1)) / 2
	}

	return fuel
}
func main() {
	fmt.Println(solve1(readFile(File)))
	fmt.Println(solve2(readFile(File)))
}
