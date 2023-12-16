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
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`, "\n")

var testDataB = strings.Trim(`
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`, "\n")

var testDataC = strings.Trim(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`, "\n")

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 2
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartOneB(t *testing.T) {
	lines := splitter.Split(testDataB, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 6
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwo(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataC, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 6
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
