package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimRight(string(f), "\n"), "\n")

	directions := [4][3][2]int{
		{{-1, -1}, {-1, 0}, {-1, 1}},
		{{1, -1}, {1, 0}, {1, 1}},
		{{-1, -1}, {0, -1}, {1, -1}},
		{{-1, 1}, {0, 1}, {1, 1}},
	}

	grid := make(map[[2]int]bool)
	elves := [][2]int{}
	for r, line := range lines {
		for c, rune := range line {
			if rune == '#' {
				grid[[2]int{r, c}] = true
				elves = append(elves, [2]int{r, c})
			}
		}
	}

	for t := 0; ; t++ {
		moved := 0
		proposed := make(map[[2]int][][2]int)
		for _, e := range elves {
			r, c := e[0], e[1]
			count := 0
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					if dr == 0 && dc == 0 {
						continue
					}
					if grid[[2]int{r + dr, c + dc}] {
						count++
					}
				}
			}

			if count > 0 {
				choice := e
				for dd := 0; dd < 4; dd++ {
					d := (dd + t) % 4
					count := 0
					for _, check := range directions[d] {
						dr, dc := check[0], check[1]
						if grid[[2]int{r + dr, c + dc}] {
							count++
						}
					}
					if count == 0 {
						choice = [2]int{
							r + directions[d][1][0],
							c + directions[d][1][1],
						}
						moved++
						break
					}
				}
				proposed[choice] = append(proposed[choice], e)
			} else {
				proposed[e] = append(proposed[e], e)
			}
		}

		if moved == 0 {
			fmt.Println("Part 2:", t+1)
			break
		}
		elves = [][2]int{}
		for p, es := range proposed {
			if len(es) == 1 {
				elves = append(elves, p)
			} else {
				elves = append(elves, es...)
			}
		}

		grid = make(map[[2]int]bool, len(elves))
		for _, e := range elves {
			grid[e] = true
		}
		if t == 9 {
			minR, maxR, minC, maxC := elves[0][0], elves[0][0], elves[0][1], elves[0][1]
			for _, e := range elves {
				if e[0] < minR {
					minR = e[0]
				}
				if e[0] > maxR {
					maxR = e[0]
				}
				if e[1] < minC {
					minC = e[1]
				}
				if e[1] > maxC {
					maxC = e[1]
				}
			}
			area := (maxR - minR + 1) * (maxC - minC + 1)
			fmt.Println("Part 1:", area-len(elves))
		}
	}
}
