package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type byPacketCmp []string

func (ss byPacketCmp) Len() int {
	return len(ss)
}

func (ss byPacketCmp) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss byPacketCmp) Less(i, j int) bool {
	var left, right interface{}
	json.Unmarshal([]byte(ss[i]), &left)
	json.Unmarshal([]byte(ss[j]), &right)
	return compare(left, right) < 0
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	ans := 0
	for i := 0; 3*i < len(lines); i++ {
		leftS := lines[3*i]
		rightS := lines[3*i+1]

		var left, right interface{}
		json.Unmarshal([]byte(leftS), &left)
		json.Unmarshal([]byte(rightS), &right)
		if compare(left, right) < 0 {
			ans += i + 1
		}
	}
	fmt.Println("Part 1:", ans)

	packets := []string{"[[2]]", "[[6]]"}
	for _, line := range lines {
		if line != "" {
			packets = append(packets, line)
		}
	}
	sort.Sort(byPacketCmp(packets))

	ans2 := 1
	for i, packet := range packets {
		if packet == "[[2]]" || packet == "[[6]]" {
			ans2 *= i + 1
		}
	}
	fmt.Println("Part 2:", ans2)
}

func compare(left, right interface{}) float64 {
	switch l := left.(type) {
	case float64:
		switch r := right.(type) {
		case float64:
			return l - r
		default:
			return compare([]interface{}{l}, right)
		}
	default:
		ll := l.([]interface{})
		switch r := right.(type) {
		case float64:
			return compare(l, []interface{}{r})
		default:
			rr := r.([]interface{})
			for i := 0; ; i++ {
				switch {
				case i == len(ll) && i == len(rr):
					return 0
				case i == len(ll):
					return -1
				case i == len(rr):
					return 1
				default:
					if c := compare(ll[i], rr[i]); c != 0 {
						return c
					}
				}
			}
		}
	}
}
