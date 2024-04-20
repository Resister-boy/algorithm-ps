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

	s := make([]int, 9)

	var x, y int
	sum := 0

	for i := 0; i < 9; i++ {
		fmt.Fscan(reader, &s[i])
		sum += s[i]
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 9; j++ {
			if sum-s[i]-s[j] == 100 {
				x = s[i]
				y = s[j]
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == x {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == y {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}

	sort.Ints(s)
	for _, v := range s {
		fmt.Fprintf(writer, "%d\n", v)
	}
}
