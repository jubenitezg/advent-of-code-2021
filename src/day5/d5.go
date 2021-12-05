package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day5/d5.txt"

type Point struct {
	x int
	y int
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

func readFile(fileName string) []*Point {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var points []*Point
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " -> ")
		pstr1, pstr2 := strings.Split(split[0], ","), strings.Split(split[1], ",")
		x1, _ := strconv.Atoi(pstr1[0])
		y1, _ := strconv.Atoi(pstr1[1])
		x2, _ := strconv.Atoi(pstr2[0])
		y2, _ := strconv.Atoi(pstr2[1])
		points = append(points, &Point{
			x: x1,
			y: y1,
		}, &Point{
			x: x2,
			y: y2,
		})
	}

	return points
}

func addCount(overlaps map[Point]int, np Point) {
	if _, ok := overlaps[np]; !ok {
		overlaps[np] = 0
	}
	overlaps[np]++
}

func countOverlaps(points []*Point, diagonals bool) int {
	overlaps := make(map[Point]int)
	for i := 0; i < len(points); i += 2 {
		p1, p2 := points[i], points[i+1]
		if p1.x == p2.x {
			for ii := min(p1.y, p2.y); ii <= max(p1.y, p2.y); ii++ {
				addCount(overlaps, Point{x: p1.x, y: ii})
			}
		} else if p1.y == p2.y {
			for ii := min(p1.x, p2.x); ii <= max(p1.x, p2.x); ii++ {
				addCount(overlaps, Point{x: ii, y: p1.y})
			}
		} else if diagonals {
			currx, curry := p1.x, p1.y
			addCount(overlaps, Point{x: currx, y: curry})
			for currx != p2.x || curry != p2.y {
				if p2.x > currx {
					currx++
				} else if p2.x < currx {
					currx--
				}
				if p2.y > curry {
					curry++
				} else if p2.y < curry {
					curry--
				}
				addCount(overlaps, Point{x: currx, y: curry})
			}
		}
	}
	count := 0
	for _, val := range overlaps {
		if val > 1 {
			count++
		}
	}
	return count
}

func solve1(points []*Point) int {
	return countOverlaps(points, false)
}

func solve2(points []*Point) int {
	return countOverlaps(points, true)
}

func main() {
	points := readFile(File)
	fmt.Println(solve1(points))
	fmt.Println(solve2(points))
}
