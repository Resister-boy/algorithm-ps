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
	for i := 0; i < (a / 4); i++ {
		fmt.Fprintf(writer, "long ")
	}
	fmt.Fprintf(writer, "int")
}
