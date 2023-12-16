package main

import (
	// "fmt"
	"regexp"
	"strings"

	// "strings"
	"testing"
)

var testData1 = strings.Trim(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`, "\n")

func TestPartOne(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testData1, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 8
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwo(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testData1, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 2286
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
