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

	var a, b, c, d int

	fmt.Fscanln(reader, &a, &b, &c, &d)

	diff := a

	if diff > b {
		diff = b
	}
	if diff > c-a {
		diff = c - a
	}
	if diff > d-b {
		diff = d - b
	}
	fmt.Fprintln(writer, diff)
}
