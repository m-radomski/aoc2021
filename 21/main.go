package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type GameState struct {
	pos    [2]int
	points [2]int
	turn   int
	mult   int64
}

func part1(lines []string) int {
	player_count := len(lines)
	player_position_array := make([]int, player_count)
	player_points_array := make([]int, player_count)
	for i, line := range lines {
		pos, _ := strconv.Atoi(strings.Split(line, ": ")[1])
		player_position_array[i] = pos
	}

	dice := 1
	roll_count := 0

	roll := func() int {
		roll_count += 1
		v := dice
		dice += 1
		if dice == 101 {
			dice = 1
		}

		return v
	}

	player_index := 0
	for true {
		roll1 := roll()
		roll2 := roll()
		roll3 := roll()

		pos := player_position_array[player_index]
		pos += roll1 + roll2 + roll3
		pos = ((pos - 1) % 10) + 1
		player_position_array[player_index] = pos
		player_points_array[player_index] += pos

		next_player_index := (player_index + 1) % player_count
		if player_points_array[player_index] >= 1000 {
			return roll_count * player_points_array[next_player_index]
		}
		player_index = next_player_index
	}

	return -1
}

func part2(lines []string) int64 {
	var gs GameState
	gs.pos[0], _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	gs.pos[1], _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	gs.points[0] = 0
	gs.points[1] = 0
	gs.mult = 1

	var wins [2]int64

	mult_map := map[int]int64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	var play func(gs GameState)
	play = func(gs GameState) {
		// You are going to have 7 different game states
		// because spliting the universe is going to give you
		// only 7 unique sums of 3 dices
		// The only way we are going to account for that is to have a multiplier
		for points := 3; points <= 9; points++ {
			new_gs := gs
			new_gs.pos[new_gs.turn] = (((new_gs.pos[new_gs.turn] + points) - 1) % 10) + 1
			new_gs.points[new_gs.turn] += new_gs.pos[new_gs.turn]
			if new_gs.points[new_gs.turn] >= 21 {
				wins[new_gs.turn] += mult_map[points] * gs.mult
			} else {
				new_gs.mult *= mult_map[points]
				new_gs.turn = (new_gs.turn + 1) % 2
				play(new_gs)
			}
		}
	}

	play(gs)

	if wins[0] > wins[1] {
		return wins[0]
	} else {
		return wins[1]
	}
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
