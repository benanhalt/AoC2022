package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

type Bliz struct {
	r  int
	c  int
	ch rune
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimRight(string(f), "\n"), "\n")

	grid := make(map[[2]int]rune)
	bs := make(map[[2]int][]rune)
	var start [2]int
	var end [2]int
	for r, line := range lines {
		for c, ch := range line {
			grid[[2]int{r, c}] = ch

			if r == 0 && ch == '.' {
				start = [2]int{r, c}
			}

			if r == len(lines)-1 && ch == '.' {
				end = [2]int{r, c}
			}

			if ch == '>' || ch == '<' || ch == '^' || ch == 'v' {
				bs[[2]int{r, c}] = []rune{ch}
			}
		}
	}
	bts := make(map[int]map[[2]int][]rune)
	bts[0] = bs

	type State struct {
		time int
		pos  [2]int
	}

	maxR := end[0]

	solve := func(s0 State, stop [2]int) State {
		seen := make(map[State]bool)
		q := deque.New()
		q.PushBack(s0)
		for {
			s := q.PopFront().(State)
			if seen[s] {
				continue
			}
			seen[s] = true
			if s.pos == stop {
				return s
			}
			bs := bts[s.time+1]
			if len(bs) == 0 {
				bs = make(map[[2]int][]rune)
				for l, chs := range bts[s.time] {
					for _, ch := range chs {
						var ll [2]int
						switch ch {
						case '^':
							ll = [2]int{l[0] - 1, l[1]}
							if ll[0] == start[0] {
								ll[0] = end[0] - 1
							}
						case 'v':
							ll = [2]int{l[0] + 1, l[1]}
							if ll[0] == end[0] {
								ll[0] = start[0] + 1
							}
						case '<':
							ll = [2]int{l[0], l[1] - 1}
							if ll[1] == start[1]-1 {
								ll[1] = end[1]
							}
						case '>':
							ll = [2]int{l[0], l[1] + 1}
							if ll[1] == end[1]+1 {
								ll[1] = start[1]
							}
						}
						old := bs[ll]
						bs[ll] = append(old, ch)
					}
				}
				bts[s.time+1] = bs
			}

			var p [2]int

			p = [2]int{s.pos[0], s.pos[1]}
			if chs := bs[p]; len(chs) == 0 {
				q.PushBack(State{s.time + 1, p})
			}

			p = [2]int{s.pos[0] + 1, s.pos[1]}
			if chs := bs[p]; p[0] <= maxR && len(chs) == 0 && grid[p] != '#' {
				q.PushBack(State{s.time + 1, p})
			}

			p = [2]int{s.pos[0] - 1, s.pos[1]}
			if chs := bs[p]; p[0] >= 0 && len(chs) == 0 && grid[p] != '#' {
				q.PushBack(State{s.time + 1, p})
			}

			p = [2]int{s.pos[0], s.pos[1] + 1}
			if chs := bs[p]; len(chs) == 0 && grid[p] != '#' {
				q.PushBack(State{s.time + 1, p})
			}

			p = [2]int{s.pos[0], s.pos[1] - 1}
			if chs := bs[p]; len(chs) == 0 && grid[p] != '#' {
				q.PushBack(State{s.time + 1, p})
			}
		}
	}
	s := solve(State{0, start}, end)
	fmt.Println("Part 1:", s.time)
	s = solve(s, start)
	s = solve(s, end)
	fmt.Println("Part 2:", s.time)
}
