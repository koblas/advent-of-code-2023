package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

type Pos [2]int
type Grid map[Pos]rune
type Number struct {
	strval string
	value  int
	startX int
	startY int
	len    int
}

type Game struct {
	grid    Grid
	numbers []*Number
}

func ParseInput(lines []string) (Game, error) {
	result := Game{
		grid: Grid{},
	}

	for y, line := range lines {
		var number *Number

		storeNum := func() {
			if number == nil {
				return
			}

			number.len = len(number.strval)
			val, _ := strconv.Atoi(number.strval)
			number.value = val

			result.numbers = append(result.numbers, number)

			number = nil
		}

		for x, ch := range line {
			if unicode.IsDigit(ch) {
				if number == nil {
					number = &Number{
						startX: x,
						startY: y,
					}

				}
				number.strval += string(ch)
				continue
			} else {
				storeNum()
			}
			if ch == '.' {
				continue
			}
			result.grid[Pos{x, y}] = ch
		}
		storeNum()
	}

	return result, nil
}

func isAdjacent(number *Number, grid Grid) bool {
	for y := number.startY - 1; y < number.startY+2; y++ {
		for x := number.startX - 1; x < number.startX+number.len+1; x++ {
			// fmt.Println("CHECK", x, y)
			if _, found := grid[[2]int{x, y}]; found {
				// fmt.Println("FOUND", number.strval, x, y)
				return true
			}
		}
	}
	// fmt.Println("NOT", number.strval)

	return false
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, number := range input.numbers {
		if isAdjacent(number, input.grid) {
			sum += number.value
		}
	}

	return sum, nil
}

func findNearby(pos Pos, numbers []*Number) []*Number {
	result := []*Number{}
	x := pos[0]
	y := pos[1]

	for _, number := range numbers {
		minX := number.startX - 1
		maxX := number.startX + number.len + 1
		minY := number.startY - 1
		maxY := number.startY + 2

		if x < minX || x >= maxX || y < minY || y >= maxY {
			continue
		}
		result = append(result, number)
	}

	return result
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sum := 0

	for pos, value := range input.grid {
		if value != '*' {
			continue
		}

		nearby := findNearby(pos, input.numbers)

		if len(nearby) != 2 {
			// fmt.Println("POS ", pos)
			// panic(fmt.Sprintf("shouldn't happen %d", len(nearby)))
			continue
		}

		sum += nearby[0].value * nearby[1].value
	}

	return sum, nil
}
