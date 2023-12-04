package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	var sum int = 0
	for _, line := range lines {
		var first *rune = nil
		var last rune

		for _, char := range line {
			if char >= '0' && char <= '9' {
				var c rune = char
				if first == nil {
					first = &c
				}
				last = char
			}

		}
		val, _ := strconv.Atoi(string(*first) + string(last))
		sum = sum + val
	}

	fmt.Println(sum)
}
