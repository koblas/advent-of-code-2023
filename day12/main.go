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

type Row struct {
	pattern string
	counts  []int
}

type Board struct {
	rows []Row
	row2 []Row
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			return result, fmt.Errorf("invalid part count %s", line)
		}

		var counts []int
		for _, str := range strings.Split(parts[1], ",") {
			v, err := strconv.Atoi(str)
			if err != nil {
				return result, err
			}
			counts = append(counts, v)

		}

		countsTwo := []int{}
		valueTwo := strings.Join([]string{
			parts[0],
			parts[0],
			parts[0],
			parts[0],
			parts[0],
		}, "?")
		for i := 0; i < 5; i += 1 {
			countsTwo = append(countsTwo, counts...)
		}

		result.rows = append(result.rows, Row{
			counts:  counts,
			pattern: parts[0],
		})
		result.row2 = append(result.row2, Row{
			counts:  countsTwo,
			pattern: valueTwo,
		})
	}
	return result, nil
}

type xyz [3]int

func (r Row) eval() int {
	dp := map[xyz]int{}
	return recurse(dp, (r.pattern), 0, r.counts, 0, 0)
}

func recurse(dp map[xyz]int, pattern string, pidx int, numbers []int, nidx int, grouplen int) int {
	if len(pattern) == pidx {
		if (nidx == len(numbers)-1 && numbers[nidx] == grouplen) || (nidx == len(numbers) && grouplen == 0) {
			return 1
		}
		return 0
	}

	pos := xyz{pidx, nidx, grouplen}
	if dp[pos] != 0 {
		return dp[pos] - 1
	}
	sum := 0
	char := pattern[pidx]

	if char == '?' || char == '#' {
		// place a '#' and increment the grouplen
		sum += recurse(dp, pattern, pidx+1, numbers, nidx, grouplen+1)
	}
	if char == '?' || char == '.' {
		// if grouplen > 0, we can place a '.' and close the group if the count matches
		if grouplen > 0 && nidx < len(numbers) && numbers[nidx] == grouplen {
			sum += recurse(dp, pattern, pidx+1, numbers, nidx+1, 0)
		}
		// if no group, place a '.' and simply move ahead without any matching
		if grouplen == 0 {
			sum += recurse(dp, pattern, pidx+1, numbers, nidx, 0)
		}
	}

	dp[pos] = sum + 1
	return sum
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, row := range input.rows {
		sum += row.eval()
	}

	return sum, nil
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, row := range input.row2 {
		sum += row.eval()
	}

	return sum, nil
}
