package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input10.txt")
	lines := strings.Split(string(f), "\n")

	cycle := 0
	x := 1
	ans := 0
	var out []string
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "addx"):
			p := strings.Split(line, " ")
			v, _ := strconv.Atoi(p[1])

			ans, out = click(&cycle, x, ans, out)
			ans, out = click(&cycle, x, ans, out)
			x += v
		case line == "noop":
			ans, out = click(&cycle, x, ans, out)
		}
	}

	fmt.Println("Part 1:", ans)
	fmt.Println("Part 2:")
	fmt.Println(strings.Join(out, ""))
}

func click(cycle *int, x int, ans int, out []string) (int, []string) {
	if *cycle%40 == 20 {
		ans += *cycle * x
	}

	if *cycle%40 == 0 {
		out = append(out, "\n")
	}

	if *cycle%40 >= x-1 && *cycle%40 <= x+1 {
		out = append(out, "#")
	} else {
		out = append(out, ".")
	}

	*cycle += 1
	return ans, out
}
