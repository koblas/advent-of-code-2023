package main

import (
	// "fmt"
	"regexp"
	"strings"

	// "strings"
	"testing"
)

var splitter = regexp.MustCompile("\r?\n")

var testDataA = strings.Trim(`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`, "\n")

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 114
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwo(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataA, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 2
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
