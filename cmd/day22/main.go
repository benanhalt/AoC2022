package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Grid map[[2]int]rune

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimRight(string(f), "\n"), "\n")

	path := lines[len(lines)-1]

	d := "R"
	r := 0
	c := 0
	for cc, p := range lines[0] {
		if p != ' ' {
			c = cc
			break
		}
	}
	fmt.Println(r, c, d)

	cMax := 0
	for _, line := range lines[:len(lines)-1] {
		if len(line) > cMax {
			cMax = len(line)
		}
	}

	rMax := len(lines) - 2
	fmt.Println(rMax, cMax)

	grid := make(Grid)
	for rr := 0; rr < rMax; rr++ {
		line := lines[rr]
		for cc, p := range line {
			if p != ' ' {
				grid[[2]int{rr, cc}] = p
			}
		}
	}

	s := 0
	p := make(Grid)
	for i := 0; i < len(path); i++ {
		if unicode.IsDigit(rune(path[i])) {
			s = 10*s + int(path[i]-'0')
			continue
		} else {
			for t := 0; t < s; t++ {
				r, c = advance(grid, rMax, cMax, r, c, d)
				p[[2]int{r, c}] = rune(d[0])
				// fmt.Println(r, c, d, t)
			}
			d = steer(path[i], d)
			p[[2]int{r, c}] = rune(d[0])
			// fmt.Println(r, c, d)
			s = 0
		}
	}
	for t := 0; t < s; t++ {
		r, c = advance(grid, rMax, cMax, r, c, d)
		p[[2]int{r, c}] = rune(d[0])
		// fmt.Println(r, c, d, t)
	}
	var q int
	switch d {
	case "R":
		q = 0
	case "D":
		q = 1
	case "L":
		q = 2
	case "U":
		q = 3
	default:
		panic(d)
	}
	// show(rMax, cMax, grid, p)
	fmt.Println(r, c, d)
	fmt.Println(q + 4*(c+1) + 1000*(r+1))
}

func steer(lr byte, d string) string {
	switch lr {
	case 'R':
		switch d {
		case "R":
			d = "D"
		case "D":
			d = "L"
		case "L":
			d = "U"
		case "U":
			d = "R"
		}
	case 'L':
		switch d {
		case "R":
			d = "U"
		case "D":
			d = "R"
		case "L":
			d = "D"
		case "U":
			d = "L"
		}
	}
	return d
}

func advance(grid Grid, rMax, cMax, r, c int, d string) (int, int) {
	rr, cc := r, c
	switch d {
	case "R":
		cc = (cc + 1) % cMax
		for {
			if grid[[2]int{rr, cc}] == 0 {
				cc = (cc + 1) % cMax
			} else {
				break
			}
		}
	case "L":
		cc = (cc - 1 + cMax) % cMax
		for {
			if grid[[2]int{rr, cc}] == 0 {
				cc = (cc - 1 + cMax) % cMax
			} else {
				break
			}
		}
	case "U":
		rr = (rr - 1 + rMax) % rMax
		for {
			if grid[[2]int{rr, cc}] == 0 {
				rr = (rr - 1 + rMax) % rMax
			} else {
				break
			}
		}
	case "D":
		rr = (rr + 1) % rMax
		for {
			if grid[[2]int{rr, cc}] == 0 {
				rr = (rr + 1) % rMax
			} else {
				break
			}
		}
	}
	if grid[[2]int{rr, cc}] == '#' {
		return r, c
	}
	if grid[[2]int{rr, cc}] == '.' {
		return rr, cc
	}
	panic(fmt.Sprint(rr, cc, grid[[2]int{rr, cc}]))
}

func show(rMax, cMax int, grid Grid, path Grid) {
	for r := -1; r < rMax+1; r++ {
		for c := -1; c < cMax+1; c++ {
			if path[[2]int{r, c}] == 0 {
				if grid[[2]int{r, c}] == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(string(grid[[2]int{r, c}]))
				}
			} else {
				fmt.Print(string(path[[2]int{r, c}]))
			}
		}
		fmt.Print("\n")
	}
}
