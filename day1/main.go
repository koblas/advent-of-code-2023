package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func ParseInput(lines []string) ([]string, error) {
	first := 0
	last := len(lines)
	if lines[0] == "" {
		first = 1
	}
	if lines[len(lines)-1] == "" {
		last = len(lines) - 1
	}
	return lines[first:last], nil
}

var vtable = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func toValue(value string) string {
	return vtable[value]
}

func PartOneSolution(lines []string) (int, error) {
	expr := regexp.MustCompile("^[^0-9]*(\\d).*?(\\d)?[^0-9]*$")
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range input {
		parsed := expr.FindStringSubmatch(line)

		var num string
		if len(parsed) == 3 {
			if parsed[2] == "" {
				num = toValue(parsed[1]) + toValue(parsed[1])
			} else {
				num = toValue(parsed[1]) + toValue(parsed[2])
			}
		} else {
			return 0, fmt.Errorf("invalid line: %s", line)
		}

		v, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}

		sum += v
	}

	return sum, nil
}

func PartTwoSolution(lines []string) (int, error) {
	exprFirst := regexp.MustCompile("^.*?(\\d|one|two|three|four|five|six|seven|eight|nine)")
	exprLast := regexp.MustCompile("^.*(\\d|one|two|three|four|five|six|seven|eight|nine).*?$")

	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range input {
		first := exprFirst.FindStringSubmatch(line)
		if len(first) != 2 {
			return 0, fmt.Errorf("first no match: %s", line)
		}
		last := exprLast.FindStringSubmatch(line)
		if len(last) != 2 {
			return 0, fmt.Errorf("last no match: %s", line)
		}

		num := toValue(first[1]) + toValue(last[1])

		v, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}

		sum += v
	}

	return sum, nil
}
