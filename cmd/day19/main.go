package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Blueprint struct {
	id     int
	matrix [4][4]int
}

type State struct {
	time      int
	resources [4]int
	robots    [4]int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	//Blueprint 1: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 2 ore and 15 obsidian.
	var bps []Blueprint
	for _, line := range lines {
		bp := Blueprint{}
		words := strings.Split(line, " ")
		bp.id, _ = strconv.Atoi(strings.TrimSuffix(words[1], ":"))
		bp.matrix[0][0], _ = strconv.Atoi(words[6])
		bp.matrix[1][0], _ = strconv.Atoi(words[12])
		bp.matrix[2][0], _ = strconv.Atoi(words[18])
		bp.matrix[2][1], _ = strconv.Atoi(words[21])
		bp.matrix[3][0], _ = strconv.Atoi(words[27])
		bp.matrix[3][2], _ = strconv.Atoi(words[30])
		bps = append(bps, bp)
	}

	optimize := func(bp Blueprint, s State) int {
		maxRate := [3]int{}
		for i := 0; i < 3; i++ {
			maxRate[i] = max([]int{bp.matrix[0][i], bp.matrix[1][i], bp.matrix[2][i], bp.matrix[3][i]})
		}
		fmt.Println(maxRate)
		cache := make(map[State]int)
		var optimal func(s State) int
		optimal = func(s State) int {
			// fmt.Println(s)
			if s.time < 1 {
				return s.resources[3]
			}
			if r, found := cache[s]; found {
				return r
			}
			nexts := s
			nexts.time--
			for j := range nexts.resources {
				nexts.resources[j] += nexts.robots[j]
				if j != 3 && nexts.resources[j] > maxRate[j]*(s.time) {
					// fmt.Println("too much", j, nexts.resources, maxRate[j]*nexts.time)
					nexts.resources[j] = maxRate[j] * (s.time)
				}
			}

			best := optimal(nexts)
			for i := range s.robots {
				if i == 3 || nexts.resources[i]+(nexts.time*nexts.robots[i]) <= maxRate[i]*(s.time) {
					ns := nexts
					ok := true
					for j := range ns.resources {
						//fmt.Println("ns.resources[i] -= bp.matrix[i][j]", i, j, ns.resources[i], bp.matrix[i][j])
						ns.resources[j] -= bp.matrix[i][j]
						if s.resources[j] < bp.matrix[i][j] {
							ok = false
							break
						}
					}
					if ok {
						ns.robots[i]++
						// fmt.Println("add robot", i, s, ns)
						sub := optimal(ns)
						if sub > best {
							best = sub
						}
					}
				}
				// fmt.Println("i ns", i, ns)
			}
			cache[s] = best
			return best
		}
		return optimal(s)
	}
	ans := 0
	// for i, bp := range bps {
	// 	ans += optimize(bp, State{time: 24, robots: [4]int{1}}) * (i + 1)
	// }
	// fmt.Println(ans)

	ans = 1
	for _, bp := range bps[1:] {
		score := optimize(bp, State{time: 32, robots: [4]int{1}})
		ans *= score
		fmt.Println(score)
	}
	fmt.Println(ans)
	// fmt.Println(optimize(bps[0], State{time: 24, robots: [4]int{1}}))
	// fmt.Println(optimize(bps[0], State{time: 1, robots: [4]int{1, 4, 2, 2}, resources: [4]int{5, 37, 6, 7}}))
}

func max(xs []int) int {
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max
}
