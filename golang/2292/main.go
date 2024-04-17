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

	var a int
	m := 1
	p := 6
	r := 1

	fmt.Fscanln(reader, &a)

	for m < a {
		r++
		m += p
		p += 6
	}

	fmt.Fprintln(writer, r)
}
