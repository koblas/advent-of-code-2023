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
-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`, "\n")

// var testDataA = strings.Trim(`
// .....
// .S-7.
// .|.|.
// .L-J.
// .....
// `, "\n")

var testDataB = strings.Trim(`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`, "\n")

var testDataC = strings.Trim(`
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`, "\n")

var testDataD = strings.Trim(`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`, "\n")

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 4
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
	expect := 8
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwoC(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataC, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 4
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwoD(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataD, -1)
	value, err := PartTwoSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 10
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
