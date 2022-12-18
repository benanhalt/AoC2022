package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))

	lines := strings.Split(input, "\n")

	maxCoord := 0
	blocks := make(map[[3]int]bool)
	for _, line := range lines {
		coordS := strings.Split(line, ",")
		coords := [3]int{}
		for i, c := range coordS {
			coords[i], _ = strconv.Atoi(c)
			if coords[i] > maxCoord {
				maxCoord = coords[i]
			}
		}
		blocks[coords] = true
	}

	dirs := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	sa := 0
	for b := range blocks {
		for _, d := range dirs {
			if !blocks[[3]int{b[0] + d[0], b[1] + d[1], b[2] + d[2]}] {
				sa++
			}
		}
	}
	fmt.Println("Part 1:", sa)

	outside := make(map[[3]int]bool)
	q := deque.New()
	q.PushBack([3]int{maxCoord + 1, maxCoord + 1, maxCoord + 1})
	for q.Len() > 0 {
		cell := q.PopFront().([3]int)
		for _, d := range dirs {
			x := [3]int{cell[0] + d[0], cell[1] + d[1], cell[2] + d[2]}
			ok := true
			for i := range x {
				if x[i] < -12 || x[i] > maxCoord+12 {
					ok = false
				}
			}
			if ok && !outside[x] && !blocks[x] {
				outside[x] = true
				q.PushBack(x)
			}
		}
	}

	osa := 0
	for b := range blocks {
		for _, d := range dirs {
			if outside[[3]int{b[0] + d[0], b[1] + d[1], b[2] + d[2]}] {
				osa++
			}
		}
	}
	fmt.Println("Part 2:", osa)
}
