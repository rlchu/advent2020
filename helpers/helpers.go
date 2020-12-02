package helpers

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// CleanInput reads txt file input based on name passed in and converts from []byte -> []string -> []int64
func CleanInput(filename string) []int {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n")
	ci := make([]int, len(input))
	for i, val := range input {
		ci[i], _ = strconv.Atoi(val)
	}
	return ci
}
