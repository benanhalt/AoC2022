package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	ans := 0
	for _, line := range lines {
		f := 0
		for _, ch := range line {
			f *= 5
			switch ch {
			case '-':
				f += -1
			case '=':
				f += -2
			case '1':
				f += 1
			case '2':
				f += 2
			case '0':
			default:
				panic(ch)
			}
		}
		ans += f
	}
	fmt.Println(ans)

	result := ""
	for {
		switch ans % 5 {
		case 0:
			result = "0" + result
		case 1:
			result = "1" + result
			ans -= 1
		case 2:
			result = "2" + result
			ans -= 2
		case 3:
			result = "=" + result
			ans += 5
			ans -= 3
		case 4:
			result = "-" + result
			ans += 5
			ans -= 4
		}
		if ans%5 != 0 {
			panic(ans)
		}
		ans /= 5
		if ans == 0 {
			break
		}
	}
	fmt.Println(result)
}
