package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func part1(lines []string) int {
	onesCount := make([]int, len(lines[0]))
	zerosCount := make([]int, len(lines[0]))

	for _, line := range lines {
		for ic, c := range line {
			if c == '1' {
				onesCount[ic] += 1
			} else if c == '0' {
				zerosCount[ic] += 1
			}
		}
	}

	gamma_r := 0
	for i := 0; i < len(onesCount); i++ {
		if onesCount[i] > zerosCount[i] {
			gamma_r = (gamma_r << 1) | 1
		} else {
			gamma_r = (gamma_r << 1) | 0
		}
	}

	epsilon_r := ^gamma_r & ((1 << len(onesCount)) - 1)

	return gamma_r * epsilon_r
}

func countBits(lines []string) ([]int, []int) {
	onesCount := make([]int, len(lines[0]))
	zerosCount := make([]int, len(lines[0]))

	for _, line := range lines {
		for ic, c := range line {
			if c == '1' {
				onesCount[ic] += 1
			} else if c == '0' {
				zerosCount[ic] += 1
			}
		}
	}

	return onesCount, zerosCount
}

func part2(lines []string) int {

	filtered := lines
	for i := 0; true; i++ {
		onesCount, zerosCount := countBits(filtered)

		if onesCount[i] >= zerosCount[i] {
			test := func(s string) bool { return s[i] == '1' }
			filtered = filter(filtered, test)
		} else {
			test := func(s string) bool { return s[i] == '0' }
			filtered = filter(filtered, test)
		}

		if len(filtered) == 1 {
			break
		}
	}

	oxygen, _ := strconv.ParseInt(filtered[0], 2, 32)

	filtered = lines
	for i := 0; true; i++ {
		onesCount, zerosCount := countBits(filtered)

		if onesCount[i] < zerosCount[i] {
			test := func(s string) bool { return s[i] == '1' }
			filtered = filter(filtered, test)
		} else {
			test := func(s string) bool { return s[i] == '0' }
			filtered = filter(filtered, test)
		}

		if len(filtered) == 1 {
			break
		}
	}

	co2, _ := strconv.ParseInt(filtered[0], 2, 32)

	return int(oxygen * co2)
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
