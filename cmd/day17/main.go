package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(f))
	fmt.Printf("'%s'\n", input)

	fb, _ := os.ReadFile("blocks.txt")
	lines := strings.Split(strings.TrimSpace(string(fb)), "\n")
	fmt.Println(lines)

	n, r := 0, 0
	block := [5][4]uint32{}
	thisBlock := []uint32{}

	for _, line := range lines {
		var br uint32
		if line == "" {
			for i := 0; i < r; i++ {
				block[n][i] = thisBlock[r-i-1]
			}
			n++
			r = 0
			thisBlock = []uint32{}
			continue
		}
		for c, p := range line {
			if p == '#' {
				br += 1 << (5 - c)
			}
		}
		thisBlock = append(thisBlock, br)
		r++
	}
	for i := 0; i < r; i++ {
		block[n][i] = thisBlock[r-i-1]
	}
	fmt.Println(block)

	seq := []uint32{0, 0, 0, 0, 0, 0, 0, 0}

	jet := 0
	for y := 0; y < 2022; y++ {
		b := block[y%5]
		//		print(b[:], 0)
		i := len(seq) - 1
		for ; ; i-- {
			if i == -1 || seq[i] != 0 {
				break
			}
		}
		i += 4
		fmt.Println(i)
		shift := 0
		print2(seq, b, shift, i)
		for i >= 0 {
			prevShift := shift
			if input[jet] == '>' {
				shift--
				fmt.Println("right")
			} else if input[jet] == '<' {
				shift++
				fmt.Println("left")
			} else {
				panic("stohsut")
			}
			for j := 0; j < 4; j++ {
				bj := Shift(b[j], shift)
				if (bj%2 == 1) || (bj > 255) {
					// fmt.Println("hit wall", shift, b, b[j], bj)
					// print(b[:], shift)
					shift = prevShift
					break
				}
				var intersect uint32
				for j := 0; j < 4; j++ {
					if i+j < len(seq) {
						intersect += Shift(b[j], shift) & seq[i+j]
					}
				}
				if intersect != 0 {
					shift = prevShift
					break
				}
			}
			print2(seq, b, shift, i)
			jet = (jet + 1) % len(input)
			i--
			if i < 0 {
				i = 0
				break
			}
			var intersect uint32
			for j := 0; j < 4; j++ {
				if i+j < len(seq) {
					intersect += Shift(b[j], shift) & seq[i+j]
				}
			}
			if intersect != 0 {
				i++
				break
			}
			print2(seq, b, shift, i)
		}
		for j := 0; j < 4; j++ {
			if j+i >= len(seq) {
				seq = append(seq, Shift(b[j], shift))
			} else {
				seq[j+i] = seq[j+i] | Shift(b[j], shift)
			}
		}
		// fmt.Println("  876543210")
		// print(seq, 0)
		// fmt.Println("  876543210\n\n")
	}
	fmt.Println("  876543210")
	print(seq, 0)
	fmt.Println("  876543210\n\n")
	fmt.Println(seq, len(seq))
}

func print(seq []uint32, shift int) {
	for i := range seq {
		b := Shift(seq[len(seq)-1-i], shift)
		for c := 10; c >= 0; c-- {
			if b&(1<<c) != 0 {
				fmt.Print("#")
			} else if c == 8 || c == 0 {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println(b)
	}
}

func print2(seq []uint32, block [4]uint32, shift int, offset int) {
	return
	for i := 0; i < 100; i++ {
		fmt.Println("")
	}
	for i := -8; i < 28; i++ {
		k := len(seq) - 1 - i
		if k < 0 {
			continue
		}
		var b uint32
		if k < len(seq) {
			b = Shift(seq[k], 0)
		}
		for c := 10; c >= 0; c-- {
			p := "."
			if b&(1<<c) != 0 {
				p = "#"
			} else if c == 8 || c == 0 {
				p = "|"
			} else {
				p = "."
			}
			j := k - offset
			if j >= 0 && j < 4 {
				if Shift(block[j], shift)&(1<<c) != 0 {
					p = "@"
				}
			}
			fmt.Print(p)
		}
		fmt.Println(b)
	}
	time.Sleep(1000 * time.Millisecond)
}

func Shift(b uint32, c int) uint32 {
	if c >= 0 {
		return b << c
	}
	return b >> (-c)
}
