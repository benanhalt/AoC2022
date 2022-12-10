package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	h := [2]int{}
	ks := [9][2]int{}
	seen := make(map[[2]int]bool)
	seen2 := make(map[[2]int]bool)
	for _, line := range lines {
		if line == "" {
			continue
		}
		p := strings.Split(line, " ")
		dir := p[0]
		steps, _ := strconv.Atoi(p[1])

		var dx, dy int
		switch dir {
		case "U":
			dx, dy = 0, 1
		case "D":
			dx, dy = 0, -1
		case "L":
			dx, dy = -1, 0
		case "R":
			dx, dy = 1, 0
		}

		for s := 0; s < steps; s++ {
			h[0] += dx
			h[1] += dy

			ks[0] = follow(h, ks[0])
			seen[ks[0]] = true
			for i := 1; i < 9; i++ {
				ks[i] = follow(ks[i-1], ks[i])
			}
			seen2[ks[8]] = true
		}
	}
	fmt.Println("Part 1:", len(seen))
	fmt.Println("Part 2:", len(seen2))
}

func follow(h, t [2]int) [2]int {
	if math.Abs(float64(h[0]-t[0])) < 2 && math.Abs(float64(h[1]-t[1])) < 2 {
		// within 1 step in both dimensions
		return t
	}

	dirs := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1}}

	// If t has to move, choose the direction that minimizes the
	// distance in both dimensions. IOW, adjacent if possible.
	best, ok := t, false
	for _, d := range dirs {
		tt := [2]int{t[0] + d[0], t[1] + d[1]}
		if math.Abs(float64(h[0]-tt[0]))+math.Abs(float64(h[1]-tt[1])) < math.Abs(float64(h[0]-best[0]))+math.Abs(float64(h[1]-best[1])) {
			best = tt
			ok = true
		}
	}

	if !ok {
		// There was no move better that the starting position!
		panic("bad")
	}
	if math.Abs(float64(h[0]-best[0])) > 1 || math.Abs(float64(h[1]-best[1])) > 1 {
		fmt.Println("bad", h, t, h[0]-t[0], h[1]-t[1])
	}
	return best
}
