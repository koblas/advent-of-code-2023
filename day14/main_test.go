package main

import (
	// "fmt"
	"fmt"
	"regexp"
	"strings"

	// "strings"
	"testing"
)

var splitter = regexp.MustCompile("\r?\n")

var testDataA = strings.Trim(`
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`, "\n")

func XTestRotate(t *testing.T) {
	lines := Grid{
		"ABC",
		"...",
		"XYZ",
	}

	fmt.Println("Start")
	Dump(lines)
	fmt.Println("One")
	lines = rotateCW(lines)
	Dump(lines)
	fmt.Println("Two")
	lines = rotateCW(lines)
	Dump(lines)
	fmt.Println("Three")
	lines = rotateCW(lines)
	Dump(lines)
	fmt.Println("Four")
	lines = rotateCW(lines)
	Dump(lines)
}

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 136
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwoA(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataA, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 64
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
