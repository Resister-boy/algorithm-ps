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

	s := make([]int, 3)

	defer writer.Flush()

	fmt.Fscanln(reader, &s[0], &s[1], &s[2])

	sort.Ints(s)

	if s[2] >= s[1]+s[0] {
		s[2] = s[1] + s[0] - 1
	}

	fmt.Fprintln(writer, s[0]+s[1]+s[2])
}
