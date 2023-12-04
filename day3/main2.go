package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numData struct {
	yMin int
	yMax int
	n    int
	x    int
}

func getStars(lines []string) [][][]int {
	symbols := [][][]int{}
	for _, line := range lines {
		tmp := [][]int{}
		for _, char := range line {
			if char == '*' {
				tmp = append(tmp, []int{})
			} else {
				tmp = append(tmp, []int{-1, -1, -1})
			}
		}
		symbols = append(symbols, tmp)
	}
	return symbols
}

func addToSymbols(numD numData, symbolsP *[][][]int) {
	symbols := *symbolsP
	maxX := len(symbols) - 1
	maxY := len((symbols)[0]) - 1
	for x := max(0, numD.x-1); x <= min(maxX, numD.x+1); x++ {
		for y := max(0, numD.yMin-1); y <= min(maxY, numD.yMax+1); y++ {
			if len(symbols[x][y]) < 2 {
				symbols[x][y] = append(symbols[x][y], numD.n)
			}
		}
	}
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	// var sum int = 0

	symbols := getStars(lines)
	for x, line := range lines {
		n := ""
		numD := numData{yMin: -1, yMax: -1, x: x, n: -1}
		for y, char := range line {
			if char >= '0' && char <= '9' {
				n = n + string(char)
				if numD.yMin == -1 {
					numD.yMin = y
				}
				numD.yMax = y
			}
			if (!(char >= '0' && char <= '9') || y+1 == len(line)) && n != "" {
				num, _ := strconv.Atoi(n)
				numD.n = num
				addToSymbols(numD, &symbols)
				n = ""
				numD = numData{yMin: -1, yMax: -1, x: x, n: -1}

			}
		}
	}
	sum := 0
	for _, line := range symbols {
		for _, arr := range line {
			if len(arr) == 2 {
				sum = sum + arr[0]*arr[1]
			}
		}
	}
	fmt.Println(sum)
}
