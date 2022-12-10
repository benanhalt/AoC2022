package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	w := len(lines[0])
	h := len(lines) - 1

	var max byte
	vis := make(map[string]bool)
	for r := 0; r < h; r++ {
		vis[fmt.Sprintf("%d,%d", r, 0)] = true
		max = lines[r][0]
		for c := 1; c < w; c++ {
			if lines[r][c] > max {
				vis[fmt.Sprintf("%d,%d", r, c)] = true
				max = lines[r][c]
			}
		}
		vis[fmt.Sprintf("%d,%d", r, w-1)] = true
		max = lines[r][w-1]
		for c := w-2; c >= 0; c-- {
			if lines[r][c] > max {
				vis[fmt.Sprintf("%d,%d", r, c)] = true
				max = lines[r][c]
			}
		}
	}
	for c := 0; c < w; c++ {
		vis[fmt.Sprintf("%d,%d", 0, c)] = true
		max = lines[0][c]
		for r:= 1; r < h; r++ {
			if lines[r][c] > max {
				vis[fmt.Sprintf("%d,%d", r, c)] = true
				max = lines[r][c]
			}
		}
		vis[fmt.Sprintf("%d,%d", h-1, c)] = true
		max = lines[h-1][c]
		for r:= h-2; r >= 0; r-- {
			if lines[r][c] > max {
				vis[fmt.Sprintf("%d,%d", r, c)] = true
				max = lines[r][c]
			}
		}
	}

	fmt.Println("Part 1:", w, h, len(vis))
	for r := 0; r < h; r++ {
		l := []string{}
		for c := 0; c < w; c++ {
			if vis[fmt.Sprintf("%d,%d", r, c)] {
				l = append(l, "*")
			} else {
				l = append(l, " ")
			}
		}
		fmt.Println(strings.Join(l, ""))
	}

	fmt.Println("Part 2:", )
}
