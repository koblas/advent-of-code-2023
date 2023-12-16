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

type Race struct {
	time   int
	record int
}

func ParseInput(lines []string) ([]Race, Race, error) {
	single := Race{}

	if len(lines) == 0 {
		return nil, single, fmt.Errorf("bad input")
	}

	var result []Race

	rex := regexp.MustCompile(" +")

	for _, line := range lines {
		parts := rex.Split(line, -1)
		accum := ""
		for idx, value := range parts[1:] {
			if value == "" {
				continue
			}

			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, single, err
			}

			accum += value
			if parts[0] == "Time:" {
				result = append(result, Race{time: num})
			} else if parts[0] == "Distance:" {
				result[idx].record = num
			}
		}
		num, err := strconv.Atoi(accum)
		if err != nil {
			return nil, single, err
		}
		if parts[0] == "Time:" {
			single.time = num
		} else {
			single.record = num
		}
	}

	return result, single, nil
}

func computeDist(holdTime, raceTime int) int {
	remain := raceTime - holdTime

	return holdTime * remain
}

func PartOneSolution(lines []string) (int, error) {
	input, _, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, race := range input {
		ways := 0
		for time := 0; time <= race.time; time += 1 {
			dist := computeDist(time, race.time)
			if dist > race.record {
				ways += 1
			}
		}
		if sum == 0 {
			sum = ways
		} else {
			sum = sum * ways
		}
	}

	return sum, nil
}

func PartTwoSolution(lines []string) (int, error) {
	_, race, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	ways := 0
	for time := 0; time <= race.time; time += 1 {
		dist := computeDist(time, race.time)
		if dist > race.record {
			ways += 1
		}
	}

	return ways, nil
}
