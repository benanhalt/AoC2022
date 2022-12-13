package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

type state struct {
	r     int
	c     int
	steps int
	seen  map[[2]int]bool
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	var start [3]int
	for r, line := range lines {
		for c, b := range line {
			if b == 'E' {
				start = [3]int{r, c, 0}
			}
		}
	}

	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1}}

	var q deque.Deque
	var part1, part2 bool
	seen := make(map[[2]int]bool)
	q.PushBack(state{r: start[0], c: start[1], steps: 0, seen: map[[2]int]bool{{start[0], start[1]}: true}})
	for i := 0; ; i++ {
		if q.Len() < 1 {
			for r := range lines {
				for c := range lines[0] {
					if seen[[2]int{r, c}] {
						fmt.Print(" ")
					} else {
						fmt.Print(string(lines[r][c]))
					}
				}
				fmt.Println("")
			}
		}
		s := q.PopFront().(state)
		seen[[2]int{s.r, s.c}] = true
		// if i%1000 == 0 {
		// 	fmt.Println(i, s.r, s.c, string(lines[s.r][s.c]), s.steps, q.Len(), len(seen))
		// }

		if lines[s.r][s.c] == 'S' && !part1 {
			fmt.Println("Part 1:", s.steps)
			part1 = true
		}
		if lines[s.r][s.c] == 'a' && !part2 {
			fmt.Println("Part 2:", s.steps)
			part2 = true
		}
		if part1 && part2 {
			break
		}
		for _, d := range dirs {
			rr, cc := s.r+d[0], s.c+d[1]
			if rr >= 0 && rr < len(lines) && cc >= 0 && cc < len(lines[0]) {
				b, bb := lines[s.r][s.c], lines[rr][cc]
				if b == 'E' {
					b = 'z'
				}
				if bb == 'S' {
					bb = 'z'
				}
				// if b == 'n' && bb == 'n' {
				// 	fmt.Println(s.r, s.c, string(bb), rr, cc, q.Len(), seen[[2]int{rr, cc}])
				// }
				if bb >= b-1 && !seen[[2]int{rr, cc}] {
					// seen := make(map[[2]int]bool)
					// for k, v := range s.seen {
					// 	seen[k] = v
					// }
					seen[[2]int{rr, cc}] = true
					q.PushBack(state{r: rr, c: cc, steps: s.steps + 1, seen: seen})
				}
			}
		}
	}
}
