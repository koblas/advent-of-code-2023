package main

import (
	// "fmt"
	"regexp"
	"strings"

	// "strings"
	"testing"
)

var splitter = regexp.MustCompile("\r?\n")

var testData = strings.Trim(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`, "\n")

func TestPartOne(t *testing.T) {
	lines := splitter.Split(testData, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 6440
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwo(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testData, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 5905
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
