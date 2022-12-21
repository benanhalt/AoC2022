package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	prev *Node
	next *Node
	n    int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	input := make([]int, len(lines))
	for i, line := range lines {
		input[i], _ = strconv.Atoi(line)
	}
	part1(input)
	part2(input)
}

func part1(input []int) {
	n0, nodes := createLL(1, input)

	mix(nodes)

	fmt.Println("Part 1:", coords(n0, nodes))
}

func part2(input []int) {
	n0, nodes := createLL(811589153, input)

	for j := 0; j < 10; j++ {
		mix(nodes)
	}

	fmt.Println("Part 2:", coords(n0, nodes))
}

func createLL(key int, input []int) (*Node, []*Node) {
	N := len(input)
	nodes := make([]*Node, N)
	var n0 *Node
	for i, n := range input {
		nodes[i] = &Node{n: n * key}
		if n == 0 {
			n0 = nodes[i]
		}
	}
	for i, node := range nodes {
		node.prev = nodes[(i+N-1)%N]
		node.next = nodes[(i+1)%N]
	}
	return n0, nodes
}

func coords(n0 *Node, nodes []*Node) int {
	N := len(nodes)
	ans := 0
	i := 0
	for node := n0; ; {
		node = node.next
		i++
		if i == (1000%N) || i == (2000%N) || i == (3000%N) {
			ans += node.n
		}
		if node == n0 {
			break
		}
	}
	return ans
}

func mix(nodes []*Node) {
	N := len(nodes)
	for i := range nodes {
		node := nodes[i]

		if node.n != 0 {
			prev, next := nodes[i].prev, nodes[i].next
			prev.next = next
			next.prev = prev
		}

		if node.n < 0 {
			for i := node.n % (N - 1); i < 0; i++ {
				node = node.prev
			}
			nodes[i].prev = node.prev
			nodes[i].next = node
			node.prev.next = nodes[i]
			node.prev = nodes[i]
		} else if node.n > 0 {
			for i := node.n % (N - 1); i > 0; i-- {
				node = node.next
			}
			nodes[i].next = node.next
			nodes[i].prev = node
			node.next.prev = nodes[i]
			node.next = nodes[i]
		}
	}
}
