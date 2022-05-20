package main

import (
	"bufio"
	"fmt"
	"os"
)

const File = "./src/day17/d17.txt"

func readFile(fileName string) (int, int, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()
	var xi, xf, yi, yf int
	fmt.Sscanf(template, "target area: x=%d..%d, y=%d..%d", &xi, &xf, &yi, &yf)

	return xi, xf, yi, yf
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMaxAbs(a, b int) int {
	a = abs(a)
	b = abs(b)
	if a > b {
		return a
	}
	return b
}

func inside(xi, xf, yi, yf, vx, vy int) bool {
	x := 0
	y := 0
	for {
		// Overshoot
		if x > xf || y < yi {
			return false
		}
		// can't reach if no velocity is left
		if vx == 0 && x < xi {
			return false
		}
		if xi <= x && x <= xf && yi <= y && y <= yf {
			return true
		}
		x += vx
		y += vy
		// steps
		vy--
		if vx > 0 {
			vx--
		}
	}
}

func solve1(yi, yf int) int {
	maxYtarget := getMaxAbs(yi, yf)
	maxY := (maxYtarget - 1) * maxYtarget / 2

	return maxY
}

func solve2(xi, xf, yi, yf int) int {
	yMax := getMaxAbs(yi, yf)
	total := 0
	for vx := 0; vx <= xf; vx++ {
		for vy := -yMax; vy <= yMax; vy++ {
			if inside(xi, xf, yi, yf, vx, vy) {
				total++
			}
		}
	}

	return total
}

func main() {
	xi, xf, yi, yf := readFile(File)
	fmt.Println(solve1(yi, yf))
	fmt.Println(solve2(xi, xf, yi, yf))
}
