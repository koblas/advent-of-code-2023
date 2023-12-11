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

type RangeMap struct {
	source      int
	destination int
	count       int
}

type Almanac struct {
	seeds []int

	maps map[string][]RangeMap
}

func strsToInts(values []string) []int {
	result := make([]int, len(values))

	for idx, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		result[idx] = num
	}

	return result
}

func ParseInput(lines []string) (*Almanac, error) {
	almanac := Almanac{
		maps: map[string][]RangeMap{},
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("Bad input")
	}

	parts := strings.Split(lines[0], " ")
	if parts[0] != "seeds:" {
		return nil, fmt.Errorf("Bad seed line")
	}
	almanac.seeds = strsToInts(parts[1:])

	group := ""
	for _, line := range lines[1:] {
		// fmt.Println(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if parts[0][0] >= '0' && parts[0][0] <= '9' {
			if len(parts) != 3 {
				return nil, fmt.Errorf("wrong count")
			}
			ranges := strsToInts(parts)

			item := RangeMap{
				source:      ranges[1],
				destination: ranges[0],
				count:       ranges[2],
			}

			almanac.maps[group] = append(almanac.maps[group], item)
		} else {
			group = parts[0]
			almanac.maps[group] = []RangeMap{}
		}
	}

	return &almanac, nil
}

func find(value int, mapping []RangeMap) int {
	for _, item := range mapping {
		if value >= item.source && value < item.source+item.count {
			return (value - item.source) + item.destination
		}
	}
	return value
}

func mapIt(seed int, almanac *Almanac) int {
	pos := find(seed, almanac.maps["seed-to-soil"])
	pos = find(pos, almanac.maps["soil-to-fertilizer"])
	pos = find(pos, almanac.maps["fertilizer-to-water"])
	pos = find(pos, almanac.maps["water-to-light"])
	pos = find(pos, almanac.maps["light-to-temperature"])
	pos = find(pos, almanac.maps["temperature-to-humidity"])
	pos = find(pos, almanac.maps["humidity-to-location"])

	return pos
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	bestPos := -1
	for _, seed := range input.seeds {
		pos := mapIt(seed, input)
		if bestPos == -1 || pos < bestPos {
			bestPos = pos
		}
	}

	return bestPos, nil
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	bestPos := -1
	for i := 0; i < len(input.seeds); i += 2 {
		for idx := 0; idx < input.seeds[i+1]; idx++ {
			seed := input.seeds[i] + idx
			pos := mapIt(seed, input)
			if bestPos == -1 || pos < bestPos {
				bestPos = pos
			}
		}
	}

	return bestPos, nil
}
