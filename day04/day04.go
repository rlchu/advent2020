package day04

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var byr = regexp.MustCompile(`byr:\d{4}\b`)
var iyr = regexp.MustCompile(`iyr:\d{4}\b`)
var eyr = regexp.MustCompile(`eyr:\d{4}\b`)
var hgt = regexp.MustCompile(`hgt:(\d+)(cm|in)`)
var hcl = regexp.MustCompile(`hcl:#[0-9a-f]{6}\b`)
var ecl = regexp.MustCompile(`ecl:amb|ecl:blu|ecl:brn|ecl:gry|ecl:grn|ecl:hzl|ecl:oth`)
var pid = regexp.MustCompile(`pid:\d{9}\b`)

func checkForRequiredFields(passport string) bool {
	requiredFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	for _, field := range requiredFields {
		if strings.Contains(passport, field) == false {
			return false
		}
	}
	return true
}

func yearConstraintsCheck(re *regexp.Regexp, passport string, min, max int) bool {
	submatch := re.FindStringSubmatch(passport)
	year, _ := strconv.Atoi(strings.Split(submatch[0], ":")[1])

	if year <= max && year >= min {
		return false
	}
	return true
}

func checkHeight(passport string) bool {
	if !hgt.MatchString(passport) {
		return true
	}
	submatch := hgt.FindStringSubmatch(passport)
	height := strings.Split(submatch[0], ":")[1]
	re := regexp.MustCompile(`(\d+)([c|i])`)
	match := re.FindStringSubmatch(height)
	magnitude, unit := match[1], match[2]
	m, _ := strconv.Atoi(magnitude)
	if unit == "c" && (m <= 193 && m >= 150) {
		return false
	} else if unit == "i" && (m <= 76 && m >= 59) {
		return false
	}
	return true
}

func checkForValidEntries(passport string) bool {
	if checkForRequiredFields(passport) == false {
		return false
	}

	if yearConstraintsCheck(byr, passport, 1920, 2002) {
		return false
	}

	if yearConstraintsCheck(iyr, passport, 2010, 2020) {
		return false
	}

	if yearConstraintsCheck(eyr, passport, 2020, 2030) {
		return false
	}

	if checkHeight(passport) {
		return false
	}

	if !hcl.MatchString(passport) {
		return false
	}

	if !ecl.MatchString(passport) {
		return false
	}

	if !pid.MatchString(passport) {
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

	return validCount, nil
}
