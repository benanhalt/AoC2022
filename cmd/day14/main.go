package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	var bottom int
	grid := make(map[[2]int]byte)
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for i := 1; i < len(points); i++ {
			start := strings.Split(points[i-1], ",")
			x0, _ := strconv.Atoi(start[0])
			y0, _ := strconv.Atoi(start[1])
			end := strings.Split(points[i], ",")
			x1, _ := strconv.Atoi(end[0])
			y1, _ := strconv.Atoi(end[1])
			if x0 > x1 {
				x0, x1 = x1, x0
			}
			if y0 > y1 {
				y0, y1 = y1, y0
			}
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					grid[[2]int{x, y}] = '#'
					if y > bottom {
						bottom = y
					}
				}
			}
		}
	}

	// need a fresh grid for part 2
	grid2 := make(map[[2]int]byte)
	for k, v := range grid {
		grid2[k] = v
	}

	for done := false; !done; {
		sX, sY := 500, 0
		for {
			if sY > bottom {
				done = true
				break
			}
			if _, p := grid[[2]int{sX, sY + 1}]; !p {
				sY++
			} else if _, p := grid[[2]int{sX - 1, sY + 1}]; !p {
				sX--
				sY++
			} else if _, p := grid[[2]int{sX + 1, sY + 1}]; !p {
				sX++
				sY++
			} else {
				grid[[2]int{sX, sY}] = 'o'
				break
			}
		}
	}

	part1 := 0
	for _, v := range grid {
		if v == 'o' {
			part1++
		}
	}
	fmt.Println("Part 1:", part1)

	grid = grid2
	for done := false; !done; {
		sX, sY := 500, 0
		for {
			if sY == bottom+1 {
				grid[[2]int{sX, sY}] = 'o'
				break
			} else if _, p := grid[[2]int{sX, sY + 1}]; !p {
				sY++
			} else if _, p := grid[[2]int{sX - 1, sY + 1}]; !p {
				sX--
				sY++
			} else if _, p := grid[[2]int{sX + 1, sY + 1}]; !p {
				sX++
				sY++
			} else {
				grid[[2]int{sX, sY}] = 'o'
				if sX == 500 && sY == 0 {
					done = true
				}
				break
			}
		}
	}
	part2 := 0
	for _, v := range grid {
		if v == 'o' {
			part2++
		}
	}
	fmt.Println("Part 2:", part2)
}
