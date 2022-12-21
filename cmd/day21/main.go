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

	jobs := make(map[string][]string, len(lines))
	for _, line := range lines {
		words := strings.Split(line, " ")
		v := strings.Trim(words[0], ":")
		jobs[v] = words[1:]
	}

	var eval func(v string, part2 bool, humn float64) float64
	eval = func(v string, part2 bool, humn float64) float64 {
		if part2 && v == "humn" {
			return humn
		}

		words := jobs[v]
		if len(words) == 1 {
			r, _ := strconv.Atoi(words[0])
			return float64(r)
		}
		a := eval(words[0], part2, humn)
		b := eval(words[2], part2, humn)
		op := words[1]
		if part2 && v == "root" {
			op = "-"
		}
		switch op {
		case "+":
			return a + b
		case "*":
			return a * b
		case "-":
			return a - b
		case "/":
			return a / b
		default:
			panic("what")
		}
	}

	humn := 5.0
	fmt.Println("Part 1:", int(eval("root", false, humn)))

	for {
		r := eval("root", true, humn)
		dr := eval("root", true, humn+1) - r

		humn = humn - r/dr
		fmt.Println("Part 2:", int(humn))
		if r/dr == 0.0 {
			break
		}
	}

}
