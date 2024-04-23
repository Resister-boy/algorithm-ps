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

	var a, b, tt int

	fmt.Fscanln(reader, &a, &b)

	s := make([]int, a)
	t := make([]int, 4)

	t[0] = -1

	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &s[i])
	}

	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			for k := 0; k < a; k++ {
				if i != j && i != k && j != k {
					tt = b - s[i] - s[j] - s[k]
					if tt >= 0 && (tt < t[0] || t[0] == -1) {
						t[0] = tt
						t[1] = s[i]
						t[2] = s[j]
						t[3] = s[k]
					}
				}
			}
		}
	}

	fmt.Fprintln(writer, t[1]+t[2]+t[3])
}
