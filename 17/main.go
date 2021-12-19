package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func between(min, v, max int) bool {
	return min <= v && max >= v
}

func sim(vx, vy, xbmin, xbmax, ybmin, ybmax int) (bool, int) {
	x, y := 0, 0
	ymax := 0

	for true {
		x += vx
		y += vy

		if y > ymax {
			ymax = y
		}

		if vx > 0 {
			vx -= 1
		} else if vx < 0 {
			vx += 1
		}

		vy -= 1

		if between(xbmin, x, xbmax) && between(ybmin, y, ybmax) {
			return true, ymax
		} else if vy < 0 && y < ybmin {
			return false, ymax
		}
	}

	return false, ymax
}

func part1(lines []string) int {
	coords := strings.Trim(strings.Split(lines[0], ": ")[1], " ")
	parts := strings.Split(coords, ", ")

	xpart := parts[0][2:]
	ypart := parts[1][2:]

	xbounds := strings.Split(xpart, "..")
	ybounds := strings.Split(ypart, "..")

	xbmin, _ := strconv.Atoi(xbounds[0])
	xbmax, _ := strconv.Atoi(xbounds[1])
	ybmin, _ := strconv.Atoi(ybounds[0])
	ybmax, _ := strconv.Atoi(ybounds[1])

	yoptmax := 0
	for vx := -1000; vx < 1000; vx++ {
		for vy := -1000; vy < 1000; vy++ {
			good, ymax := sim(vx, vy, xbmin, xbmax, ybmin, ybmax)
			if good && ymax > yoptmax {
				yoptmax = ymax
			}
		}
	}

	return yoptmax
}

func part2(lines []string) int {
	coords := strings.Trim(strings.Split(lines[0], ": ")[1], " ")
	parts := strings.Split(coords, ", ")

	xpart := parts[0][2:]
	ypart := parts[1][2:]

	xbounds := strings.Split(xpart, "..")
	ybounds := strings.Split(ypart, "..")

	xbmin, _ := strconv.Atoi(xbounds[0])
	xbmax, _ := strconv.Atoi(xbounds[1])
	ybmin, _ := strconv.Atoi(ybounds[0])
	ybmax, _ := strconv.Atoi(ybounds[1])

	acc := 0
	for vx := -1000; vx < 1000; vx++ {
		for vy := -1000; vy < 1000; vy++ {
			good, _ := sim(vx, vy, xbmin, xbmax, ybmin, ybmax)
			if good {
				acc += 1
			}
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
