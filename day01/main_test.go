package main

import (
	// "fmt"
	"regexp"
	// "strings"
	"testing"
)

var testData1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var testData2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestPartOne(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testData1, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 142
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwo(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testData2, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 281
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
