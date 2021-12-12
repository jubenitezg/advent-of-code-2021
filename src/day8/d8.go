package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const File = "./src/day8/d8.txt"

// This problem and this code (╯°□°）╯︵ ┻━┻

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func solve1(input []string) int {
	total := 0
	for _, line := range input {
		outputValues := strings.Fields(strings.Split(line, " | ")[1])
		for _, value := range outputValues {
			n := len(value)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				total++
			}
		}
	}

	return total
}

func in(findChar int32, str string) bool {
	for _, ch := range str {
		if findChar == ch {
			return true
		}
	}
	return false
}

func equals(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, ch := range str2 {
		if !in(ch, str1) {
			return false
		}
	}
	return true
}

func makesString(make, from string) bool {
	makes := true
	for _, ch := range make {
		makes = makes && in(ch, from)
	}

	return makes
}

func solve2(input []string) int {
	final := 0
	for _, line := range input {
		numbers := make([]string, 10)
		decode := strings.Fields(strings.Split(line, " | ")[0])
		for _, outputValue := range decode {
			n := len(outputValue)
			switch n {
			case 2:
				numbers[1] = outputValue
			case 3:
				numbers[7] = outputValue
			case 4:
				numbers[4] = outputValue
			case 7:
				numbers[8] = outputValue
			}
		}
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 5 {
				if makesString(numbers[7], decodeValue) {
					numbers[3] = decodeValue
					break
				}
			}
		}
		// The two letters not found previously are the horizontal ones
		horizontal := ""
		for _, ch := range numbers[3] {
			if !in(ch, numbers[7]) {
				horizontal += string(ch)
			}
		}
		// From three, we can find 9, contains 3 which contains 7
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 6 {
				if makesString(numbers[3], decodeValue) {
					numbers[9] = decodeValue
					break
				}
			}
		}
		// With 4 find middle horizontal letter
		var middleHorizontal int32
		for _, ch := range horizontal {
			if in(ch, numbers[4]) {
				middleHorizontal = ch
				break
			}
		}
		// Zero is the only one which doesn't have a middle line
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 6 {
				if !in(middleHorizontal, decodeValue) {
					numbers[0] = decodeValue
					break
				}
			}
		}
		// Six is just not 0 and not 9
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 6 {
				if decodeValue != numbers[0] && decodeValue != numbers[9] {
					numbers[6] = decodeValue
					break
				}
			}
		}
		// To get 5 we need part of 4
		partOfFour := ""
		for _, ch := range numbers[4] {
			if !in(ch, numbers[1]) {
				partOfFour += string(ch)
			}
		}
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 5 {
				if makesString(partOfFour, decodeValue) {
					numbers[5] = decodeValue
					break
				}
			}
		}
		// 2 is just not 3 and not 5
		for _, decodeValue := range decode {
			n := len(decodeValue)
			if n == 5 {
				if decodeValue != numbers[5] && decodeValue != numbers[3] {
					numbers[2] = decodeValue
					break
				}
			}
		}
		number := 0
		for _, outputValue := range strings.Fields(strings.Split(line, " | ")[1]) {
			for num, strValue := range numbers {
				if equals(strValue, outputValue) {
					number = (number * 10) + num
				}
			}
		}
		final += number
	}

	return final
}

func main() {
	fmt.Println(solve1(readFile(File)))
	fmt.Println(solve2(readFile(File)))
}
