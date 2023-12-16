package main

import (
	"bufio"
	"fmt"
	"os"
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

type Card map[string]struct{}
type Game struct {
	winning Card
	holding Card
	all     Card
}

func ParseInput(lines []string) ([]Game, error) {
	result := []Game{}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("bad row format: %s", line)
		}

		parts = strings.Split(parts[1], "|")
		if len(parts) != 2 {
			return nil, fmt.Errorf("bad card format: %s", line)
		}
		game := Game{}
		// fmt.Printf("'%s' | '%s'\n", parts[0], parts[1])
		all := Card{}
		for idx, group := range parts {
			numbers := Card{}
			for _, num := range strings.Split(strings.TrimSpace(group), " ") {
				if num == "" {
					continue
				}
				numbers[num] = struct{}{}
				all[num] = struct{}{}
			}
			// fmt.Println(numbers)
			if idx == 0 {
				game.winning = numbers
			} else {
				game.holding = numbers
			}
		}
		game.all = all
		result = append(result, game)
	}

	return result, nil
}

func matches(game Game) int {
	count := 0
	for value := range game.winning {
		if _, found := game.holding[value]; found {
			count += 1
		}
	}
	return count
}

func score(count int) int {
	if count == 0 {
		return 0
	}

	return 1 << (count - 1)
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, game := range input {
		count := matches(game)
		value := score(count)

		// fmt.Println("Got ", idx, value)

		sum += value
	}

	return sum, nil
}

func decend(idx int, counts []int, memo map[int]int) int {
	if value, found := memo[idx]; found {
		return value
	}

	count := counts[idx]
	sum := 1
	for pos := idx + 1; pos < idx+count+1; pos += 1 {
		sum += decend(pos, counts, memo)
	}

	memo[idx] = sum
	return sum
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	counts := make([]int, len(input))
	for idx, game := range input {
		counts[idx] = matches(game)
	}

	sum := 0
	memo := map[int]int{}
	for idx := range input {
		sum += decend(idx, counts, memo)
	}

	return sum, nil
}
