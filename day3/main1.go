package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getSymbols(lines []string) [][]bool {
	symbols := [][]bool{}
	for _, line := range lines {
		tmp := []bool{}
		for _, char := range line {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '.' {
				tmp = append(tmp, true)
			} else {
				tmp = append(tmp, false)
			}
		}
		symbols = append(symbols, tmp)
	}
	return symbols
}

func checkIfValidNumber(ids [][]int, symbols [][]bool) bool {
	maxX := len(symbols)
	maxY := len(symbols[0])
	for _, id := range ids {
		if id[1]+1 < maxY && symbols[id[0]][id[1]+1] ||
			id[1] != 0 && symbols[id[0]][id[1]-1] ||
			id[0]+1 < maxX && symbols[id[0]+1][id[1]] ||
			id[0] != 0 && symbols[id[0]-1][id[1]] ||
			id[1]+1 < maxY && id[0]+1 < maxX && symbols[id[0]+1][id[1]+1] ||
			id[1] != 0 && id[0] != 0 && symbols[id[0]-1][id[1]-1] ||
			id[0]+1 < maxX && id[1] != 0 && symbols[id[0]+1][id[1]-1] ||
			id[1]+1 < maxY && id[0] != 0 && symbols[id[0]-1][id[1]+1] {

			return true
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	var sum int = 0

	symbols := getSymbols(lines)
	for x, line := range lines {
		n := ""
		ids := [][]int{}
		for y, char := range line {
			if char >= '0' && char <= '9' {
				n = n + string(char)
				xy := []int{x, y}
				ids = append(ids, xy)
			}
			if (!(char >= '0' && char <= '9') || y+1 == len(line)) && n != "" {
				if checkIfValidNumber(ids, symbols) {
					num, _ := strconv.Atoi(n)
					fmt.Println(num)
					sum = sum + num

				}
				n = ""
				ids = [][]int{}
			}
		}

	}
	fmt.Println(sum)
}
