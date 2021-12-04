package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	acc := 0
	lastDepth, _ := strconv.Atoi(lines[0])

	for _, line := range lines[1:] {
		depth, _ := strconv.Atoi(line)
		if depth > lastDepth {
			acc += 1
		}

		lastDepth = depth
	}

	return acc
}

func part2(lines []string) int {
	lwin := 0
	rwin := 0

	if len(lines) < 4 {
		return 0
	}

	nums := make([]int, len(lines))
	for i, line := range lines {
		v, _ := strconv.Atoi(line)
		nums[i] = v
	}

	acc := 0

	for i := 0; i < len(nums)-3; i++ {
		lwin = nums[i] + nums[i+1] + nums[i+2]
		rwin = nums[i+1] + nums[i+2] + nums[i+3]

		if rwin > lwin {
			acc += 1
		}
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
