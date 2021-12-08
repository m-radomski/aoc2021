package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(lines []string) int {
	v := 0

	for _, line := range lines {
		parts := strings.Split(line, "|")
		readout := strings.Trim(parts[1], " ")

		for _, valEncStr := range strings.Split(readout, " ") {
			l := len(valEncStr)

			if l == 2 || l == 3 || l == 4 || l == 7 {
				v += 1
			}
		}
	}

	return v
}

func part2(lines []string) int {
	v := 0

	encAsInt := func(s string) int {
		v := 0
		for _, c := range s {
			v |= (1 << int(c-'a'))
		}

		return v
	}

	for _, line := range lines {
		parts := strings.Split(line, "|")

		// parts := int[7]
		var dict [10]int

		var unmatched6 [3]int
		var unmatched5 [3]int
		k6 := 0
		k5 := 0

		readin := strings.Trim(parts[0], " ")
		for _, valEncStr := range strings.Split(readin, " ") {
			l := len(valEncStr)

			num := 0
			if l == 2 {
				num = 1
			} else if l == 3 {
				num = 7
			} else if l == 4 {
				num = 4
			} else if l == 7 {
				num = 8
			} else {
				if len(valEncStr) == 6 {
					unmatched6[k6] = encAsInt(valEncStr)
					k6 += 1
				} else {
					unmatched5[k5] = encAsInt(valEncStr)
					k5 += 1
				}
			}

			if num > 0 {
				dict[num] = encAsInt(valEncStr)
			}
		}

		for i, v := range unmatched6 {
			if v&dict[4] == dict[4] { // Match 9
				dict[9] = v
				unmatched6[i] = 0
			} else if v&dict[7] == dict[7] { // Match 0
				dict[0] = v
				unmatched6[i] = 0
			}
		}

		// Last is 6
		for _, v := range unmatched6 {
			if v != 0 {
				dict[6] = v
				break
			}
		}

		for i, v := range unmatched5 {
			if v&dict[1] == dict[1] { // Match 3
				dict[3] = v
				unmatched5[i] = 0
			} else if v&((^dict[4])&0x7f) == (^dict[4])&0x7f { // Match 2
				dict[2] = v
				unmatched5[i] = 0
			}
		}

		// Last is 5
		for _, v := range unmatched5 {
			if v != 0 {
				dict[5] = v
				break
			}
		}

		output := 0
		readout := strings.Trim(parts[1], " ")
		for _, valEncStr := range strings.Split(readout, " ") {
			v := encAsInt(valEncStr)

			for num, dval := range dict {
				if v == dval {
					output = (output * 10) + num
				}
			}
		}

		v += output
	}

	return v

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
