package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	pair   = iota
	number = iota
)

type snumber struct {
	tag    int
	num    int
	lhs    *snumber
	rhs    *snumber
	parent *snumber
}

type Consumer struct {
	line string
	pos  int
}

func (c *Consumer) snumber_read() *snumber {
	if c.line[c.pos] == '[' {
		c.pos += 1

		root := new(snumber)
		root.tag = pair
		root.lhs = c.snumber_read()
		root.lhs.parent = root
		c.pos += 1
		root.rhs = c.snumber_read()
		root.rhs.parent = root
		c.pos += 1

		return root
	} else {
		spos := c.pos

		for true {
			if c.line[c.pos] == ',' || c.line[c.pos] == ']' {
				break
			} else {
				c.pos += 1
			}
		}

		val := new(snumber)
		val.tag = number
		val.num, _ = strconv.Atoi(c.line[spos:c.pos])
		return val
	}
}

func snumber_read(line string) *snumber {
	cons := Consumer{line: line, pos: 0}
	return cons.snumber_read()
}

func (s *snumber) show() {
	var recur func(s *snumber)
	recur = func(s *snumber) {
		if s.tag == pair {
			fmt.Printf("[")
			recur(s.lhs)
			fmt.Printf(",")
			recur(s.rhs)
			fmt.Printf("]")
		} else if s.tag == number {
			fmt.Printf("%d", s.num)
		}
	}

	recur(s)
	fmt.Printf("\n")
}

func (s *snumber) add(s2 *snumber) *snumber {
	root := new(snumber)
	root.tag = pair
	root.lhs = s
	root.rhs = s2
	root.lhs.parent = root
	root.rhs.parent = root

	for true {
		if root.can_explode() {
			root.explode()
		} else {
			v := root.split()
			if !v {
				break
			}
		}
	}

	return root
}

func (s *snumber) can_explode() bool {
	var recur func(num *snumber, depth int) bool
	recur = func(num *snumber, depth int) bool {
		if num == nil {
			return false
		} else if depth == 4 && num.tag == pair {
			return true
		} else {
			return recur(num.lhs, depth+1) || recur(num.rhs, depth+1)
		}
	}

	return recur(s, 0)
}

func (s *snumber) explode() {
	exploded := false
	var expcopy *snumber
	var expnew *snumber
	var exploded_left bool

	var traverse func(num *snumber, depth int)
	traverse = func(num *snumber, depth int) {
		if num == nil {
			return
		} else if depth == 4 && num.tag == pair {
			if !exploded {
				exploded_left = (num.parent.lhs == num)

				exploded = true
				expcopy = new(snumber)
				expcopy.tag = num.tag
				expcopy.lhs = num.lhs
				expcopy.rhs = num.rhs
				expcopy.parent = num.parent

				num.tag = number
				num.num = 0
				expnew = num
			}
		} else {
			traverse(num.lhs, depth+1)
			traverse(num.rhs, depth+1)
		}
	}
	var findleftmost func(num, prev *snumber, goleft bool)

	findleftmost = func(num, prev *snumber, goleft bool) {
		var node *snumber
		if num == nil {
			return
		}

		if goleft {
			node = num.lhs
		} else {
			node = num.rhs
		}

		if node == prev {
			findleftmost(num.parent, num, goleft)
		} else if node.tag == pair {
			findleftmost(node, num, false)
		} else if node.tag == number {
			node.num += expcopy.lhs.num
		}
	}

	var findrightmost func(num, prev *snumber, goright bool)
	findrightmost = func(num, prev *snumber, goright bool) {
		var node *snumber
		if num == nil {
			return
		}

		if goright {
			node = num.rhs
		} else {
			node = num.lhs
		}

		if node == prev {
			findrightmost(num.parent, num, goright)
		} else if node.tag == pair {
			findrightmost(node, num, false)
		} else if node.tag == number {
			node.num += expcopy.rhs.num
		}
	}

	traverse(s, 0)
	findleftmost(expnew.parent, expnew, true)
	findrightmost(expnew.parent, expnew, true)
}

func (s *snumber) split() bool {
	splitted := false

	var recur func(num *snumber)
	recur = func(num *snumber) {
		if num == nil {
			return
		} else if num.tag == number && num.num >= 10 {
			if !splitted {
				splitted = true
				num.tag = pair

				num.lhs = new(snumber)
				num.lhs.tag = number
				num.lhs.num = num.num / 2
				num.lhs.parent = num

				num.rhs = new(snumber)
				num.rhs.tag = number
				num.rhs.num = num.num - num.lhs.num
				num.rhs.parent = num
			}

		} else {
			recur(num.lhs)
			recur(num.rhs)
		}
	}

	recur(s)
	return splitted
}

func (s *snumber) magnitude() int {
	if s.tag == number {
		return s.num
	} else {
		return 3*s.lhs.magnitude() + 2*s.rhs.magnitude()
	}
}

func part1(lines []string) int {
	snum := snumber_read(lines[0])
	lines = lines[1:]

	for _, line := range lines {
		snumnew := snumber_read(line)
		snumsum := snum.add(snumnew)

		snum = snumsum
	}
	snum.show()

	return snum.magnitude()
}

func part2(lines []string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	magmax := math.MinInt32
	for i := 0; i < len(lines); i++ {
		for j := i; j < len(lines); j++ {
			mag1 := snumber_read(lines[i]).add(snumber_read(lines[j])).magnitude()
			mag2 := snumber_read(lines[j]).add(snumber_read(lines[i])).magnitude()
			magmax = max(magmax, max(mag1, mag2))
		}
	}

	return magmax
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
