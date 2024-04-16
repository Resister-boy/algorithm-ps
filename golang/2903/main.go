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

	fmt.Fscanln(reader, &a)

	r := 1
	for i := 1; i <= a; i++ {
		r *= 2
	}

	fmt.Fprintln(writer, (r+1)*(r+1))
}
