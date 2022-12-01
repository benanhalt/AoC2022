package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	f, _ := os.Open("input1.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sums := []float64{}
	currentSum := 0.0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			sums = append(sums, currentSum)
			currentSum = 0.0
		} else {
			v, _ := strconv.ParseFloat(l, 64)
			currentSum += v
		}
	}

	f.Close()

	sort.Float64s(sums)
	fmt.Println("Part 1:", sums[len(sums)-1])

	total := 0.0
	for _, v := range sums[len(sums)-3:] {
		total += v
	}
	fmt.Println("Part 2:", total)
}
