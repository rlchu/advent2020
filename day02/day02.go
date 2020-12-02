package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passwordLine struct {
	low      int
	high     int
	letter   string
	password string
	count    int
}

// parseInput translates password input file into []passwordLine
func parseInput(filename string) ([]passwordLine, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return []passwordLine{}, err
	}
	input := strings.Split(string(body), "\n")
	ci := make([]passwordLine, len(input))
	for i, val := range input {
		s := strings.Split(val, " ")
		ci[i].letter = string(s[1][0])
		ci[i].password = s[2]

		n := strings.Split(s[0], "-")
		ci[i].low, err = strconv.Atoi(string(n[0]))
		ci[i].high, err = strconv.Atoi(string(n[1]))
	}
	return ci, nil
}

func main2() {
	passwordLine, _ := parseInput("day02_input.txt")
	for k, line := range passwordLine {
		passwordLine[k].count = 0
		for i, passwordLetter := range line.password {
			contLetterCount := 0
			if line.letter == string(passwordLetter) {
				for _, nextLetter := range []byte(line.password)[i:] {
					if string(nextLetter) == line.letter {
						contLetterCount++
						continue
					} else {
						break
					}
				}
			}
			if contLetterCount >= passwordLine[k].count {
				passwordLine[k].count = contLetterCount
			}
		}
	}
	finalCount := 0
	for _, line := range passwordLine {
		if line.low <= line.count && line.count <= line.high {
			finalCount++
		}
	}
	fmt.Println(finalCount)
}

func main3() {
	passwordLine, _ := parseInput("day02_input.txt")
	// passwordLine, _ := parseInput("day02_example.txt")
	for k, line := range passwordLine {
		passwordLine[k].count = 0
		for _, passwordLetter := range line.password {
			if line.letter == string(passwordLetter) {
				passwordLine[k].count++
			}

		}
	}
	finalCount := 0
	for _, line := range passwordLine {
		if line.low <= line.count && line.count <= line.high {
			finalCount++
		}
	}
	fmt.Println(finalCount)
}

func main() {
	passwordLine, _ := parseInput("day02_input.txt")
	// passwordLine, _ := parseInput("day02_example.txt")
	count := 0
	for _, line := range passwordLine {
		if (string(line.password[line.low-1]) == line.letter) != (string(line.password[line.high-1]) == line.letter) {
			count++
		}
	}

	fmt.Println(count)
}
