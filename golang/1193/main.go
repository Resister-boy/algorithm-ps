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
	i := 1
	fmt.Fscanln(reader, &a)

	for i < a {
		a -= i
		i++
	}

	if i%2 == 1 {
		fmt.Fprintf(writer, "%d/%d\n", i+1-a, a)
	} else {
		fmt.Fprintf(writer, "%d/%d\n", a, i+1-a)
	}
}
