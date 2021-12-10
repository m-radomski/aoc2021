package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func part1(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, cols)
	}

	between := func(v, min, max int) bool {
		return v >= min && v <= max
	}

	var recurse func(x, y, xdir, ydir int) int
	recurse = func(x, y, xdir, ydir int) int {
		if x >= cols || y >= rows {
			return 0
		} else {
			v := 0
			lower := true
			if between(x+1, 0, cols-1) {
				lower = lower && board[y][x] < board[y][x+1]
			}
			if between(x-1, 0, cols-1) {
				lower = lower && board[y][x] < board[y][x-1]
			}
			if between(y+1, 0, rows-1) {
				lower = lower && board[y][x] < board[y+1][x]
			}
			if between(y-1, 0, rows-1) {
				lower = lower && board[y][x] < board[y-1][x]
			}

			if lower {
				v += board[y][x] + 1
			}

			if xdir == 1 && ydir == 1 {
				v += recurse(x+1, y+0, 1, 0)
				v += recurse(x+1, y+1, 1, 1)
				v += recurse(x+0, y+1, 0, 1)
			} else {
				v += recurse(x+xdir, y+ydir, xdir, ydir)
			}

			return v
		}
	}

	for row, line := range lines {
		for col, c := range line {
			board[row][col] = (int)(c - '0')
		}
	}

	acc := 0

	acc += recurse(1, 0, 1, 0)
	acc += recurse(1, 1, 1, 1)
	acc += recurse(0, 1, 0, 1)

	return acc
}

func part2(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, cols)
	}

	between := func(v, min, max int) bool {
		return v >= min && v <= max
	}

	var bsizes []int

	var flood_fill func(x, y int, seen [][]bool) int
	flood_fill = func(x, y int, seen [][]bool) int {
		if x >= cols || y >= rows || y < 0 || x < 0 {
			return 0
		} else {
			if board[y][x] == 9 || seen[y][x] {
				return 0
			} else {
				seen[y][x] = true
				xv := flood_fill(x+1, y, seen) + flood_fill(x-1, y, seen)
				yv := flood_fill(x, y+1, seen) + flood_fill(x, y-1, seen)

				return 1 + xv + yv
			}
		}
	}

	var recurse func(x, y, xdir, ydir int)
	recurse = func(x, y, xdir, ydir int) {
		if x >= cols || y >= rows {
			return
		} else {
			lower := true
			if between(x+1, 0, cols-1) {
				lower = lower && board[y][x] < board[y][x+1]
			}
			if between(x-1, 0, cols-1) {
				lower = lower && board[y][x] < board[y][x-1]
			}
			if between(y+1, 0, rows-1) {
				lower = lower && board[y][x] < board[y+1][x]
			}
			if between(y-1, 0, rows-1) {
				lower = lower && board[y][x] < board[y-1][x]
			}

			if lower {
				// start search for basin
				seen := make([][]bool, rows)
				for i := 0; i < rows; i++ {
					seen[i] = make([]bool, cols)
				}
				bsize := flood_fill(x, y, seen)
				bsizes = append(bsizes, bsize)
			}

			if xdir == 1 && ydir == 1 {
				recurse(x+1, y+0, 1, 0)
				recurse(x+1, y+1, 1, 1)
				recurse(x+0, y+1, 0, 1)
			} else {
				recurse(x+xdir, y+ydir, xdir, ydir)
			}
		}
	}

	for row, line := range lines {
		for col, c := range line {
			board[row][col] = (int)(c - '0')
		}
	}

	recurse(1, 0, 1, 0)
	recurse(1, 1, 1, 1)
	recurse(0, 1, 0, 1)

	sort.SliceStable(bsizes, func(i, j int) bool { return bsizes[i] > bsizes[j] })

	return bsizes[0] * bsizes[1] * bsizes[2]
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
