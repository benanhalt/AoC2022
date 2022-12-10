package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	total := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		l := strings.Split(line, " ")

		var elfPlay int
		switch l[0] {
		case "A": elfPlay = 0
		case "B": elfPlay = 1
		case "C": elfPlay = 2
		}

		var myPlay int
		switch l[1] {
		case "X": myPlay = 0
		case "Y": myPlay = 1
		case "Z": myPlay = 2
		}

		score := myPlay + 1
		if myPlay == (elfPlay + 1) % 3 {
			score += 6
		} else if myPlay == elfPlay {
			score += 3
		}
		total += score

		var myPlay2 int
		switch l[1] {
		case "X": myPlay2 = (elfPlay + 2) % 3
		case "Y": myPlay2 = elfPlay
		case "Z": myPlay2 = (elfPlay + 1) % 3
		}

		score2 := myPlay2 + 1
		if myPlay2 == (elfPlay + 1) % 3 {
			score2 += 6
		} else if myPlay2 == elfPlay {
			score2 += 3
		}
		total2 += score2
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
}
