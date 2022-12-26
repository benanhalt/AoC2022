package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nValves = 15

type Room struct {
	idx     int
	label   string
	rate    int
	tunnels []string
}

type State struct {
	closed   uint64
	location int
	opening  bool
	time     int
	players  int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	totalRate := 0
	rooms := make(map[string]Room, len(lines))
	roomsByIdx := make([]Room, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		label := words[1]
		rate, _ := strconv.Atoi(
			strings.TrimSuffix(
				strings.TrimPrefix(words[4], "rate="),
				";"))
		tunnels := []string{}
		for _, word := range words[9:] {
			tunnels = append(tunnels, strings.TrimSuffix(word, ","))
		}
		rooms[label] = Room{i, label, rate, tunnels}
		roomsByIdx[i] = rooms[label]
		totalRate += rate
	}
	fmt.Println("Part 1:", solve(rooms, roomsByIdx, 30, 1))
	fmt.Println("Part 2:", solve(rooms, roomsByIdx, 26, 2))
}

func solve(rooms map[string]Room, roomsByIdx []Room, time, players int) int {
	posActions := func(s State, n int) []string {
		result := []string{}
		for _, t := range roomsByIdx[s.location].tunnels {
			result = append(result, t)
		}
		if s.closed&(1<<s.location) != 0 {
			result = append(result, "open")
		}
		return result
	}

	update := func(s State, a string) State {
		closed := s.closed
		var opening bool
		var location int

		for i := 0; i < len(rooms); i++ {
			switch {
			case i == s.location && a == "open":
				closed = closed & ^(1 << i)
			case i == s.location && a == "open":
				closed = closed & ^(1 << i)
			}
		}
		switch a {
		case "open":
			opening = true
			location = s.location
		default:
			opening = false
			location = rooms[a].idx
		}

		result := State{
			time:     s.time - 1,
			closed:   closed,
			location: location,
			opening:  opening,
			players:  s.players,
		}
		return result
	}

	added := func(s State, a string) int {
		amount := 0
		if a == "open" {
			amount += roomsByIdx[s.location].rate * (s.time - 1)

		}
		return amount
	}

	cache := make(map[State]int, 10000000)

	var optimal func(s State) int
	optimal = func(s State) int {
		// if s.location[0] > s.location[1] {
		// 	s.location[0], s.location[1] = s.location[1], s.location[0]
		// 	s.opening[0], s.opening[1] = s.opening[1], s.opening[0]
		// }
		if cached, found := cache[s]; found {
			return cached
		}
		if s.time < 1 {
			if s.players == 0 {
				return 0
			}
			return optimal(State{
				time:     26,
				location: rooms["AA"].idx,
				opening:  false,
				closed:   s.closed,
				players:  s.players - 1,
			})
		}

		best := 0
		for _, action := range posActions(s, 0) {
			// action1 := ""
			ns := update(s, action)
			sub := optimal(ns)
			score := sub + added(s, action)

			if score > best {
				best = score
			}
		}

		cache[s] = best

		if len(cache)%1000000 == 0 {
			fmt.Println("Seen", len(cache))
		}
		return cache[s]
	}

	var valves uint64
	for _, r := range roomsByIdx {
		if r.rate > 0 {
			valves = valves | (1 << r.idx)
		}
	}
	return optimal(State{
		time:     time,
		location: rooms["AA"].idx,
		opening:  false,
		closed:   valves,
		players:  players - 1,
	})
}

func contains(ss []string, s string) bool {
	for _, t := range ss {
		if t == s {
			return true
		}
	}
	return false
}
