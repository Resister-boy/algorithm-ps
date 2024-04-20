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

	m := -1
	s := 0

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	for i := a; i <= b; i++ {
		for j := 1; j <= i/2; j++ {
			if j != 1 && i%j == 0 {
				break
			}

			if j+1 > i/2 {
				s += i
				if m == -1 || i < m {
					m = i
				}
			}
		}
	}

	if m == -1 && s == 0 {
		fmt.Fprintf(writer, "%d\n", -1)
		return
	}
	fmt.Fprintf(writer, "%d\n%d\n", s, m)
}
