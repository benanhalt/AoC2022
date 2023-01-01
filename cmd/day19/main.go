package main

import (
	"fmt"
	"os"
	"runtime"
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
		fmt.Println("optimize", bp.id)
		maxRate := [3]int{}
		for i := 0; i < 3; i++ {
			maxRate[i] = max([]int{bp.matrix[0][i], bp.matrix[1][i], bp.matrix[2][i], bp.matrix[3][i]})
		}
		cache := make(map[State]int)
		var optimal func(s State) int
		optimal = func(s State) int {
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
					nexts.resources[j] = maxRate[j] * (s.time)
				}
			}

			best := optimal(nexts)
			for i := range s.robots {
				if i == 3 || nexts.resources[i]+(nexts.time*nexts.robots[i]) <= maxRate[i]*(s.time) {
					ns := nexts
					ok := true
					for j := range ns.resources {
						ns.resources[j] -= bp.matrix[i][j]
						if s.resources[j] < bp.matrix[i][j] {
							ok = false
							break
						}
					}
					if ok {
						ns.robots[i]++
						sub := optimal(ns)
						if sub > best {
							best = sub
						}
					}
				}
			}
			cache[s] = best
			return best
		}
		return optimal(s)
	}
	ans := 0
	for i, bp := range bps {
		ans += optimize(bp, State{time: 24, robots: [4]int{1}}) * (i + 1)
		runtime.GC()
	}
	fmt.Println("Part 1:", ans)

	ans = 1
	for _, bp := range bps[:3] {
		score := optimize(bp, State{time: 32, robots: [4]int{1}})
		runtime.GC()
		ans *= score
		fmt.Println(score)
	}
	fmt.Println("Part 2:", ans)
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
