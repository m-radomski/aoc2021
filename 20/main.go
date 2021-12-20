package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type v2i struct {
	row int
	col int
}

func solve(lines []string, steps int) int {
	pattern := make([]bool, 512)
	i := 0

	for _, line := range lines[0:1] {
		for _, c := range line {
			pattern[i] = (c == '#')
			i += 1
		}
	}

	// If the pattern will make the entire board lit up because
	// the zeroth pattern is enabled then handle the case differently
	// Basically create a boundry condition for that case
	grid_has_boundry_condition := pattern[0] && !pattern[511]

	img := make(map[v2i]bool)
	rows := len(lines[2:])
	cols := len(lines[2])
	for row, line := range lines[2:] {
		for col, c := range line {
			img[v2i{row, col}] = (c == '#')
		}
	}

	for k := 0; k < steps; k++ {
		nrows := rows + 2
		ncols := cols + 2
		nimg := make(map[v2i]bool)

		for row := 0; row < nrows; row++ {
			for col := 0; col < ncols; col++ {
				idx := 0

				for rm := 0; rm < 3; rm++ {
					for cm := 0; cm < 3; cm++ {
						idx <<= 1
						val := 0

						rowindex := row + rm - 2
						colindex := col + cm - 2
						on_edge := rowindex < 0 || colindex < 0 || colindex >= rows || rowindex >= cols
						if grid_has_boundry_condition && ((k % 2) == 1) && on_edge {
							val = 1
						} else if img[v2i{rowindex, colindex}] {
							val = 1
						}

						idx = idx | val
					}
				}

				nimg[v2i{row, col}] = pattern[idx]
			}
		}

		rows = nrows
		cols = ncols
		img = nimg
	}

	acc := 0
	for _, val := range img {
		if val {
			acc += 1
		}
	}

	return acc
}

func part1(lines []string) int {
	return solve(lines, 2)
}

func part2(lines []string) int {
	return solve(lines, 50)
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
