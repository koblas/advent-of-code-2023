package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

type Game struct {
	red   int
	green int
	blue  int
}

func ParseInput(lines []string) ([][]Game, error) {
	expr := regexp.MustCompile(`(\d+) (blue|red|green)`)
	result := [][]Game{}

	for i, line := range lines {
		parts := strings.Split(line, ";")

		round := []Game{}
		for _, part := range parts {
			values := expr.FindAllStringSubmatch(part, 100)

			game := Game{}
			for _, item := range values {
				if len(item) != 3 {
					return nil, fmt.Errorf("not enoguh values, line=%d %v %v", i, line, values)
				}
				val, err := strconv.Atoi(item[1])
				if err != nil {
					return nil, fmt.Errorf("bad number, line=%d %v", i, err)
				}

				switch item[2] {
				case "red":
					game.red = val
				case "green":
					game.green = val
				case "blue":
					game.blue = val
				}
			}

			round = append(round, game)
		}
		result = append(result, round)

		// fmt.Println(values)
	}

	return result, nil
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0
	for id, games := range input {
		skip := false
		for _, game := range games {
			skip = skip || (game.red > 12 || game.green > 13 || game.blue > 14)
		}
		if skip {
			continue
		}
		sum += id + 1
	}

	return sum, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, games := range input {
		maxBag := Game{}
		for _, game := range games {
			maxBag.red = max(maxBag.red, game.red)
			maxBag.green = max(maxBag.green, game.green)
			maxBag.blue = max(maxBag.blue, game.blue)
		}

		sum += maxBag.red * maxBag.green * maxBag.blue
	}

	return sum, nil
}
