package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.ReadFile("input6.txt")

	fmt.Println("Part 1:", findMarker(f, 4))
	fmt.Println("Part 2:", findMarker(f, 14))
}

func findMarker(f []byte, length int) int {
	for i := length; i < len(f); i++ {
		if countChars(f[i-length:i]) == length {
			return i
		}
	}
	return -1
}

func countChars(cs []byte) int {
	m := make(map[byte]bool)
	for _, c := range cs {
		m[c] = true
	}
	return len(m)
}
