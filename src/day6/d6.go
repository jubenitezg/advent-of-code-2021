package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day6/d6.txt"

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

func solve(lanterns []int, days int) int {
	counter := make([]int, 10)
	for _, v := range lanterns {
		counter[v]++
	}
	for i := 0; i < days-1; i++ {
		for lantern := 1; lantern <= 9; lantern++ {
			quantity := counter[lantern]
			//all lanterns are now -1
			counter[lantern-1] += quantity
			counter[lantern] = 0
		}
		zeros := counter[0]
		counter[9] += zeros // new lanterns
		// reset lanterns
		counter[7] += zeros
		counter[0] = 0
	}
	sum := 0
	for _, quantity := range counter {
		sum += quantity
	}

	return sum
}

func main() {
	fmt.Println(solve(readFile(File), 80))
	fmt.Println(solve(readFile(File), 256))
}
