package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, cols)
	}

	for row, line := range lines {
		for col, c := range line {
			board[row][col] = int(c - '0')
		}
	}

	steps := 100

	var recur_flash func(row, col int)

	totflash := 0
	recur_flash = func(row, col int) {
		if board[row][col] >= 10 {
			totflash += 1
			board[row][col] = 0

			for offrow := -1; offrow <= 1; offrow++ {
				for offcol := -1; offcol <= 1; offcol++ {
					if offrow == 0 && offcol == 0 {
						continue
					}

					mrow := row + offrow
					mcol := col + offcol

					if mrow >= 0 && mrow < rows && mcol >= 0 && mcol < cols && board[mrow][mcol] != 0 {
						board[mrow][mcol] += 1
						recur_flash(mrow, mcol)
					}
				}
			}
		}
	}

	for step := 0; step < steps; step++ {
		// Increment all board values by one
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				board[row][col] += 1
			}
		}

		// Flash
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				recur_flash(row, col)
			}
		}
	}

	return totflash
}

func part2(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, cols)
	}

	for row, line := range lines {
		for col, c := range line {
			board[row][col] = int(c - '0')
		}
	}

	var recur_flash func(row, col int) int

	recur_flash = func(row, col int) int {
		flashcnt := 0

		if board[row][col] >= 10 {
			flashcnt += 1
			board[row][col] = 0

			for offrow := -1; offrow <= 1; offrow++ {
				for offcol := -1; offcol <= 1; offcol++ {
					if offrow == 0 && offcol == 0 {
						continue
					}

					mrow := row + offrow
					mcol := col + offcol

					if mrow >= 0 && mrow < rows && mcol >= 0 && mcol < cols && board[mrow][mcol] != 0 {
						board[mrow][mcol] += 1
						flashcnt += recur_flash(mrow, mcol)
					}
				}
			}
		}

		return flashcnt
	}

	for step := 0; ; step++ {
		// Increment all board values by one
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				board[row][col] += 1
			}
		}

		// Flash
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				flash := recur_flash(row, col)
				if flash == rows*cols {
					return step + 1
				}
			}
		}
	}

	return -1
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
