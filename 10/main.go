package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func part1(lines []string) int {
	var points [4]int

	openers := []rune{'(', '[', '{', '<'}
	scores := []int{3, 57, 1197, 25137}
	closer_to_bindex := map[rune]int{')': 0, ']': 1, '}': 2, '>': 3}

	for _, line := range lines {
		stop := 0
		stack := make([]rune, len(line))
		for _, c := range line {
			if c == '{' || c == '(' || c == '[' || c == '<' {
				stack[stop] = c
				stop += 1
			} else {
				bindex := closer_to_bindex[c]
				popped := stack[stop-1]

				if popped == openers[bindex] {
					stop -= 1
				} else {
					points[bindex] += 1
					break
				}
			}
		}
	}

	v := 0
	for i, score := range scores {
		v += score * points[i]
	}

	return v
}

func part2(lines []string) int {
	var points [4]int

	openers := []rune{'(', '[', '{', '<'}
	scores := make([]int, 0)
	opener_to_bindex := map[rune]int{'(': 0, '[': 1, '{': 2, '<': 3}
	closer_to_bindex := map[rune]int{')': 0, ']': 1, '}': 2, '>': 3}

	for _, line := range lines {
		stop := 0
		stack := make([]rune, len(line))
		invalid := false
		for _, c := range line {
			if c == '{' || c == '(' || c == '[' || c == '<' {
				stack[stop] = c
				stop += 1
			} else {
				bindex := closer_to_bindex[c]
				popped := stack[stop-1]

				if popped == openers[bindex] {
					stop -= 1
				} else {
					points[bindex] += 1
					invalid = true
					break
				}
			}
		}

		if !invalid {
			score := 0
			for cidx := stop - 1; cidx >= 0; cidx-- {
				opener := stack[cidx]
				bindex := opener_to_bindex[opener]

				score *= 5
				score += bindex + 1
			}

			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	fn := "input.txt"
	//fn := "test.txt"
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(buf), "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	acc := part1(lines)
	fmt.Println(acc)

	acc2 := part2(lines)
	fmt.Println(acc2)
}
