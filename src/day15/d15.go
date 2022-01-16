package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day15/d15.txt"

func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input [][]int
	for scanner.Scan() {
		snumbers := strings.Split(scanner.Text(), "")
		numbers := make([]int, len(snumbers))
		for i, snum := range snumbers {
			atoi, _ := strconv.Atoi(snum)
			numbers[i] = atoi
		}
		input = append(input, numbers)
	}

	return input
}

type Point struct {
	i int
	j int
}

type Item struct {
	point    *Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

var di = [4]int{-1, 0, 1, 0}
var dj = [4]int{0, 1, 0, -1}

func dijkstra(si, sj int, riskMap [][]int) int {
	n := len(riskMap)
	m := len(riskMap[0])
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist[i] = append(dist[i], math.MaxInt32/2)
		}
	}
	dist[si][sj] = 0
	pq := make(PriorityQueue, 1)
	idx := 0
	pq[0] = &Item{
		point:    &Point{si, sj},
		priority: dist[si][sj],
		index:    idx,
	}
	idx++
	heap.Init(&pq)
	for pq.Len() > 0 {
		topNode := heap.Pop(&pq).(*Item)
		u := topNode.point.i
		v := topNode.point.j
		d := topNode.priority
		for k := 0; k < 4; k++ {
			newU := u + di[k]
			newV := v + dj[k]
			if newU < 0 || newU >= n || newV < 0 || newV >= m {
				continue
			}
			if d+riskMap[newU][newV] < dist[newU][newV] {
				dist[newU][newV] = d + riskMap[newU][newV]
				heap.Push(&pq, &Item{
					point:    &Point{newU, newV},
					priority: dist[newU][newV],
					index:    idx,
				})
				idx++
			}
		}

	}

	return dist[n-1][m-1]
}

func updateRiskMap(riskMap [][]int) [][]int {
	n := len(riskMap)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// skip tile 0,0
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				// On most left side tile look up
				for currI := n * (i - 1); currI < n*i; currI++ {
					newRow := make([]int, n)
					idx := 0
					for currJ := 0; currJ < n; currJ++ {
						newVal := (riskMap[currI][currJ] + 1) % 10
						if newVal == 0 {
							newVal = 1
						}
						newRow[idx] = newVal
						idx++
					}
					riskMap = append(riskMap, newRow)
				}
			} else {
				// look left
				for currI := n * i; currI < n*(i+1); currI++ {
					newRow := make([]int, n)
					idx := 0
					for currJ := n * (j - 1); currJ < n*j; currJ++ {
						newVal := (riskMap[currI][currJ] + 1) % 10
						if newVal == 0 {
							newVal = 1
						}
						newRow[idx] = newVal
						idx++
					}
					for _, newNum := range newRow {
						riskMap[currI] = append(riskMap[currI], newNum)
					}
				}
			}
		}
	}
	return riskMap
}

func solve1(riskMap [][]int) int {
	return dijkstra(0, 0, riskMap)
}

func solve2(riskMap [][]int) int {
	riskMap = updateRiskMap(riskMap)
	return dijkstra(0, 0, riskMap)
}

func main() {
	input := readFile(File)
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
