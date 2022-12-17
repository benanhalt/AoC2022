package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//f, _ := os.ReadFile("ex.txt")
	//input := strings.TrimSpace(string(f))

	fb, _ := os.ReadFile("blocks.txt")
	lines := strings.Split(strings.TrimSpace(string(fb)), "\n")
	fmt.Println(lines)

	n := 0
	r := 0
	maxC := 0
	blocks := make(map[[3]int]bool)
	blockH := make(map[int]int)
	blockW := make(map[int]int)
	for _, line := range lines {
		if line == "" {
			blockH[n] = r
			blockW[n] = maxC + 1
			n++
			r = 0
			maxC = 0
			continue
		}
		for c, char := range line {
			blocks[[3]int{n, r, c}] = char == '#'
			if c > maxC {
				maxC = c
			}
		}
		r++
	}
	blockH[n] = r
	blockW[n] = maxC + 1

	type blockPlacement struct {
		blockN int
		dr     int
		dc     int
	}

	// var stack []blockPlacement

	// overlaps := func(p blockPlacement, stack []blockPlacement) bool {
	// 	for brc := range blocks {
	// 		if brc[0] == p.blockN {
	// 			r, c := brc[1], brc[2]
	// 			for _, pp := range stack {
	// 				if blocks[[3]int{pp.blockN, pp.dr + r, pp.dc + c}] && blocks[[3]int{p.blockN, p.dr + r, p.dc + c}] {
	// 					return true
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return false
	// }

	// draw := func(p blockPlacement) {
	// 	for r := 0; r < p.dr+blockH[p.blockN]; r++ {
	// 		for c := 0; c < p.dc+blockW[p.blockN]; c++ {
	// 			if blocks[[3]int{p.blockN, r - p.dr, c - p.dc}] {
	// 				fmt.Print("#")
	// 			} else {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 		fmt.Print("\n")
	// 	}
	// }

	drawStack := func(ps []blockPlacement, h, w int) {
		for r := h - 1; r >= 0; r-- {
			for c := 0; c < w; c++ {
				o := false
				for _, p := range ps {
					o = o || blocks[[3]int{p.blockN, r - p.dr, c - p.dc}]
				}
				if o {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
	}

	stack := []blockPlacement{}
	for i := range blockH {
		stack = append(stack, blockPlacement{blockN: i, dr: i * 5, dc: 2})
	}

	drawStack(stack, 25, 7)
}
