package day06

import (
	"io/ioutil"
	"strings"
)

func dedupe(str string) string {
	letters := make(map[string]bool)
	splitString := strings.Split(str, "")
	dedupedSplitString := []string{}

	for _, entry := range splitString {
		if _, value := letters[entry]; !value {
			letters[entry] = true
			dedupedSplitString = append(dedupedSplitString, entry)
		}
	}

	return strings.Join(dedupedSplitString, "")
}

// Part1 counts the number of questions to which anyone answered "yes" and returns the sum across groups
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	groups := strings.Split(string(body), "\n\n")
	answer := 0

	for _, group := range groups {
		group = strings.Replace(group, " ", "", -1)
		group = strings.Replace(group, "\n", "", -1)
		answer += len(dedupe(group))
	}

	return answer, nil
}

// Part2 counts the number of questions to which everyone answered "yes" and returns sum of those counts
func Part2(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	groups := strings.Split(string(body), "\n\n")
	answer := 0

	for _, group := range groups {
		frequencies := make(map[string]int)
		groupParsed := strings.Split(strings.Replace(group, " ", "", -1), "\n")
		groupLength := len(groupParsed)

		for _, individualAnswers := range groupParsed {
			for _, letter := range strings.Split(individualAnswers, "") {
				frequencies[letter]++
			}
		}

		for _, v := range frequencies {
			if v == groupLength {
				answer++
			}
		}
	}

	return answer, nil
}
