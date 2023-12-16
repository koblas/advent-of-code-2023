package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fd)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	value, err := PartOneSolution(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1: ", value)

	values, err := PartTwoSolution(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2: ", values)
}

type Group []int

type Board struct {
	Order []Group
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var values Group
		parts := strings.Split(line, " ")

		for _, val := range parts {
			v, err := strconv.Atoi(val)
			if err != nil {
				return Board{}, err
			}
			values = append(values, v)
		}

		result.Order = append(result.Order, values)
	}

	return result, nil
}

func allZeros(set Group) bool {
	for _, v := range set {
		if v != 0 {
			return false
		}
	}
	return true
}

func runSolutionOne(set Group) int {
	rows := []Group{set}

	sum := set[len(set)-1]
	for !allZeros(rows[len(rows)-1]) {
		//
		next := Group{}
		lastRow := rows[len(rows)-1]
		for idx, v := range lastRow[1:] {
			next = append(next, v-lastRow[idx])
		}
		rows = append(rows, next)
		sum += next[len(next)-1]
	}

	return sum
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, set := range input.Order {
		sum += runSolutionOne(set)
	}

	return sum, nil
}

func runSolutionTwo(set Group) int {
	rows := []Group{set}

	for !allZeros(rows[len(rows)-1]) {
		//
		next := Group{}
		lastRow := rows[len(rows)-1]
		for idx, v := range lastRow[1:] {
			next = append(next, v-lastRow[idx])
		}
		rows = append(rows, next)
	}

	previous := 0
	for idx := len(rows) - 1; idx >= 0; idx -= 1 {
		previous = rows[idx][0] - previous
	}

	return previous
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, set := range input.Order {
		sum += runSolutionTwo(set)
	}

	return sum, nil
}
