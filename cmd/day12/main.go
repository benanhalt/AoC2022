package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

type state struct {
	r    int
	c    int
	path map[[2]int]bool
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
	q.PushBack(state{r: start[0], c: start[1], path: map[[2]int]bool{{start[0], start[1]}: true}})
	for i := 0; ; i++ {
		if q.Len() < 1 {
			// no solution
			render(lines, seen)
		}
		s := q.PopFront().(state)
		seen[[2]int{s.r, s.c}] = true

		if lines[s.r][s.c] == 'S' && !part1 {
			fmt.Println("Part 1:", len(s.path)-1)
			render(lines, s.path)
			part1 = true
		}
		if lines[s.r][s.c] == 'a' && !part2 {
			fmt.Println("Part 2:", len(s.path)-1)
			render(lines, s.path)
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
				if bb >= b-1 && !seen[[2]int{rr, cc}] {
					path := make(map[[2]int]bool)
					for k, v := range s.path {
						path[k] = v
					}
					seen[[2]int{rr, cc}] = true
					path[[2]int{rr, cc}] = true
					q.PushBack(state{r: rr, c: cc, path: path})
				}
			}
		}
	}
}

func render(lines []string, seen map[[2]int]bool) {
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
