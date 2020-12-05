package day04

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func checkForRequiredFields(passport string) bool {
	requiredFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	for _, field := range requiredFields {
		if strings.Contains(passport, field) == false {
			return false
		}
	}
	return true
}

func checkForValidEntries(passport string) bool {
	if checkForRequiredFields(passport) == false {
		return false
	}
	byr := regexp.MustCompile(`byr:\w{4}`)
	e := byr.FindStringSubmatch(passport)
	// if strings.Split(e[0], ":")[1] > 2002 || strings.Split(e[0], ":")[1] < 1920 {
	// 	return false
	// }
	fmt.Println(strings.Split(e[0], ":")[1])

	// for _, field := range requiredFields {
	// 	if strings.Contains(passport, field) == false {
	// 		return false
	// 	}
	// }
	return true
}

// Part1 ..
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n\n")
	validCount := 0
	for _, passport := range input {
		if checkForRequiredFields(passport) {
			validCount++
		}
	}
	return validCount, nil
}

// Part2 ...
func Part2(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n\n")
	validCount := 0
	for _, passport := range input {
		if checkForValidEntries(passport) {
			validCount++
		}
	}

	return 0, nil
}
