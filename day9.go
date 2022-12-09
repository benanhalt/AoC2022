package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	f, _ := os.ReadFile("input9.txt")
	lines := strings.Split(string(f), "\n")

	h := [2]int{}
	t := [2]int{}
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
			t = follow(h, t)
			seen[t] = true

			ks[0] = follow(h, ks[0])
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
	switch h[0] - t[0] {
	case 0:
		switch h[1] - t[1] {
		case 2:
			t[1] += 1
		case -2:
			t[1] -= 1
		}
	case 1:
		switch h[1] - t[1] {
		case 2:
			t[0] += 1
			t[1] += 1
		case -2:
			t[0] += 1
			t[1] -= 1
		}
	case -1:
		switch h[1] - t[1] {
		case 2:
			t[0] -= 1
			t[1] += 1
		case -2:
			t[0] -= 1
			t[1] -= 1
		}
	case 2:
		switch h[1] - t[1] {
		case 0:
			t[0] += 1
		case 1, 2:
			t[0] += 1
			t[1] += 1
		case -1, -2:
			t[0] += 1
			t[1] -= 1
		}
	case -2:
		switch h[1] - t[1] {
		case 0:
			t[0] -= 1
		case 1, 2:
			t[0] -= 1
			t[1] += 1
		case -1, -2:
			t[0] -= 1
			t[1] -= 1
		}
	}
	if math.Abs(float64(h[0] - t[0])) > 1 || math.Abs(float64(h[1] - t[1])) > 1 {
		fmt.Println("bad", h, t, h[0] - t[0], h[1] - t[1])
	}
	return t
}
