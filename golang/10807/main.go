package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var t, v int

	fmt.Fscanln(reader, &t)

	s := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &s[i])
	}
	fmt.Fscan(reader, &v)
	fmt.Fprintln(writer, countEl(s, v))
}

func countEl(s []int, v int) int {
	c := 0

	for _, i := range s {
		if i == v {
			c += 1
		}
	}
	return c
}
