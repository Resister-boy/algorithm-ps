package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	key   int
	left  *node
	right *node
}

func addNode(n *node, key int) *node {
	if n == nil {
		n = &node{key: key, left: nil, right: nil}
	} else if key < n.key {
		n.left = addNode(n.left, key)
	} else {
		n.right = addNode(n.right, key)
	}
	return n
}

func printNode(n *node) {
	if n == nil {
		return
	}
	printNode(n.left)
	printNode(n.right)
	fmt.Println(n.key)
}

func main() {
	var n *node = nil
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		a := scanner.Text()

		if len(a) == 0 || a == "" || a == "\n" {
			break
		}
		aa, _ := strconv.Atoi(a)
		n = addNode(n, aa)
	}

	printNode(n)
}
