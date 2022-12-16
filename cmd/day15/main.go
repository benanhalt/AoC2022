package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sensor struct {
	x  int
	y  int
	bx int
	by int
}

type beacon struct {
	x int
	y int
}

type Interval struct {
	min int
	max int
}

type byStart []Interval

func (is byStart) Len() int {
	return len(is)
}

func (is byStart) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is byStart) Less(i, j int) bool {
	return is[i].min < is[j].min
}

func Range(s sensor) int {
	dx := s.x - s.bx
	dy := s.y - s.by
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func intervalY(y int, s sensor) []Interval {
	dy := y - s.y
	if dy < 0 {
		dy = -dy
	}
	dx := Range(s) - dy
	if dx < 0 {
		return []Interval{}
	}
	return []Interval{{s.x - dx, s.x + dx}}
}

func union(is []Interval) []Interval {
	sort.Sort(byStart(is))
	result := []Interval{is[0]}
	for _, i1 := range is[1:] {
		i0 := &result[len(result)-1]
		switch {
		case i0.max < i1.min:
			result = append(result, i1)
		case i0.max < i1.max:
			i0.max = i1.max
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intersect(is []Interval, i Interval) []Interval {
	result := []Interval{}
	for _, i0 := range is {
		if i0.max < i.min {
			continue
		}
		if i0.min > i.max {
			continue
		}
		result = append(result, Interval{max(i0.min, i.min), min(i0.max, i.max)})
	}
	return result
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	sensors := []sensor{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		x, _ := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(words[2], "x="), ","))
		y, _ := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(words[3], "y="), ":"))
		bx, _ := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(words[8], "x="), ","))
		by, _ := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(words[9], "y="), ","))
		s := sensor{x, y, bx, by}
		sensors = append(sensors, s)
	}

	beacons := make(map[[2]int]bool)
	for _, s := range sensors {
		beacons[[2]int{s.bx, s.by}] = true
	}

	for y := 0; y <= 4000000; y++ {
		intervals := []Interval{}
		for _, s := range sensors {
			intervals = append(intervals, intervalY(y, s)...)
		}
		u := union(intervals)
		p := intersect(u, Interval{0, 4000000})
		part1 := 0
		for _, i := range u {
			part1 += i.max - i.min + 1
		}
		for b := range beacons {
			if b[1] == y {
				part1--
			}
		}
		if y == 2000000 {
			fmt.Println("Part 1:", part1)
		}
		if len(p) == 2 && p[0].max+2 == p[1].min {
			x := p[0].max + 1
			fmt.Println("Part 2:", y+x*4000000)
		}
	}
}
