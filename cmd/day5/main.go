package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var stacksL []string
	var ops []string
	gotStacks := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			gotStacks = true
			continue
		}
		if gotStacks {
			ops = append(ops, line)
		} else {
			stacksL = append(stacksL, line)
		}
	}
	part1(stacksL, ops)
	part2(stacksL, ops)
}

func part1(stacksL []string, ops []string) {
	var stacks [9][]string
	for l := len(stacksL)-2; l >= 0; l-- {
		line := stacksL[l]
		for i := 0; i < 9; i++ {
			if c := line[i*4+1]; c != ' ' {
				stacks[i] = append(stacks[i], string(c))
			}
		}
	}
	re, _ := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	for _, op := range ops {
		match := re.FindStringSubmatch(op)
		amnt, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		for i := 0; i < amnt; i++ {
			lFrom := len(stacks[from-1])
			c := stacks[from-1][lFrom-1]
			stacks[from-1] = stacks[from-1][:lFrom-1]
			stacks[to-1] = append(stacks[to-1], c)
		}
	}
	var part1 []string
	for i := 0; i < 9; i++ {
		part1 = append(part1, stacks[i][len(stacks[i])-1])
	}
	fmt.Println("Part 1:", strings.Join(part1, ""))
}

func part2(stacksL []string, ops []string) {
	var stacks [9][]string
	for l := len(stacksL)-2; l >= 0; l-- {
		line := stacksL[l]
		for i := 0; i < 9; i++ {
			if c := line[i*4+1]; c != ' ' {
				stacks[i] = append(stacks[i], string(c))
			}
		}
	}
	re, _ := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	for _, op := range ops {
		match := re.FindStringSubmatch(op)
		amnt, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])

		lFrom := len(stacks[from-1])
		moved := stacks[from-1][lFrom-amnt:]
		stacks[from-1] = stacks[from-1][:lFrom-amnt]
		stacks[to-1] = append(stacks[to-1], moved...)
	}
	var part2 []string
	for i := 0; i < 9; i++ {
		part2 = append(part2, stacks[i][len(stacks[i])-1])
	}
	fmt.Println("Part 2:", strings.Join(part2, ""))
}
