package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type v2i struct {
	row int
	col int
}

func dijkstra(start, end v2i, graph [][]int) int {
	rows := len(graph)
	cols := len(graph[0])

	dist := make([][]int, rows)
	Q := make([]v2i, 0)

	for i := 0; i < cols; i++ {
		dist[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			dist[i][j] = math.MaxInt32
			Q = append(Q, v2i{i, j})
		}
	}

	dist[start.row][start.col] = 0

	for len(Q) != 0 {
		minv := math.MaxInt32
		minqi := 0
		for i, q := range Q {
			if dist[q.row][q.col] < minv {
				minv = dist[q.row][q.col]
				minqi = i
			}
		}

		minq := Q[minqi]
		Q = append(Q[:minqi], Q[minqi+1:]...)

		check := func(u, v v2i) {
			nlen := dist[u.row][u.col] + graph[v.row][v.col]
			if nlen < dist[v.row][v.col] {
				dist[v.row][v.col] = nlen
			}
		}

		if minq.row > 0 {
			check(minq, v2i{minq.row - 1, minq.col})
		}
		if minq.row < rows-1 {
			check(minq, v2i{minq.row + 1, minq.col})
		}
		if minq.col > 0 {
			check(minq, v2i{minq.row, minq.col - 1})
		}
		if minq.col < cols-1 {
			check(minq, v2i{minq.row, minq.col + 1})
		}
	}

	return dist[end.row][end.col]
}

func part1(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	graph := make([][]int, rows)
	for i := 0; i < cols; i++ {
		graph[i] = make([]int, cols)
	}

	for row, line := range lines {
		for col, c := range line {
			graph[row][col] = int(c - '0')
		}
	}

	return dijkstra(v2i{0, 0}, v2i{rows - 1, cols - 1}, graph)
}

func part2(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])
	rows5 := rows * 5
	cols5 := cols * 5

	graph := make([][]int, rows5)
	for i := 0; i < cols5; i++ {
		graph[i] = make([]int, cols5)
	}

	for row, line := range lines {
		for col, c := range line {
			for ro := 0; ro < 5; ro++ {
				for co := 0; co < 5; co++ {
					wrappedv := (((int(c-'0') + ro + co) - 1) % 9) + 1
					graph[row+ro*rows][col+co*cols] = wrappedv
				}
			}
		}
	}

	return dijkstra(v2i{0, 0}, v2i{rows5 - 1, cols5 - 1}, graph)
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
