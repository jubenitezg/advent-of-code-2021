package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const File = "./src/day14/d14.txt"

func readFile(fileName string) (string, map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()
	scanner.Scan() // white space
	pairs := make(map[string]string)
	var key, value string
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s -> %s", &key, &value)
		pairs[key] = value
	}
	return template, pairs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(template string, pairs map[string]string, steps int) int {
	currPairCount := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		p := template[i : i+2]
		currPairCount[p]++
	}
	for step := 0; step < steps; step++ {
		nextPairCount := make(map[string]int)
		for pair, times := range currPairCount {
			p1 := string(pair[0]) + pairs[pair]
			p2 := pairs[pair] + string(pair[1])
			//fmt.Printf("%s produces (%s,%s)\n", pair, p1, p2)
			// for each pair the new pair should also appear the same amount
			nextPairCount[p1] += times
			nextPairCount[p2] += times
		}
		currPairCount = nextPairCount
	}
	// (NC,CN)
	// (NB,BC)
	// (CH,HB)
	// N C N B C HB
	currCount := make(map[uint8]int)
	for pair, times := range currPairCount {
		// count first letter, next pair contains the other 'repeated' letter
		currCount[pair[0]] += times
	}
	// last letter always left out
	currCount[template[len(template)-1]]++
	maxV := math.MinInt
	minV := math.MaxInt
	for _, count := range currCount {
		maxV = max(maxV, count)
		minV = min(minV, count)
	}

	return maxV - minV
}

func main() {
	template, pairs := readFile(File)
	fmt.Println(solve(template, pairs, 10))
	fmt.Println(solve(template, pairs, 40))
}
