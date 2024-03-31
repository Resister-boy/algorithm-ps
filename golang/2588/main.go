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

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	k := b
	i := 1
	for (k / 10) != 0 {
		fmt.Fprintln(writer, a*(k%10))
		k /= 10
		i *= 10
	}

	fmt.Fprintln(writer, a*(b/i))
	fmt.Fprintln(writer, a*b)
}
