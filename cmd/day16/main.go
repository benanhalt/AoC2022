package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nValves = 60

type Room struct {
	label   string
	rate    int
	tunnels []string
}

type State struct {
	time     int
	closed   [nValves]string
	location [2]string
	opening  [2]bool
}

type Result struct {
	amount  int
	actions []string
}

func main() {
	f, _ := os.ReadFile("ex.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	rooms := make(map[string]Room, len(lines))
	for _, line := range lines {
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
		rooms[label] = Room{label, rate, tunnels}
		fmt.Println(rooms[label])
	}

	posActions := func(s State, n int) []string {
		// if s.opening[n] {
		// 	return []string{"finish"}
		// }
		result := []string{}
		for _, t := range rooms[s.location[n]].tunnels {
			result = append(result, t)
		}
		if contains(s.closed[:], s.location[n]) {
			result = append(result, "open")
		}
		return result
	}

	update := func(s State, as [2]string) State {
		//fmt.Println("update", s, as)
		var closed [nValves]string
		var opening [2]bool
		var location [2]string

		for i, r := range s.closed {
			switch {
			case r == s.location[0] && as[0] == "open":
				closed[i] = "  "
			case r == s.location[1] && as[1] == "open":
				closed[i] = "  "
			default:
				closed[i] = s.closed[i]
			}
		}
		for i := range as {
			switch as[i] {
			case "open":
				opening[i] = true
				location[i] = s.location[i]
			// case "finish":
			// 	opening[i] = false
			// 	location[i] = s.location[i]
			default:
				opening[i] = false
				location[i] = as[i]
			}
		}
		result := State{
			time:     s.time - 1,
			closed:   closed,
			location: location,
			opening:  opening,
		}
		return result
	}

	added := func(s State, as [2]string) int {
		amount := 0
		for i := range as {
			if as[i] == "open" {
				amount += rooms[s.location[i]].rate * (s.time - 1)
			}
		}
		return amount
	}

	cache := make(map[State]Result)

	var optimal func(s State) Result
	optimal = func(s State) Result {
		if cached, found := cache[s]; found {
			return cached
		}
		if s.time < 1 {
			return Result{0, nil}
		}
		best := 0
		var bestActions []string
		for _, action0 := range posActions(s, 0) {
			for _, action1 := range posActions(s, 1) {
				if action0 == "open" && action1 == "open" && s.location[0] == s.location[1] {
					continue
				}
				actions := [2]string{action0, action1}
				ns := update(s, actions)
				sub := optimal(ns)
				score := sub.amount + added(s, actions)
				if s.location[0] == "DD" && s.time == 30 {
					fmt.Println(s.closed, s.time, s.location, s.opening, sub, score)
				}
				if score > best {
					best = score
					bestActions = append([]string{actions[0] + fmt.Sprint(31-s.time), actions[1]}, sub.actions...)
				}
			}
		}
		cache[s] = Result{best, bestActions}
		return cache[s]
	}

	valves := [nValves]string{}
	i := 0
	for l, r := range rooms {
		if r.rate > 0 {
			valves[i] = l
		}
		i++
	}
	fmt.Println(optimal(State{
		time:     26,
		location: [2]string{"AA", "AA"},
		opening:  [2]bool{},
		closed:   valves,
	}))
}

func contains(ss []string, s string) bool {
	for _, t := range ss {
		if t == s {
			return true
		}
	}
	return false
}
