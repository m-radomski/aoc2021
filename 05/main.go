package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	xmax := int64(0)
	ymax := int64(0)

	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	spxs := make([]int64, len(lines))
	spys := make([]int64, len(lines))
	epxs := make([]int64, len(lines))
	epys := make([]int64, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		sp := parts[0]
		ep := parts[1]

		sp_parts := strings.Split(sp, ",")
		ep_parts := strings.Split(ep, ",")

		spx, _ := strconv.ParseInt(sp_parts[0], 10, 64)
		spy, _ := strconv.ParseInt(sp_parts[1], 10, 64)

		epx, _ := strconv.ParseInt(ep_parts[0], 10, 64)
		epy, _ := strconv.ParseInt(ep_parts[1], 10, 64)

		spxs[i] = spx
		spys[i] = spy
		epxs[i] = epx
		epys[i] = epy

		xmax = max(spx, xmax)
		xmax = max(epx, xmax)

		ymax = max(spy, ymax)
		ymax = max(epy, ymax)
	}

	board := make([][]int64, xmax+1)
	for i := int64(0); i <= xmax; i++ {
		board[i] = make([]int64, ymax+1)
	}

	for i := 0; i < len(lines); i++ {
		spx := spxs[i]
		spy := spys[i]
		epx := epxs[i]
		epy := epys[i]

		if spx == epx {
			startx := min(spy, epy)
			endx := max(spy, epy)

			for col := startx; col <= endx; col++ {
				board[spx][col] += 1
			}
		} else if spy == epy {
			starty := min(spx, epx)
			endy := max(spx, epx)

			for row := starty; row <= endy; row++ {
				board[row][spy] += 1
			}
		} else {
			// skip this one
		}
	}

	v := 0
	for row := int64(0); row <= xmax; row++ {
		for col := int64(0); col <= ymax; col++ {
			if board[row][col] >= 2 {
				//fmt.Println(row, col)
				v += 1
			}
		}
	}

	return v
}

func part2(lines []string) int {
	xmax := int64(0)
	ymax := int64(0)

	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	spxs := make([]int64, len(lines))
	spys := make([]int64, len(lines))
	epxs := make([]int64, len(lines))
	epys := make([]int64, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		sp := parts[0]
		ep := parts[1]

		sp_parts := strings.Split(sp, ",")
		ep_parts := strings.Split(ep, ",")

		spx, _ := strconv.ParseInt(sp_parts[0], 10, 64)
		spy, _ := strconv.ParseInt(sp_parts[1], 10, 64)

		epx, _ := strconv.ParseInt(ep_parts[0], 10, 64)
		epy, _ := strconv.ParseInt(ep_parts[1], 10, 64)

		spxs[i] = spx
		spys[i] = spy
		epxs[i] = epx
		epys[i] = epy

		xmax = max(spx, xmax)
		xmax = max(epx, xmax)

		ymax = max(spy, ymax)
		ymax = max(epy, ymax)
	}

	board := make([][]int64, xmax+1)
	for i := int64(0); i <= xmax; i++ {
		board[i] = make([]int64, ymax+1)
	}

	for i := 0; i < len(lines); i++ {
		spx := spxs[i]
		spy := spys[i]
		epx := epxs[i]
		epy := epys[i]

		if spx == epx {
			startx := min(spy, epy)
			endx := max(spy, epy)

			for col := startx; col <= endx; col++ {
				board[spx][col] += 1
			}
		} else if spy == epy {
			starty := min(spx, epx)
			endy := max(spx, epx)

			for row := starty; row <= endy; row++ {
				board[row][spy] += 1
			}
		} else {
			xdir := int64(1)
			if spx >= epx {
				xdir = -xdir
			}

			ydir := int64(1)
			if spy >= epy {
				ydir = -ydir
			}

			test := func(v, vmax, dir int64) bool {
				if dir > 0 {
					return v <= vmax
				} else {
					return v >= vmax
				}
			}

			for row, col := spx, spy; test(row, epx, xdir); row, col = row+xdir, col+ydir {
				board[row][col] += 1
			}
		}
	}

	v := 0
	for row := int64(0); row <= xmax; row++ {
		for col := int64(0); col <= ymax; col++ {
			if board[row][col] >= 2 {
				v += 1
			}
		}
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
