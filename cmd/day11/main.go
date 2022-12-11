package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func monkeyOp(monkey int, old uint64) uint64 {
	switch monkey {
	case 0:
		return old * 17
	case 1:
		return old + 1
	case 2:
		return old + 3
	case 3:
		return old + 5
	case 4:
		return old * old
	case 5:
		return old + 2
	case 6:
		return old + 4
	case 7:
		return old * 19
	}
	panic("no monkey")
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	monkeyItems := [8][]uint64{}
	monkeyDiv := [8]int{}
	monkeyTrue := [8]int{}
	monkeyFalse := [8]int{}
	monkeyN := -1
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey") {
			monkeyN++
		}
		if strings.HasPrefix(line, "  Starting items") {
			p := strings.Split(line, ":")
			pp := strings.Split(p[1], ",")
			for _, item := range pp {
				itemN, _ := strconv.Atoi(strings.TrimSpace(item))
				monkeyItems[monkeyN] = append(monkeyItems[monkeyN], uint64(itemN))
			}
		}
		if strings.HasPrefix(line, "  Test:") {
			p := strings.Split(line, "by")
			monkeyDiv[monkeyN], _ = strconv.Atoi(strings.TrimSpace(p[1]))
		}
		if strings.HasPrefix(line, "    If true: ") {
			p := strings.Split(line, " ")
			monkeyTrue[monkeyN], _ = strconv.Atoi(strings.TrimSpace(p[len(p)-1]))
		}
		if strings.HasPrefix(line, "    If false: ") {
			p := strings.Split(line, " ")
			monkeyFalse[monkeyN], _ = strconv.Atoi(strings.TrimSpace(p[len(p)-1]))
		}
	}
	fmt.Println(monkeyItems)
	fmt.Println(monkeyDiv)
	fmt.Println(monkeyTrue)
	fmt.Println(monkeyFalse)

	lcm := uint64(1)
	for m := 0; m < 8; m++ {
		lcm *= uint64(monkeyDiv[m])
	}

	var counts [8]int
	for round := 0; round < 10000; round++ {
		for m := 0; m < 8; m++ {
			for _, item := range monkeyItems[m] {
				counts[m]++
				worry := monkeyOp(m, item)
				worry %= lcm
				var throwTo int
				if worry%uint64(monkeyDiv[m]) == 0 {
					throwTo = monkeyTrue[m]
				} else {
					throwTo = monkeyFalse[m]
				}
				monkeyItems[throwTo] = append(monkeyItems[throwTo], worry)
			}
			monkeyItems[m] = []uint64{}
		}
	}

	sort.Ints(counts[:])
	fmt.Println("Part 1:", counts[7]*counts[6])
	fmt.Println("Part 2:")
}
