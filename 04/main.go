package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type board struct {
	numbers   [5][5]int
	marked    [5][5]bool
	markedNum int

	wonCol bool
	wonRow bool
}

func (b *board) fromText(lines []string) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			b.marked[row][col] = false
			numStr := strings.Trim(lines[row][col*2+col:(col+1)*2+col], " ")
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			b.numbers[row][col] = int(num)
		}
	}
}

func (b *board) mark(num int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b.numbers[row][col] == num {
				b.marked[row][col] = true
			}
		}
	}
}

func (b *board) hasWon() bool {
	for row := 0; row < 5; row++ {
		rowWins := true
		colWins := true

		for col := 0; col < 5; col++ {
			rowWins = b.marked[row][col] && rowWins
			colWins = b.marked[col][row] && colWins
			if !rowWins && !colWins {
				continue
			}
		}

		if rowWins {
			b.wonRow = true
			return true
		}

		if colWins {
			b.wonCol = true
			return true
		}
	}

	return false
}

func (b *board) sumUnmarked() int {
	v := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b.marked[row][col] {
				v += b.numbers[row][col]
			}
		}
	}

	return v
}

func part1(lines []string) int {
	draw := strings.Split(lines[0], ",")
	lines = lines[1:] // skip draw
	boards := make([]board, len(lines)/6)

	for boardIdx := 0; boardIdx < len(boards); boardIdx++ {
		boardLines := lines[1:6]
		lines = lines[6:]
		boards[boardIdx].fromText(boardLines)
	}

	for drawIndex := 0; ; drawIndex += 1 {
		if drawIndex >= len(draw) {
			fmt.Println("Nothing found", drawIndex)
			break
		}

		num64, _ := strconv.ParseInt(draw[drawIndex], 10, 64)
		num := int(num64)

		for boardIndex, _ := range boards {
			boards[boardIndex].mark(num)

			if boards[boardIndex].hasWon() {
				fmt.Println(boardIndex, "has won")
				if boards[boardIndex].wonRow {
					v := boards[boardIndex].sumUnmarked()
					return v * num
				} else if boards[boardIndex].wonCol {
					v := boards[boardIndex].sumUnmarked()
					return v * num
				} else {
					panic("VERY BAD!")
				}
			}
		}
	}

	return -1
}

func part2(lines []string) int {
	draw := strings.Split(lines[0], ",")
	lines = lines[1:] // skip draw
	boards := make([]board, len(lines)/6)

	for boardIdx := 0; boardIdx < len(boards); boardIdx++ {
		boardLines := lines[1:6]
		lines = lines[6:]

		boards[boardIdx].fromText(boardLines)
	}

	wins := make([]bool, len(boards))

	for drawIndex := 0; ; drawIndex += 1 {
		if drawIndex >= len(draw) {
			fmt.Println("Nothing found", drawIndex)
			break
		}

		num64, _ := strconv.ParseInt(draw[drawIndex], 10, 64)
		num := int(num64)

		for boardIndex, _ := range boards {
			boards[boardIndex].mark(num)

			if boards[boardIndex].hasWon() {
				wins[boardIndex] = true

				nowins := 0
				for _, win := range wins {
					if !win {
						nowins += 1
					}
				}

				if nowins == 0 {
					fmt.Println(boardIndex, "has won")
					if boards[boardIndex].wonRow {
						v := boards[boardIndex].sumUnmarked()
						return v * num
					} else if boards[boardIndex].wonCol {
						v := boards[boardIndex].sumUnmarked()
						return v * num
					} else {
						panic("VERY BAD!")
					}
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
