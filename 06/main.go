package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var times [9]int

	numStrs := strings.Split(lines[0], ",")
	for _, numStr := range numStrs {
		num64, _ := strconv.ParseInt(numStr, 10, 64)
		num := int(num64)

		times[num] += 1
	}

	for i := 0; i < 80; i++ {
		var newTimes [9]int
		for cycle := 1; cycle <= 8; cycle++ {
			newTimes[cycle-1] = times[cycle]
		}
		newTimes[8] += times[0]
		newTimes[6] += times[0]
		times = newTimes
	}

	acc := 0
	for _, v := range times {
		acc += v
	}

	return acc
}

func part2(lines []string) int {
	var times [9]int

	numStrs := strings.Split(lines[0], ",")
	for _, numStr := range numStrs {
		num64, _ := strconv.ParseInt(numStr, 10, 64)
		num := int(num64)

		times[num] += 1
	}

	for i := 0; i < 256; i++ {
		var newTimes [9]int
		for cycle := 1; cycle <= 8; cycle++ {
			newTimes[cycle-1] = times[cycle]
		}
		newTimes[8] += times[0]
		newTimes[6] += times[0]
		times = newTimes
	}

	acc := 0
	for _, v := range times {
		acc += v
	}

	return acc
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
