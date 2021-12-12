package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

// https://stackoverflow.com/questions/59293525/how-to-check-if-a-string-is-all-upper-or-lower-case-in-go
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func part1(lines []string) int {
	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start := parts[0]
		end := parts[1]

		if _, ok := graph[start]; !ok {
			graph[start] = make([]string, 0)
		}

		if _, ok := graph[end]; !ok {
			graph[end] = make([]string, 0)
		}

		graph[start] = append(graph[start], end)
		graph[end] = append(graph[end], start)
	}

	var recur func(node, path string) int
	recur = func(node, path string) int {
		links, _ := graph[node]

		v := 0
		if node == "end" {
			v = 1
		}

		for _, link := range links {
			contains := false
			for _, prev := range strings.Split(path, ",") {
				if !IsUpper(link) && link == prev {
					contains = true
				}
			}

			if !contains {
				v += recur(link, path+","+link)
			}
		}

		return v
	}

	return recur("start", "start")
}

func part2(lines []string) int {
	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start := parts[0]
		end := parts[1]

		if _, ok := graph[start]; !ok {
			graph[start] = make([]string, 0)
		}

		if _, ok := graph[end]; !ok {
			graph[end] = make([]string, 0)
		}

		graph[start] = append(graph[start], end)
		graph[end] = append(graph[end], start)
	}

	var recur func(node string, path []string) int
	recur = func(node string, path []string) int {
		links, _ := graph[node]

		v := 0
		if node == "end" {
			v = 1
		}

		for _, link := range links {
			invalid := false
			if !IsUpper(link) {
				visit := 0
				already_visited_twice := false
				small := make([]string, 0)

				for _, prev := range path {
					if !IsUpper(prev) {
						if link == prev {
							visit += 1
						}

						already_visited := false
						for _, s := range small {
							if s == prev {
								already_visited = true
							}
						}

						if already_visited {
							already_visited_twice = true
						} else {
							small = append(small, prev)
						}
					}
				}

				if already_visited_twice || link == "end" || link == "start" {
					invalid = visit >= 1
				}
			}

			if !invalid {
				v += recur(link, append(path, link))
			}
		}

		return v
	}

	path := make([]string, 1)
	path[0] = "start"
	return recur("start", path)
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
