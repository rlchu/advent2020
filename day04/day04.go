package day04

import (
	"io/ioutil"
	"regexp"
	"strconv"
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
	byrYear, _ := strconv.Atoi(strings.Split(e[0], ":")[1])

	if byrYear > 2002 || byrYear < 1920 {
		return false
	}

	iyr := regexp.MustCompile(`iyr:\w{4}`)
	e = iyr.FindStringSubmatch(passport)
	iyrYear, _ := strconv.Atoi(strings.Split(e[0], ":")[1])

	if iyrYear > 2020 || iyrYear < 2010 {
		return false
	}

	eyr := regexp.MustCompile(`eyr:\w{4}`)
	e = eyr.FindStringSubmatch(passport)
	eyrYear, _ := strconv.Atoi(strings.Split(e[0], ":")[1])

	if eyrYear > 2030 || eyrYear < 2020 {
		return false
	}

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
