package main

import (
	"fmt"
	"os"
	"strings"
)

func getScore(played []string, winning []string) int {
	sum := 0
	set := make(map[string]bool)
	for _, p := range played {
		set[p] = true
	}
	for _, w := range winning {
		if set[w] {
			if sum == 0 {
				sum = 1
			} else {
				sum = sum * 2
			}
		}
	}
	return sum
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	var sum int = 0

	for _, line := range lines {
		card := strings.Split(line, ":")[1]

		playedS := strings.Trim(strings.Split(card, "|")[0], " ")
		winningS := strings.Trim(strings.Split(card, "|")[1], " ")

		played := strings.Fields(playedS)
		winning := strings.Fields(winningS)

		count := getScore(played, winning)

		sum = sum + count

	}
	fmt.Println(sum)
}
