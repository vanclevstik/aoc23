package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkString(s string) (bool, bool) {
	validStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, validString := range validStrings {
		if strings.HasPrefix(validString, s) {
			if s == validString {
				return true, true
			}
			return true, false
		}

	}
	return false, false
}
func main() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	var sum int = 0
	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"
	m["four"] = "4"
	m["five"] = "5"
	m["six"] = "6"
	m["seven"] = "7"
	m["eight"] = "8"
	m["nine"] = "9"

	//lines := [1]string{"oneightwoneight"}
	for _, line := range lines {
		var first *string = nil
		var last string

		var d string
		for _, char := range line {
			if char >= '0' && char <= '9' {
				c := string(char)
				if first == nil {
					first = &c
				}
				last = c

			} else {
				d = d + string(char)
				valid, exact := checkString(d)
				if exact {
					c := m[d]
					if first == nil {
						first = &c
					}
					last = c
					switch d {
					case "seven":
						d = "n"
					case "eight":
						d = "t"
					case "two":
						d = "o"
					case "three":
						d = "e"
					case "five":
						d = "e"
					case "nine":
						d = "e"
					case "one":
						d = "e"
					default:
						d = ""
					}

				} else if !valid {
					for d != "" && !valid {
						//fmt.Println(d)
						if len(d) == 1 {
							d = ""
						} else {
							d = d[1:]
						}
						valid, _ = checkString(d)
					}
				}
				//fmt.Println("continue")
			}

		}
		val, _ := strconv.Atoi(*first + last)
		fmt.Printf(" word: %s, val: %d\n", line, val)
		sum = sum + val
	}

	fmt.Println(sum)
}

// string is 0   string is valid
//  yes          yes                   stop
//  yes          no                    stop
//  no           yes                   stop
//  no           no				keep goin
