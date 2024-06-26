package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a, b float64

	fmt.Fscanln(reader, &a, &b)

	fmt.Fprintln(writer, a/b)
}
