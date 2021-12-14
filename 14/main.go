package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func solve(lines []string, steps int) int64 {
	template := lines[0]
	recipeArray := lines[2:]

	recipeMap := make(map[string]byte)
	for _, recipe := range recipeArray {
		parts := strings.Split(recipe, " -> ")
		recipeMap[parts[0]] = parts[1][0]
	}

	p := make(map[string]int64)
	for pp := 0; pp < len(template)-1; pp++ {
		p[template[pp:pp+2]] += 1
	}

	occur := make(map[byte]int64)
	for _, c := range template {
		occur[byte(c)] += 1
	}

	for i := 0; i < steps; i++ {
		np := make(map[string]int64)

		for key, val := range p {
			insert := recipeMap[key]

			key1 := fmt.Sprintf("%c%c", key[0], insert)
			key2 := fmt.Sprintf("%c%c", insert, key[1])

			np[key1] += val
			np[key2] += val
			occur[insert] += val
		}

		p = np
	}

	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)

	for _, val := range occur {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	return max - min
}

func part1(lines []string) int {
	return int(solve(lines, 10))
}

func part2(lines []string) int64 {
	return solve(lines, 40)
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
