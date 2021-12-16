package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Consumer struct {
	data    []int
	wordPos int
	bitPos  int
}

func ConsumerFromDataLine(dataLine string) Consumer {
	var cons Consumer
	cons.data = make([]int, (len(dataLine)+7)/8)
	for i := 0; i < len(cons.data); i++ {
		val64 := int64(0)
		if len(dataLine) < (i+1)*8 {
			val64, _ = strconv.ParseInt(dataLine[i*8:], 16, 64)
			val64 <<= 4 * ((i+1)*8 - len(dataLine))
		} else {
			val64, _ = strconv.ParseInt(dataLine[i*8:(i+1)*8], 16, 64)
		}
		cons.data[i] = int(val64)
	}

	return cons
}

func (c *Consumer) ConsumeBits(n int) int {
	get := func(word, bit, n int) int {
		mask := ((1 << n) - 1)
		maskS := mask << (32 - n - bit)
		val := c.data[word] & maskS
		valS := val >> (32 - n - bit)
		return valS
	}

	if c.bitPos+n >= 32 {
		lhs := 32 - c.bitPos
		rhs := n - lhs

		lv := get(c.wordPos, c.bitPos, lhs)
		rv := get(c.wordPos+1, 0, rhs)

		c.bitPos = rhs
		c.wordPos += 1

		return rv | (lv << rhs)
	} else {
		valS := get(c.wordPos, c.bitPos, n)
		c.bitPos += n
		if c.bitPos == 32 {
			c.bitPos = 0
			c.wordPos += 1
		}
		return valS
	}

	return 0
}

func (c *Consumer) ConsumedBits() int {
	return c.wordPos*32 + c.bitPos
}

func ParsePacket(cons *Consumer, fn func(ver, typeId, lit int, values []int) int) int {
	ver := cons.ConsumeBits(3)
	typeId := cons.ConsumeBits(3)
	lit := 0

	var values []int
	switch typeId {
	case 4:
		for true {
			cont := cons.ConsumeBits(1) == 1
			lit = (lit << 4) | cons.ConsumeBits(4)
			if !cont {
				break
			}
		}
	default:
		label := cons.ConsumeBits(1)

		if label == 0 {
			bitCnt := cons.ConsumeBits(15)
			consumed := cons.ConsumedBits()
			values = make([]int, 0)
			for cons.ConsumedBits() != bitCnt+consumed {
				v := ParsePacket(cons, fn)
				values = append(values, v)
			}
		} else {
			subPacketNum := cons.ConsumeBits(11)
			values = make([]int, subPacketNum)
			for i := 0; i < subPacketNum; i++ {
				v := ParsePacket(cons, fn)
				values[i] = v
			}
		}
	}

	return fn(ver, typeId, lit, values)
}

func part1(lines []string) int {
	dataLine := lines[0]
	cons := ConsumerFromDataLine(dataLine)

	fn := func(ver, typeId, lit int, values []int) int {
		acc := ver
		for _, val := range values {
			acc += val
		}
		return acc
	}

	return ParsePacket(&cons, fn)
}

func part2(lines []string) int {
	dataLine := lines[0]
	cons := ConsumerFromDataLine(dataLine)

	fn := func(ver, typeId, lit int, values []int) int {
		boolToInt := func(b bool) int {
			if b {
				return 1
			} else {
				return 0
			}
		}
		acc := 0

		switch typeId {
		case 0:
			acc = lit
		case 1:
			acc = 1
		case 2:
			acc = math.MaxInt32
		case 3:
			acc = math.MinInt32
		case 4:
			acc = lit
		case 5:
			acc = 1
		case 6:
			acc = 1
		case 7:
			acc = 1
		}

		if typeId == 0 || typeId == 1 || typeId == 2 || typeId == 3 {
			for _, val := range values {
				switch typeId {
				case 0:
					acc += val
				case 1:
					acc *= val
				case 2:
					if val < acc {
						acc = val
					}
				case 3:
					if val > acc {
						acc = val
					}
				}
			}
		} else if typeId == 5 || typeId == 6 || typeId == 7 {
			switch typeId {
			case 5:
				acc = boolToInt((acc == 1) && (values[0] > values[1]))
			case 6:
				acc = boolToInt((acc == 1) && (values[0] < values[1]))
			case 7:
				acc = boolToInt((acc == 1) && (values[0] == values[1]))
			}
		}
		return acc
	}

	return ParsePacket(&cons, fn)
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
