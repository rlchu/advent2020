package day07

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func checkSubBags(bag string, bagCapabilities map[string][]string) bool {
	for _, subBag := range bagCapabilities[bag] {
		if subBag == "shiny gold" {
			return true
		}
		if subBag == "no other bags" {
			return false
		}
		if checkSubBags(subBag, bagCapabilities) {
			return true
		}
	}

	return false
}

// Part1 ...
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	groups := strings.Split(string(body), "\n")

	re1 := regexp.MustCompile(`\w+ \w+`)
	re2 := regexp.MustCompile(`\d+ \w+ \w+`)
	bagCapabilities := make(map[string][]string)

	for _, group := range groups {
		subjectBag := re1.FindString(group)
		if len(re2.FindAllString(group, -1)) != 0 {
			for _, val := range re2.FindAllString(group, -1) {
				temp := strings.Split(val, " ")[1:]
				val = strings.Join(temp, " ")
				bagCapabilities[subjectBag] = append(bagCapabilities[subjectBag], val)
			}
		} else {
			bagCapabilities[subjectBag] = append(bagCapabilities[subjectBag], "no other bags")
		}
	}

	total := 0
	for bag := range bagCapabilities {
		if checkSubBags(bag, bagCapabilities) {
			total++
		}

	}

	return total, nil
}

// Part2 ...
// func Part2(filename string) (int, error) {
// 	body, _ := ioutil.ReadFile(filename)
// 	groups := strings.Split(string(body), "\n\n")

// 	return answer, nil
// }
