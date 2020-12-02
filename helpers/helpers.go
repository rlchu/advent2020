package helpers

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// CleanInput reads txt file input based on name passed in and converts from []byte -> []int
func CleanInput(filename string) ([]int, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return []int{}, err
	}
	input := strings.Split(string(body), "\n")
	ci := make([]int, len(input))
	for i, val := range input {
		ci[i], err = strconv.Atoi(val)
		if err != nil {
			return []int{}, err
		}
	}
	return ci, nil
}
