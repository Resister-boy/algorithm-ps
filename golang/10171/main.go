package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fprintln(writer, "\\    /\\")
	fmt.Fprintln(writer, " )  ( ')")
	fmt.Fprintln(writer, "(  /  )")
	fmt.Fprintln(writer, " \\(__)|")
}
