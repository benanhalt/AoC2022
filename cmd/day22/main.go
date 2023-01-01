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

	r := 0
	c := 0
	for cc, p := range lines[0] {
		if p != ' ' {
			c = cc
			break
		}
	}

	cMax := 0
	for _, line := range lines[:len(lines)-1] {
		if len(line) > cMax {
			cMax = len(line)
		}
	}

	rMax := len(lines) - 2

	grid := make(Grid)
	for rr := 0; rr < rMax; rr++ {
		line := lines[rr]
		for cc, p := range line {
			if p != ' ' {
				grid[[2]int{rr, cc}] = p
			}
		}
	}
	solve(1, grid, rMax, cMax, path, r, c)
	solve(2, grid, rMax, cMax, path, r, c)
}

func solve(part int, grid Grid, rMax, cMax int, path string, r, c int) {
	var advance func(grid Grid, rMax, cMax, r, c int, d string) (int, int, string)
	if part == 2 {
		advance = advance2
	} else {
		advance = advance1
	}
	s := 0
	d := "R"
	for i := 0; i < len(path); i++ {
		if unicode.IsDigit(rune(path[i])) {
			s = 10*s + int(path[i]-'0')
			continue
		} else {
			for t := 0; t < s; t++ {
				r, c, d = advance(grid, rMax, cMax, r, c, d)
				// fmt.Println(r, c, d, t)
			}
			d = steer(path[i], d)
			// fmt.Println(r, c, d)
			s = 0
		}
	}
	for t := 0; t < s; t++ {
		r, c, d = advance(grid, rMax, cMax, r, c, d)
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
	//fmt.Println(r, c, d)
	fmt.Println("Part ", part, ":", q+4*(c+1)+1000*(r+1))
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

func advance1(grid Grid, rMax, cMax, r, c int, d string) (int, int, string) {
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
		return r, c, d
	}
	if grid[[2]int{rr, cc}] == '.' {
		return rr, cc, d
	}
	panic(fmt.Sprint(rr, cc, grid[[2]int{rr, cc}]))
}

func advance2(grid Grid, rMax, cMax, r, c int, d string) (int, int, string) {
	face := [2]int{r / 50, c / 50}
	rr, cc, dd := r, c, d
	switch d {
	case "R":
		cc += 1
		if cc%50 == 0 {
			switch face {
			case [2]int{0, 2}:
				cc = 50*1 + 49
				rr = 50*2 + 49 - rr
				dd = "L"
			case [2]int{1, 1}:
				cc = 50*2 + rr%50
				rr = 49
				dd = "U"
			case [2]int{2, 1}:
				cc = 50 + 2 + 49
				rr = 49 - rr%50
				dd = "L"
			case [2]int{3, 0}:
				cc = 50 + rr%50
				rr = 50*2 + 49
				dd = "U"
			}
		}
	case "L":
		cc -= 1
		if (cc+50)%50 == 49 {
			switch face {
			case [2]int{0, 1}:
				cc = 0
				rr = 50*2 + 49 - rr
				dd = "R"
			case [2]int{1, 1}:
				cc = rr % 50
				rr = 50 * 2
				dd = "D"
			case [2]int{2, 0}:
				cc = 50
				rr = 49 - rr%50
				dd = "R"
			case [2]int{3, 0}:
				cc = 50 + rr%50
				rr = 0
				dd = "D"
			}
		}
	case "U":
		rr -= 1
		if (rr+50)%50 == 49 {
			switch face {
			case [2]int{2, 0}:
				rr = 50*1 + cc%50
				cc = 50 * 1
				dd = "R"
			case [2]int{0, 1}:
				rr = 50*3 + cc%50
				cc = 0
				dd = "R"
			case [2]int{0, 2}:
				rr = 50*3 + 49
				cc = cc % 50
				dd = "U"
			}
		}
	case "D":
		rr += 1
		if rr%50 == 0 {
			switch face {
			case [2]int{3, 0}:
				rr = 0
				cc = 50*2 + cc%50
				dd = "D"
			case [2]int{2, 1}:
				rr = 50*3 + cc%50
				cc = 49
				dd = "L"
			case [2]int{0, 2}:
				rr = 50 + cc%50
				cc = 50*1 + 49
				dd = "L"
			}
		}
	}
	if grid[[2]int{rr, cc}] == '#' {
		return r, c, d
	}
	if grid[[2]int{rr, cc}] == '.' {
		return rr, cc, dd
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
