package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	label   string
	rate    int
	tunnels []string
}

type Args struct {
	time  int
	open  string
	label string
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	rooms := make(map[string]room, len(lines))
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
		rooms[label] = room{label, rate, tunnels}
		fmt.Println(rooms[label])
	}

	cache := make(map[Args]result)
	var optimal func(args Args) result

	optimal = func(args Args) result {
		if cached, found := cache[args]; found {
			return cached
		}
		time, open, label := args.time, args.open, args.label

		if time < 1 {
			return result{0, []string{}}
		}
		room := rooms[label]

		var bestPath []string
		var best int
		nowOpen := strings.Split(open, " ")
		if !contains(nowOpen, label) && room.rate > 0 {
			nowOpen = append(nowOpen, label)
			sort.Strings(nowOpen)

			for _, tunnel := range room.tunnels {
				sub := optimal(Args{time - 2, strings.Join(nowOpen, " "), tunnel})
				amount := sub.amount + room.rate*(time-1)
				if amount > best {
					best = amount
					bestPath = append(sub.path, fmt.Sprint("open", label))
				}
			}
		}
		for _, tunnel := range room.tunnels {
			sub := optimal(Args{time - 1, open, tunnel})
			if sub.amount > best {
				best = sub.amount
				bestPath = append(sub.path, label)
			}
		}

		cache[args] = result{best, bestPath}
		//fmt.Println(args, best, bestPath)
		return result{best, bestPath}
	}
	fmt.Println(optimal(Args{30, "", "AA"}))
}

func contains(ss []string, s string) bool {
	for _, t := range ss {
		if t == s {
			return true
		}
	}
	return false
}

type result struct {
	amount int
	path   []string
}

// func optimal(rooms *map[string]room, cache *map[args]result, time int, label string) result {
// 	if cached, found := (*cache)[args{time, label}]; found {
// 		return cached
// 	}

// 	if time < 1 {
// 		return result{0, []string{}}
// 	}
// 	room := (*rooms)[label]

// 	var bestPath []string
// 	var best int
// 	for _, tunnel := range room.tunnels {
// 		sub := optimal(rooms, cache, time-1, tunnel)
// 		if sub.amount > best {
// 			best = sub.amount
// 			bestPath = append(sub.path, label)
// 		}
// 	}

// 	for _, elem := range bestPath {
// 		if elem == fmt.Sprint("open", label) {
// 			(*cache)[args{time, label}] = result{best, bestPath}
// 			fmt.Println(time, label, room.rate, best, bestPath)
// 			return result{best, bestPath}
// 		}
// 	}

// 	for _, tunnel := range room.tunnels {
// 		sub := optimal(rooms, cache, time-2, tunnel)
// 		amount := sub.amount + room.rate*(time-1)
// 		if amount > best {
// 			best = amount
// 			bestPath = append(sub.path, fmt.Sprint("open", label))
// 		}
// 	}

// 	(*cache)[args{time, label}] = result{best, bestPath}
// 	fmt.Println(time, label, room.rate, best, bestPath)
// 	return result{best, bestPath}
// }
