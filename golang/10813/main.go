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

	var a, b, c int

	fmt.Fscanln(reader, &a, &b)
	s := make([]int, a)
	t := make([]int, 2)

	for i := range s {
		s[i] = i + 1
	}

	for i := 0; i < b; i++ {
		fmt.Fscan(reader, &t[0], &t[1])
		c = s[t[0]-1]
		s[t[0]-1] = s[t[1]-1]
		s[t[1]-1] = c
	}

	for _, v := range s {
		fmt.Fprint(writer, v, " ")
	}
}
