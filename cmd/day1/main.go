package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

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

	sort.Sort(sort.Reverse(sort.Float64Slice(sums)))
	fmt.Println("Part 1:", sums[0])
	fmt.Println("Part 2:", sums[0]+sums[1]+sums[2])
}
