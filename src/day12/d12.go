package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const File = "./src/day12/d12.txt"

func readFile(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := make(map[string][]string)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		if _, ok := input[nodes[0]]; !ok {
			input[nodes[0]] = make([]string, 0)
		}
		if _, ok := input[nodes[1]]; !ok {
			input[nodes[1]] = make([]string, 0)
		}
		input[nodes[0]] = append(input[nodes[0]], nodes[1])
		input[nodes[1]] = append(input[nodes[1]], nodes[0])
	}

	return input
}

func smallCave(node string) bool {
	if node == "start" || node == "end" {
		return false
	}

	return unicode.IsLower(rune(node[0]))
}

//var path = make([]string, 0)
var paths = 0

func solve1(from string, graph map[string][]string, vis map[string]bool) {
	if smallCave(from) {
		vis[from] = true
	}
	//path = append(path, from)
	if from == "end" {
		paths++
		//fmt.Println(path)
	} else {
		for _, adj := range graph[from] {
			if _, in := vis[adj]; !in && adj != "start" {
				solve1(adj, graph, vis)
			}
		}
	}
	//path = path[:len(path)-1]
	if smallCave(from) {
		delete(vis, from)
	}
}

func dfsCaveLimit(from string, graph map[string][]string, vis map[string]int) {
	if smallCave(from) {
		vis[from]++
	}
	// count small caves seen twice
	seenTwice := 0
	for _, count := range vis {
		if count >= 2 {
			seenTwice++
		}
	}
	//path = append(path, from)
	if from == "end" && seenTwice < 2 {
		paths++
	} else {
		if seenTwice < 2 {
			for _, adj := range graph[from] {
				if count, _ := vis[adj]; count < 2 && adj != "start" {
					dfsCaveLimit(adj, graph, vis)
				}
			}
		}
	}
	//path = path[:len(path)-1]
	if smallCave(from) {
		vis[from]--
	}
}

func solve2(graph map[string][]string) {
	paths = 0
	vis := make(map[string]int)
	for node := range graph {
		if smallCave(node) {
			vis[node] = 0
		}
	}
	dfsCaveLimit("start", graph, vis)
}
func main() {
	graph := readFile(File)
	solve1("start", graph, make(map[string]bool))
	fmt.Println(paths)
	solve2(graph)
	fmt.Println(paths)
}
