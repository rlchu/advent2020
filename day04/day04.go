package day04

import (
	"io/ioutil"
	"strings"
)

func checkForRequired(passport string) bool {
	requiredFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	for _, field := range requiredFields {
		if strings.Contains(passport, field) == false {
			return false
		}
	}
	return true
}

// Part1 ..
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n\n")
	validCount := 0
	for _, passport := range input {
		if checkForRequired(passport) {
			validCount++
		}
	}
	return validCount, nil
}

// Part2 ...
// func Part2(filename s)tring) (int, error) {
// 	stringGrid, err := parseInput(filename)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }
