package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	horz := 0
	depth := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		cmd := parts[0]
		val, _ := strconv.Atoi(parts[1])

		if cmd == "forward" {
			horz += val
		} else if cmd == "down" {
			depth += val
		} else if cmd == "up" {
			depth -= val
		}
	}

	return horz * depth
}

func part2(lines []string) int {
	horz := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		cmd := parts[0]
		val, _ := strconv.Atoi(parts[1])

		if cmd == "forward" {
			horz += val
			depth += val * aim
		} else if cmd == "down" {
			aim += val
		} else if cmd == "up" {
			aim -= val
		}
	}

	return horz * depth
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
