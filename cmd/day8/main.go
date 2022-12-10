package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	w := len(lines[0])
	h := len(lines) - 1

	var max byte
	vis := make(map[[2]int]bool)
	for r := 0; r < h; r++ {
		vis[[2]int{r, 0}] = true
		max = lines[r][0]
		for c := 1; c < w; c++ {
			if lines[r][c] > max {
				vis[[2]int{r, c}] = true
				max = lines[r][c]
			}
		}
		vis[[2]int{r, w - 1}] = true
		max = lines[r][w-1]
		for c := w - 2; c >= 0; c-- {
			if lines[r][c] > max {
				vis[[2]int{r, c}] = true
				max = lines[r][c]
			}
		}
	}
	for c := 0; c < w; c++ {
		vis[[2]int{0, c}] = true
		max = lines[0][c]
		for r := 1; r < h; r++ {
			if lines[r][c] > max {
				vis[[2]int{r, c}] = true
				max = lines[r][c]
			}
		}
		vis[[2]int{h - 1, c}] = true
		max = lines[h-1][c]
		for r := h - 2; r >= 0; r-- {
			if lines[r][c] > max {
				vis[[2]int{r, c}] = true
				max = lines[r][c]
			}
		}
	}

	fmt.Println("Part 1:", len(vis))
	// for r := 0; r < h; r++ {
	// 	l := []string{}
	// 	for c := 0; c < w; c++ {
	// 		if vis[[2]int{r, c}] {
	// 			l = append(l, "*")
	// 		} else {
	// 			l = append(l, " ")
	// 		}
	// 	}
	// 	fmt.Println(strings.Join(l, ""))
	// }

	best := 0
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			count := 0
			for cc := c + 1; cc < w; cc++ {
				count++
				if lines[r][cc] >= lines[r][c] {
					break
				}
			}
			score := count
			count = 0

			for cc := c - 1; cc >= 0; cc-- {
				count++
				if lines[r][cc] >= lines[r][c] {
					break
				}
			}

			score *= count
			count = 0
			for rr := r + 1; rr < w; rr++ {
				count++
				if lines[rr][c] >= lines[r][c] {
					break
				}
			}

			score *= count
			count = 0
			for rr := r - 1; rr >= 0; rr-- {
				count++
				if lines[rr][c] >= lines[r][c] {
					break
				}
			}

			score *= count
			if score > best {
				best = score
			}
		}
	}
	fmt.Println("Part 2:", best)
}
