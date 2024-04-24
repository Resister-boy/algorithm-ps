package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a, b, ss int

	fmt.Fscan(reader, &a)

	s := make([]int, 0)

	for j := 1; j < a; j++ {
		ss = 0
		b = j
		for b > 0 {
			ss += b % 10
			b /= 10
		}
		if a == j+ss {
			s = append(s, j)
		}
	}

	if len(s) == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	sort.Ints(s)
	fmt.Fprintln(writer, s[0])
}
