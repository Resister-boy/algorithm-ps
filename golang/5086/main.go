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

	for {
		fmt.Fscanln(reader, &a, &b)
		if a == 0 && b == 0 {
			return
		}

		if b%a == 0 {
			fmt.Fprintln(writer, "factor")
		} else if a%b == 0 {
			fmt.Fprintln(writer, "multiple")
		} else {
			fmt.Fprintln(writer, "neither")
		}
	}
}
