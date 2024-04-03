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
	var s string

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &s)

	t := 0
	for _, v := range s {
		t += int(v) - 48
	}

	fmt.Fprintf(writer, "%d", t)
}
