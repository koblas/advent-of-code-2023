package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Dirs [2]int
type Set map[int]bool

type Board struct {
	Order []int
	Steps map[int]Dirs

	StartOne  []int
	StartTwo  []int
	FinishOne Set
	FinishTwo Set
}

func ParseInput(lines []string, jokers bool) (Board, error) {
	result := Board{
		Steps:     map[int]Dirs{},
		FinishOne: Set{},
		FinishTwo: Set{},
		StartOne:  []int{},
		StartTwo:  []int{},
	}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	for _, ch := range lines[0] {
		value := 0
		if ch == 'R' {
			value = 1
		}
		result.Order = append(result.Order, value)
	}

	steps := map[string][]string{}
	indexes := map[string]int{}

	for idx, line := range lines[1:] {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		// d, _ := json.Marshal(parts)
		// fmt.Println(string(d))
		if len(parts) != 4 {
			return result, fmt.Errorf("Bad line %s", line)
		}

		step := parts[0]

		steps[step] = []string{
			parts[2][1 : len(parts[2])-1],
			parts[3][0 : len(parts[3])-1],
		}
		indexes[step] = idx

		if step == "ZZZ" {
			result.FinishOne[idx] = true
		} else if step == "AAA" {
			result.StartOne = []int{idx}
		}
		if step[2] == 'Z' {
			result.FinishTwo[idx] = true
		} else if step[2] == 'A' {
			result.StartTwo = append(result.StartTwo, idx)
		}
	}

	for key, values := range steps {
		result.Steps[indexes[key]] = Dirs{
			indexes[values[0]],
			indexes[values[1]],
		}
	}

	return result, nil
}

// smart bruteforcing Least Common Multiple
func lcm_smartbf(a []int) int {
	n := 1
	for i := range a {
		n *= a[i]
	}
	mx := slices.Max(a)

	i := 0
	l := len(a)
	for i = 1; i <= n; i++ {
		div := 0
		for j := range a {
			if (mx*i)%a[j] == 0 {
				div++
			}
		}
		if div == l {
			return mx * i
		}
	}

	return -0
}

func runSolution(input Board, start []int, finish Set) (int, error) {
	results := []int{}
	for _, pos := range start {
		count := 0
		for idx := 0; !finish[pos]; idx = (idx + 1) % len(input.Order) {
			dir := input.Order[idx]
			pos = input.Steps[pos][dir]
			count += 1
		}

		results = append(results, count)
	}

	value := lcm_smartbf(results)

	return value, nil
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	return runSolution(input, input.StartOne, input.FinishOne)
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	return runSolution(input, input.StartTwo, input.FinishTwo)
}
