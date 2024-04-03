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

	var a, d int
	var s string

	fmt.Fscanln(reader, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &d, &s)
		for k := 0; k < len(s); k++ {
			for j := 0; j < d; j++ {
				fmt.Fprintf(writer, "%c", s[k])
			}
		}
		fmt.Fprintln(writer, "")
	}
}
