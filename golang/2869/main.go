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

	var a, b, v int

	r := 1
	fmt.Fscanln(reader, &a, &b, &v)

	r += (v - a) / (a - b)
	if ((v - a) % (a - b)) != 0 {
		r++
	}

	if a >= v {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, r)
	}
}
