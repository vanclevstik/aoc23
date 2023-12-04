package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func coloursSplit(s string) map[string]int {
	m := make(map[string]int)
	for _, colour := range strings.Split(s, ",") {
		//fmt.Println(colour)
		if strings.Contains(colour, "green") {
			m["green"], _ = strconv.Atoi(strings.Split(colour, " ")[1])
		} else if strings.Contains(colour, "blue") {
			m["blue"], _ = strconv.Atoi(strings.Split(colour, " ")[1])
		} else if strings.Contains(colour, "red") {
			m["red"], _ = strconv.Atoi(strings.Split(colour, " ")[1])
		}
	}
	return m

}

func extractGames(line string) (int, []map[string]int) {
	// line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	id, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])

	games := []map[string]int{}
	for _, game := range strings.Split(strings.Split(line, ":")[1], ";") {
		games = append(games, coloursSplit(game))
	}
	return id, games
}

func validGames(games []map[string]int) bool {
	for _, game := range games {
		if game["green"] > 13 || game["blue"] > 14 || game["red"] > 12 {
			return false
		}
	}
	return true
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
		id, extractedGames := extractGames(line)

		if validGames(extractedGames) {
			sum = sum + id
		}
	}

	fmt.Println(sum)
}
