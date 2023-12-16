package main

import (
	// "fmt"
	"regexp"
	"strings"

	// "strings"
	"testing"
)

var splitter = regexp.MustCompile("\r?\n")

// var testDataA = strings.Trim(`
// ##.#..#...##...#.
// ..####.#.####.#.#
// .....##.##..##.##
// ..###.##.#..####.
// ##..###..#..#..##
// ##.#####.####.###
// ..#.#.#.#....#.#.
// ####.#....##....#
// ####..####..####.
// ....#.#........#.
// ....##.#......#.#
// `, "\n")

var testDataA = strings.Trim(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`, "\n")

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	value, err := PartOneSolution(lines)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	expect := 405
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
	expect := 400
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
