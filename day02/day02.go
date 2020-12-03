package day02

import (
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

// Part1 returns # of valid passwords
func Part1(filename string) (int, error) {
	passwordLine, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
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
	return finalCount, nil
}

// Part2 returns # of valid passwords based on positions rule
func Part2(filename string) (int, error) {
	passwordLine, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, line := range passwordLine {
		if (string(line.password[line.low-1]) == line.letter) != (string(line.password[line.high-1]) == line.letter) {
			count++
		}
	}
	return count, nil
}
