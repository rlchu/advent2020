package day01

import (
	"testing"
)

type testCase struct {
	fileName string
	want     int64
}

var testData01 = []testCase{
	{"day01_example.txt", 514579},
	{"day01_input.txt", 970816},
}

var testData02 = []testCase{
	{"day01_example.txt", 241861950},
	{"day01_input.txt", 96047280},
}

func TestPart1(t *testing.T) {
	for _, val := range testData01 {
		answer, err := Part1(val.fileName)
		if err != nil {
			t.Fatalf("%s", err)
		}
		if answer != val.want {
			t.Fatalf("wanted: %v got: %v", val.want, answer)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, val := range testData02 {
		answer, err := Part2(val.fileName)
		if err != nil {
			t.Fatalf("%s", err)
		}
		if answer != val.want {
			t.Fatalf("wanted: %v got: %v", val.want, answer)
		}
	}
}
