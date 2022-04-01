package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const File = "./src/day16/d16.txt"

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

var hexTable = map[int32]string{
	'0': "0000", '1': "0001", '2': "0010", '3': "0011", '4': "0100", '5': "0101", '6': "0110", '7': "0111",
	'8': "1000", '9': "1001", 'A': "1010", 'B': "1011", 'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
}

type Transmission struct {
	bits string
	rlen int
}

func hexToBin(hex string) string {
	var builder strings.Builder
	for _, code := range hex {
		builder.WriteString(hexTable[code])
	}
	return builder.String()
}

// Takes bits and return ex: takeBits(3) of 100101 returns 100 and the bits are now 101
func (t *Transmission) takeBits(amount int) string {
	bits := t.bits[:amount]
	t.bits = t.bits[amount:]
	t.rlen += amount
	return bits
}

func (t *Transmission) current() uint8 {
	return t.bits[0]
}

func (t *Transmission) literal() string {
	bits := t.takeBits(5)
	// Exclude fist bit
	return bits[1:]
}

func (t *Transmission) Version() uint64 {
	version, _ := strconv.ParseUint(t.takeBits(3), 2, 32)
	return version
}

func (t *Transmission) Id() uint64 {
	id, _ := strconv.ParseUint(t.takeBits(3), 2, 32)
	return id
}

func (t *Transmission) LengthType() uint64 {
	lengthType, _ := strconv.ParseUint(t.takeBits(1), 2, 32)
	return lengthType
}

func (t *Transmission) LiteralNumber() uint64 {
	var builder strings.Builder
	parseLiteral := true
	for parseLiteral {
		parseLiteral = t.current() != '0'
		literal := t.literal()
		builder.WriteString(literal)
	}
	parsed, _ := strconv.ParseUint(builder.String(), 2, 64)
	return parsed
}

func (t *Transmission) SubPackets() uint64 {
	numberSubPackets, _ := strconv.ParseUint(t.takeBits(11), 2, 64)
	return numberSubPackets
}

func (t *Transmission) Length() uint64 {
	length, _ := strconv.ParseUint(t.takeBits(15), 2, 64)
	return length
}

func sum(array []uint64) uint64 {
	var s uint64 = 0
	for _, v := range array {
		s += v
	}
	return s
}

func product(array []uint64) uint64 {
	var p uint64 = 1
	for _, v := range array {
		p *= v
	}
	return p
}

func min(array []uint64) uint64 {
	var mn uint64 = math.MaxUint64
	for _, v := range array {
		if v < mn {
			mn = v
		}
	}
	return mn
}

func max(array []uint64) uint64 {
	var mx uint64 = 0
	for _, v := range array {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func gt(array []uint64) uint64 {
	if array[0] > array[1] {
		return 1
	}
	return 0
}

func lt(array []uint64) uint64 {
	if array[0] < array[1] {
		return 1
	}
	return 0
}

func eq(array []uint64) uint64 {
	if array[0] == array[1] {
		return 1
	}
	return 0
}

func parse(t *Transmission, versionsAdd bool) uint64 {
	version := t.Version()
	id := t.Id()
	nums := make([]uint64, 0)
	if versionsAdd {
		nums = append(nums, version)
	}
	if id == 4 {
		literalNumber := t.LiteralNumber()
		if versionsAdd {
			return version
		}
		return literalNumber
	}
	lengthType := t.LengthType()
	if lengthType == 1 {
		nSubPackets := t.SubPackets()
		for nSubPackets > 0 {
			nums = append(nums, parse(t, versionsAdd))
			nSubPackets--
		}
	} else {
		length := t.Length()
		stop := t.rlen + int(length)
		for t.rlen < stop {
			nums = append(nums, parse(t, versionsAdd))
		}
	}
	if versionsAdd {
		return sum(nums)
	} else {
		switch id {
		case 0:
			return sum(nums)
		case 1:
			return product(nums)
		case 2:
			return min(nums)
		case 3:
			return max(nums)
		case 5:
			return gt(nums)
		case 6:
			return lt(nums)
		case 7:
			return eq(nums)
		default:
			return math.MaxUint64
		}
	}
}

func solve(hex string, part1 bool) uint64 {
	binString := hexToBin(hex)
	t := &Transmission{
		bits: binString,
	}
	return parse(t, part1)
}

func main() {
	input := readFile(File)
	fmt.Println(solve(input, true))
	fmt.Println(solve(input, false))
}
