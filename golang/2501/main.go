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

	var a, b int
	s := make([]int, 0)

	fmt.Fscanln(reader, &a, &b)

	for i := 1; i <= a/2; i++ {
		if a%i == 0 {
			s = append(s, i)
		}
	}
	s = append(s, a)
	if len(s) < b {
		fmt.Fprintln(writer, 0)
		return
	}
	fmt.Fprintln(writer, s[b-1])
}
