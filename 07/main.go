package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	numStrArray := strings.Split(lines[0], ",")
	nums := make([]int, len(numStrArray))
	maxNum := 0
	minNum := 0xffffffff

	for i, numStr := range numStrArray {
		num64, _ := strconv.ParseInt(numStr, 10, 64)
		nums[i] = int(num64)

		if nums[i] > maxNum {
			maxNum = nums[i]
		}

		if nums[i] < minNum {
			minNum = nums[i]
		}
	}

	minScore := 0xffffffff

	for i := minNum; i <= maxNum; i++ {
		score := 0
		for _, num := range nums {
			v := num - i
			if v < 0 {
				v = -v
			}

			score += v
		}

		if score < minScore {
			minScore = score
		}
	}

	return minScore
}

func part2(lines []string) int {
	numStrArray := strings.Split(lines[0], ",")
	nums := make([]int, len(numStrArray))
	maxNum := 0
	minNum := 0xffffffff

	for i, numStr := range numStrArray {
		num64, _ := strconv.ParseInt(numStr, 10, 64)
		nums[i] = int(num64)

		if nums[i] > maxNum {
			maxNum = nums[i]
		}

		if nums[i] < minNum {
			minNum = nums[i]
		}
	}

	minScore := 0xffffffff

	for i := minNum; i <= maxNum; i++ {
		score := 0
		for _, num := range nums {
			v := num - i
			if v < 0 {
				v = -v
			}

			score += (v*v + v)/2
		}

		if score < minScore {
			minScore = score
		}
	}

	return minScore
}

func main() {
	fn := "input.txt"
	// fn := "test.txt"
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
