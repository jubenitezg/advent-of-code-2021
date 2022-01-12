package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day13/d13.txt"

type Point struct {
	x int
	y int
}

type Instruction struct {
	axis rune
	fold int
}

func readFile(fileName string) (map[Point]bool, []Instruction) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := make(map[Point]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points[Point{x, y}] = true
	}
	instructions := make([]Instruction, 0)
	var axis rune
	var fold int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "fold along %c=%d", &axis, &fold)
		instructions = append(instructions, Instruction{axis, fold})
	}

	return points, instructions
}

func solve1(points map[Point]bool, instructions []Instruction) int {
	// only fold instruction
	instruction := instructions[0]
	for point := range points {
		if instruction.axis == 'y' && point.y >= instruction.fold {
			newY := (instruction.fold << 1) - point.y
			points[Point{point.x, newY}] = true
			delete(points, point)
		} else if instruction.axis == 'x' && point.x >= instruction.fold {
			newX := (instruction.fold << 1) - point.x
			points[Point{newX, point.y}] = true
			delete(points, point)
		}
	}
	return len(points)
}
func solve2(points map[Point]bool, instructions []Instruction) {
	for _, instruction := range instructions {
		for point := range points {
			if instruction.axis == 'y' && point.y >= instruction.fold {
				newY := (instruction.fold << 1) - point.y
				points[Point{point.x, newY}] = true
				delete(points, point)
			} else if instruction.axis == 'x' && point.x >= instruction.fold {
				newX := (instruction.fold << 1) - point.x
				points[Point{newX, point.y}] = true
				delete(points, point)
			}
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 50; j++ {
			if points[Point{x: j, y: i}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}

func main() {
	points, instructions := readFile(File)
	fmt.Println(solve1(points, instructions))
	solve2(points, instructions)
}
