package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	total := 0
	total2 := 0
	var curIntersection []byte
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		comp1 := []byte(line[:len(line)/2])
		comp2 := []byte(line[len(line)/2:])
		b :=  intersect(comp1, comp2)[0]
		total += int(priority(b))

		if i % 3 == 0 {
			curIntersection = []byte(line)
		} else {
			curIntersection = intersect(curIntersection, []byte(line))
		}

		if i % 3 == 2 {
			total2 += int(priority(curIntersection[0]))
		}
		i += 1
	}
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
}

func intersect(s, t []byte) []byte {
	m := make(map[byte]bool)
	for _, b := range t {
		m[b] = true
	}

	var result []byte
	for _, b := range s {
		if m[b] {
			result = append(result, b)
		}
	}
	return result
}

func priority(b byte) byte {
	if b <= byte('Z') {
		return 1 + b - byte('A') + 26
	}
	return 1 + b - byte('a')
}
