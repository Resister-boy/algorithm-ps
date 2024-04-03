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

	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &s)
		l := len(s)
		fmt.Fprintf(writer, "%c", s[0])
		fmt.Fprintf(writer, "%c", s[l-1])
		fmt.Fprintln(writer, "")
	}
}
