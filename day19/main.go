package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Rule struct {
	key   string
	op    rune
	value int
	next  string
}

type Part struct {
	values map[string]int
	sum    int
}

type Input struct {
	workflows map[string][]Rule
	parts     []Part
}

func runWorkflow(input Input, part Part) int {
	current := "in"

	for {
		// fmt.Println("RUNNING", current)
		if current == "A" {
			return part.sum
		} else if current == "R" {
			return 0
		}

		flow, found := input.workflows[current]
		if !found {
			panic("Rule not found " + current)
		}

		for _, step := range flow {
			// fmt.Println(step)

			if step.op == '<' {
				val := part.values[step.key]
				if val < step.value {
					current = step.next
					break
				}
			} else if step.op == '>' {
				val := part.values[step.key]
				if val > step.value {
					current = step.next
					break
				}
			} else {
				current = step.next
				break
			}
		}
	}
}

func solve(input Input) int {
	sum := 0
	for _, part := range input.parts {
		sum += runWorkflow(input, part)
	}
	return sum
}

type Set struct {
	min, max int
}
type SetValues map[string]Set

type Range struct {
	state  string
	values SetValues
}

func (vals SetValues) String() string {
	buf := ""

	for _, k := range []string{"x", "a", "m", "s"} {
		val := vals[k]
		buf += fmt.Sprintf("%s: [%d,%d] ", k, val.min, val.max)
	}

	return "{" + buf + "}"

}

func splitValues(key string, val int, values SetValues) [2]*SetValues {
	out1 := SetValues{}
	out2 := SetValues{}

	for k, v := range values {
		out1[k] = v
		out2[k] = v
	}

	out1[key] = Set{
		min: out1[key].min,
		max: val - 1,
	}
	out2[key] = Set{
		min: val,
		max: out2[key].max,
	}

	r1 := &out1
	r2 := &out2

	if out1[key].max < out1[key].min {
		r1 = nil
	}
	if out2[key].max < out2[key].min {
		r2 = nil
	}

	return [2]*SetValues{r1, r2}
}

func solveTwo(input Input) int {
	accepted := []SetValues{}

	start := Range{
		state: "in",
		values: SetValues{
			"x": Set{1, 4000},
			"m": Set{1, 4000},
			"a": Set{1, 4000},
			"s": Set{1, 4000},
		},
	}

	queue := []Range{start}
	for len(queue) != 0 {
		first := queue[0]
		queue = queue[1:]
		current := first.state

		currentValues := &first.values

		for currentValues != nil {
			// fmt.Println("RUNNING", current)
			if current == "A" {
				accepted = append(accepted, *currentValues)
				break
			} else if current == "R" {
				break
			}

			flow, found := input.workflows[current]
			if !found {
				panic("Rule not found " + current)
			}

			for _, step := range flow {
				if currentValues == nil {
					break
				}
				if step.op == '<' {
					vnew := splitValues(step.key, step.value, *currentValues)

					// Handle low part
					if vnew[0] != nil {
						queue = append(queue, Range{
							state:  step.next,
							values: *vnew[0],
						})
					}
					currentValues = vnew[1]
				} else if step.op == '>' {
					vnew := splitValues(step.key, step.value+1, *currentValues)

					// Handle high part
					if vnew[1] != nil {
						queue = append(queue, Range{
							state:  step.next,
							values: *vnew[1],
						})
					}
					currentValues = vnew[0]
				} else {
					queue = append(queue, Range{
						state:  step.next,
						values: *currentValues,
					})
					currentValues = nil
					break
				}
			}
		}
	}

	sum := 0
	for _, item := range accepted {
		possibilites := 1
		for _, val := range item {
			possibilites *= 1 + val.max - val.min
		}
		// fmt.Printf("Accepted %s   poss=%d\n", item.String(), possibilites)
		sum += possibilites
	}
	return sum
}

var ruleExp = regexp.MustCompile(`(\w+)([<>])(\d+):(\w+)`)

func ParseInput(lines []string) (Input, error) {
	input := Input{
		workflows: map[string][]Rule{},
	}

	if len(lines) == 0 {
		return input, fmt.Errorf("bad input")
	}

	cut := 0
	for idx, line := range lines {
		if line == "" {
			cut = idx
			break
		}
	}

	for _, line := range lines[0:cut] {
		idx := strings.Index(line, "{")
		if idx == -1 {
			return input, fmt.Errorf("separator not found")
		}
		name := line[0:idx]
		line = strings.TrimPrefix(line[idx:], "{")
		line = strings.TrimSuffix(line, "}")
		rules := []Rule{}
		for _, rule := range strings.Split(line, ",") {
			action := Rule{}
			match := ruleExp.FindStringSubmatch(rule)
			if len(match) != 0 {
				action.key = match[1]
				action.op = rune(match[2][0])
				val, err := strconv.Atoi(match[3])
				if err != nil {
					return input, fmt.Errorf("invalid value: %w", err)
				}
				action.value = val
				action.next = match[4]
			} else if rule == "A" {
				action.op = 'G'
				action.next = "A"
			} else if rule == "R" {
				action.op = 'G'
				action.next = "R"
			} else {
				action.op = 'G'
				action.next = rule
			}

			rules = append(rules, action)
		}
		input.workflows[name] = rules
	}
	for _, line := range lines[cut+1:] {
		line = strings.TrimPrefix(line, "{")
		line = strings.TrimSuffix(line, "}")

		row := map[string]int{}
		sum := 0
		for _, part := range strings.Split(line, ",") {
			p2 := strings.Split(part, "=")
			val, err := strconv.Atoi(p2[1])
			if err != nil {
				return input, err
			}
			sum += val
			row[p2[0]] = val
		}
		input.parts = append(input.parts, Part{
			values: row,
			sum:    sum,
		})
	}

	return input, nil
}

func PartOneSolution(input Input) (int, error) {
	return solve(input), nil
}

func PartTwoSolution(input Input) (int, error) {
	return solveTwo(input), nil
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
	input, err := ParseInput(lines)
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
