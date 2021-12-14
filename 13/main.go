package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BoolMatrix struct {
	rows int
	cols int
	data []bool // row-major order
}

func (BM BoolMatrix) Index(row, col int) int {
	return col + row*BM.cols
}

func (BM *BoolMatrix) FromLines(lines []string, rows, cols int) {
	BM.rows = rows
	BM.cols = cols
	BM.data = make([]bool, BM.rows*BM.cols)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])

		BM.data[BM.Index(row, col)] = true
	}
}

func (BM *BoolMatrix) FoldY(row int) {
	var NBM BoolMatrix
	NBM.rows = row
	NBM.cols = BM.cols
	NBM.data = make([]bool, NBM.rows*NBM.cols)

	for ri := 0; ri < row; ri++ {
		for col := 0; col < BM.cols; col++ {
			NBM.data[NBM.Index(ri, col)] = (BM.data[BM.Index(ri, col)] || BM.data[BM.Index(BM.rows-1-ri, col)])
		}
	}

	*BM = NBM
}

func (BM *BoolMatrix) FoldX(col int) {
	var NBM BoolMatrix
	NBM.rows = BM.rows
	NBM.cols = col
	NBM.data = make([]bool, NBM.rows*NBM.cols)

	for row := 0; row < BM.rows; row++ {
		for ci := 0; ci < col; ci++ {
			NBM.data[NBM.Index(row, ci)] = BM.data[BM.Index(row, ci)] || BM.data[BM.Index(row, BM.cols-1-ci)]
		}
	}

	*BM = NBM
}

func (BM *BoolMatrix) CountDots() int {
	acc := 0

	for _, dot := range BM.data {
		if dot {
			acc += 1
		}
	}

	return acc
}

func (BM *BoolMatrix) PrintCode() {
	for row := 0; row < BM.rows; row++ {
		for col := 0; col < BM.cols; col++ {
			if BM.data[BM.Index(row, col)] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func part1(lines []string) int {
	// Find separator
	var dots []string
	var folds []string
	for i, line := range lines {
		if len(line) == 0 {
			dots = lines[:i]
			folds = lines[i+1:]
			break
		}
	}

	rows := 0
	cols := 0

	for _, fold := range folds {
		parts := strings.Split(fold, "=")
		axis := parts[0][len(parts[0])-1:]
		loc, _ := strconv.Atoi(parts[1])

		if axis == "y" && rows == 0 {
			rows = loc*2 + 1
		} else if axis == "x" && cols == 0 {
			cols = loc*2 + 1
		}

		if cols != 0 && rows != 0 {
			break
		}
	}

	var BM BoolMatrix
	BM.FromLines(dots, rows, cols)

	for _, fold := range folds {
		parts := strings.Split(fold, "=")
		axis := parts[0][len(parts[0])-1:]
		loc, _ := strconv.Atoi(parts[1])

		if axis == "y" {
			BM.FoldY(loc)
		} else {
			BM.FoldX(loc)
		}

		break
	}

	return BM.CountDots()
}

func part2(lines []string) int {
	// Find separator
	var dots []string
	var folds []string
	for i, line := range lines {
		if len(line) == 0 {
			dots = lines[:i]
			folds = lines[i+1:]
			break
		}
	}

	rows := 0
	cols := 0

	for _, fold := range folds {
		parts := strings.Split(fold, "=")
		axis := parts[0][len(parts[0])-1:]
		loc, _ := strconv.Atoi(parts[1])

		if axis == "y" && rows == 0 {
			rows = loc*2 + 1
		} else if axis == "x" && cols == 0 {
			cols = loc*2 + 1
		}

		if cols != 0 && rows != 0 {
			break
		}
	}

	var BM BoolMatrix
	BM.FromLines(dots, rows, cols)

	for _, fold := range folds {
		parts := strings.Split(fold, "=")
		axis := parts[0][len(parts[0])-1:]
		loc, _ := strconv.Atoi(parts[1])

		if axis == "y" {
			BM.FoldY(loc)
		} else {
			BM.FoldX(loc)
		}
	}

	BM.PrintCode()

	return 0
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
