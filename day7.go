package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("input7.txt")
	lines := strings.Split(string(f), "\n")

	var path []string
	sizes := make(map[string]int)
	for _, line := range lines {
		switch {
		case line == "$ cd ..":
			path = path[:len(path)-1]
		case strings.HasPrefix(line, "$ cd"):
			path = append(path, line[5:])
		case !strings.HasPrefix(line, "$") && !strings.HasPrefix(line, "dir"):
			p := strings.Split(line, " ")
			size, _ := strconv.Atoi(p[0])
			for i := 0; i < len(path); i++ {
				sizes[strings.Join(path[:i+1], "/")] += size
			}
		}
	}

	part1 := 0
	for _, size := range sizes {
		if size <= 100000 {
			part1 += size
		}
	}
	fmt.Println("Part 1:", part1)

	capacity := 70000000
	free := capacity - sizes["/"]
	needed := 30000000

	currentMin := capacity
	for _, size := range sizes {
		if size + free >= needed && size < currentMin {
			currentMin = size
		}
	}
	fmt.Println("Part 2:", currentMin)
}
