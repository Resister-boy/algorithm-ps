package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a, b int
	var s string

	fmt.Fscanln(reader, &a, &b)

	for i := a; i != 0; {
		t := i % b
		if t > 9 {
			t = t - 10 + 65
			s += string(rune(t))
		} else {
			t = t + 48
			s += string(rune(t))
		}
		i /= b
	}
	fmt.Fprintln(writer, reverse(s))
}
