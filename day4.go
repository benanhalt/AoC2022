package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input4.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	fullyContained := 0
	overlapping := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		s := strings.Split(line, ",")
		r1 := strings.Split(s[0], "-")
		r2 := strings.Split(s[1], "-")
		r11, _ := strconv.Atoi(r1[0])
		r12, _ := strconv.Atoi(r1[1])
		r21, _ := strconv.Atoi(r2[0])
		r22, _ := strconv.Atoi(r2[1])
		if FullyContained(r11, r12, r21, r22) {
			fullyContained += 1
		}
		if !Disjoint(r11, r12, r21, r22) {
			overlapping += 1
		}
	}
	fmt.Println("Part 1:", fullyContained)
	fmt.Println("Part 2:", overlapping)
}

func FullyContained(a1, b1, a2, b2 int) bool {
	return (a2 <= a1 && b1 <= b2) || (a1 <= a2 && b2 <= b1)
}

func Disjoint(a1, b1, a2, b2 int) bool {
	return (a1 < a2 && b1 < a2) || (a1 > b2 && b1 > b2)
}
