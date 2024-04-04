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

	var a, b, r int

	fmt.Fscanln(reader, &a, &b)

	aa := a
	bb := b

	for i := 10; aa != 0 && bb != 0; aa /= 10 {
		if aa%i > bb%i {
			r = a
			break
		} else if aa%i < bb%i {
			r = b
			break
		}
		bb /= 10
	}

	for i := 10; r != 0; r /= 10 {
		fmt.Fprint(writer, r%i)
	}
}
