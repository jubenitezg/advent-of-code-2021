package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day2/dive_d2.txt"

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line:= scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func solve1(input []string) int {
	h := 0
	d := 0
	for _,v := range input {
		split := strings.Split(v, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}
		switch split[0] {
		case "forward":
			h += n
		case "down":
			d += n
		case "up":
			d -= n
		}
	}
	return d*h
}

func solve2(input []string) int {
	h := 0
	d := 0
	a := 0
	for _,v := range input {
		split := strings.Split(v, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			return -1
		}
		switch split[0] {
		case "forward":
			h += n
			d += a * n
		case "down":
			a += n
		case "up":
			a -= n
		}
	}
	return d*h
}

func main() {
	fmt.Println(solve1(readFile(File)))
	fmt.Println(solve2(readFile(File)))
}