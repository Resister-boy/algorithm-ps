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

	var a, b, c int

	fmt.Fscanln(reader, &a, &b)
	fmt.Fscan(reader, &c)

	b = b + c

	if b >= 60 {
		a = a + b/60
		b = b % 60

		if a > 23 {
			a = a - 24

			fmt.Fprint(writer, a, b)
		} else {
			fmt.Fprint(writer, a, b)
		}
	} else {
		fmt.Fprint(writer, a, b)
	}
}
