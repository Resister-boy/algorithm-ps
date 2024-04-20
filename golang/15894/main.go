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

	s := 1
	t := 1

	fmt.Fscanln(reader, &a)

	if a != 1 {
		s = a * 2
		t = a - 1
	}

	fmt.Fprintln(writer, a+s+1+t)
}
