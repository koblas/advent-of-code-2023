package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operation struct {
	key   string
	op    rune
	value int
}

type Input struct {
	steps []string

	part2 []Operation
}

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

	timeStart := time.Now()
	input, err := ParseInput(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Build input (%.2fms)\n", float64(time.Since(timeStart).Microseconds())/1000)

	timeStart = time.Now()
	values, err := PartOneSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)

	timeStart = time.Now()
	values, err = PartTwoSolution(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2 (%.2fms): %v\n", float64(time.Since(timeStart).Microseconds())/1000, values)
}

func ParseInput(lines []string, jokers bool) (Input, error) {
	result := Input{}

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	for _, line := range lines {
		result.steps = strings.Split(line, ",")

		for _, step := range result.steps {
			val := 0
			op := '-'
			parts := strings.Split(step, "=")
			key := parts[0]
			if len(parts) == 2 {
				op = '='
				var err error
				val, err = strconv.Atoi(parts[1])
				if err != nil {
					return result, err
				}
			} else {
				key = parts[0][0 : len(parts[0])-1]
			}
			result.part2 = append(result.part2, Operation{
				key:   key,
				op:    op,
				value: val,
			})
		}
	}

	return result, nil
}

func hash(value string) int {
	result := 0
	for _, ch := range value {
		result += int(ch)
		result *= 17
		result %= 256
	}

	return result
}

func PartOneSolution(input Input) (int, error) {
	sum := 0
	for _, step := range input.steps {
		sum += hash(step)
	}

	return sum, nil
}

func PrintBucket(idx int, bucket []Operation) {
	if len(bucket) == 0 {
		return
	}
	fmt.Printf("Box %d: ", idx)
	for _, item := range bucket {
		fmt.Printf("[%s %d] ", item.key, item.value)
	}
	fmt.Printf("\n")
}

func PartTwoSolution(input Input) (int, error) {
	buckets := [256][]Operation{}

	for _, step := range input.part2 {
		hval := hash(step.key)

		var scratch []Operation
		found := false
		for _, item := range buckets[hval] {
			if item.key == step.key {
				if step.op == '-' {
					// do nothing
				} else {
					found = true
					scratch = append(scratch, step)
				}
			} else {
				scratch = append(scratch, item)
			}
		}
		if !found && step.op == '=' {
			scratch = append(scratch, step)
		}

		// fmt.Printf("\nAfter \"%s\"\n", input.steps[idx])
		// PrintBucket(hval, buckets[hval])
		// PrintBucket(hval, scratch)
		buckets[hval] = scratch

	}
	// Print
	// fmt.Println("")
	// for idx, bucket := range buckets {
	// 	PrintBucket(idx, bucket)
	// }

	sum := 0
	for bidx, bucket := range buckets {
		for sidx, slot := range bucket {
			sum += (bidx + 1) * (sidx + 1) * slot.value
		}
	}

	// 44665 to low
	// 632055 to high

	return sum, nil
}
