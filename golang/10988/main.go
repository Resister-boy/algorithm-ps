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

	var s string

	fmt.Fscanln(reader, &s)
	l := len(s) - 1

	for i := 0; i < len(s); i++ {
		if s[i] != s[l-i] {
			fmt.Fprintln(writer, 0)
			return
		}
	}
	fmt.Fprintln(writer, 1)
}
