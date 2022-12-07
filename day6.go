package main

import (
	"fmt"
	"os"
	"math/bits"
)

func main() {
	f, _ := os.ReadFile("input6.txt")

	fmt.Println("Part 1:", findMarker(f, 4))
	fmt.Println("Part 2:", findMarker(f, 14))
}

func findMarker(f []byte, length int) int {
	var mask uint32
	for i := 0; i < len(f); i++ {
		mask = mask ^ (1 << (f[i] - 'a'))
		if i >= length {
			mask = mask ^ (1 << (f[i-length] - 'a'))
		}
		if bits.OnesCount32(mask) == length {
			return i + 1
		}
	}
	return -1
}
