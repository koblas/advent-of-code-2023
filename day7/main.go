package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type HandType int

const (
	Invalid HandType = iota
	Five    HandType = iota
	Four    HandType = iota
	Full    HandType = iota
	Three   HandType = iota
	Two     HandType = iota
	One     HandType = iota
	High    HandType = iota
)

type Hand struct {
	Orig   []int
	Sorted []int
	Cards  string
	Bid    int
	Rank   HandType
}

var rank = map[rune]int{}

func init() {
	cards := "AKQJT98765432"
	for idx, ch := range []rune("AKQJT98765432") {
		rank[ch] = len(cards) - idx + 1
	}
}

func classify(value string, jokers bool) (HandType, []int, []int) {
	counts := map[int]int{}

	cards := []int{}
	orig := []int{}
	jcount := 0
	for _, ch := range value {
		var r int
		if jokers && ch == 'J' {
			r = 0
			jcount += 1
		} else {
			r = rank[ch]
			cards = append(cards, r)
			counts[r] += 1
		}
		orig = append(orig, r)
		// orig = append(orig, rank[ch])
	}

	sort.Slice(cards, func(i, j int) bool {
		ca := cards[i]
		cb := cards[j]

		diff := counts[cb] - counts[ca]
		if diff == 0 {
			return ca > cb
		}
		return diff < 0
	})

	// fmt.Println(value, cards)
	if jcount != 0 {
		best := 0
		if len(cards) != 0 {
			best = cards[0]
		}
		for len(cards) != 5 {
			cards = append([]int{best}, cards...)
		}
	}

	rank := High
	switch len(counts) {
	case 5:
		rank = High
	case 4:
		rank = One
	case 3:
		if cards[0] == cards[2] {
			rank = Three
		} else {
			rank = Two
		}
	case 2:
		if cards[0] == cards[3] {
			rank = Four
		} else {
			rank = Full
		}
	case 1, 0:
		rank = Five
	default:
		panic("SHOULDN't HAPPEN")
	}
	// fmt.Println("INPUT", value, orig, cards, counts, rank)
	return rank, orig, cards
}

func ParseInput(lines []string, jokers bool) ([]Hand, error) {
	var result []Hand

	if len(lines) == 0 {
		return result, fmt.Errorf("bad input")
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Bad parts")
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		rank, orig, _ := classify(parts[0], jokers)
		result = append(result, Hand{
			// Sorted: sorted,
			Orig:  orig,
			Cards: parts[0],
			Bid:   value,
			Rank:  rank,
		})
	}

	return result, nil
}

type ByHand []Hand

func (a ByHand) Len() int      { return len(a) }
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool {
	if a[i].Rank != a[j].Rank {
		return a[i].Rank < a[j].Rank
	}
	for idx, value := range a[i].Orig {
		jVal := a[j].Orig[idx]
		if jVal != value {
			return jVal < value
		}
		// if a[j].Sorted[idx] != value {
		// 	return a[j].Sorted[idx] < value
		// }
	}
	panic("NOT HERE")
	return true
}

func PartOneSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, false)
	if err != nil {
		return 0, err
	}

	sort.Sort(ByHand(input))

	sum := 0
	for idx, hand := range input {
		rank := (len(input) - idx)
		score := rank * hand.Bid

		// fmt.Printf("%d: hand=%v  cards=%s   bid=%d  rank=%d score=%d\n", idx, hand.Orig, hand.Cards, hand.Bid, rank, score)
		sum += score
	}

	return sum, nil
}

func PartTwoSolution(lines []string) (int, error) {
	input, err := ParseInput(lines, true)
	if err != nil {
		return 0, err
	}

	sort.Sort(ByHand(input))

	sum := 0
	for idx, hand := range input {
		rank := (len(input) - idx)
		score := rank * hand.Bid

		// fmt.Printf("%3d: hand=%2v  cards=%s   bid=%d  type=%d  rank=%d score=%d\n", idx, hand.Orig, hand.Cards, hand.Bid, hand.Rank, rank, score)
		sum += score
	}

	return sum, nil
}
