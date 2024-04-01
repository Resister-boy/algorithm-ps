package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fprintln(writer, "|\\_/|")
	fmt.Fprintln(writer, "|q p|   /}")
	fmt.Fprintln(writer, "( 0 )\"\"\"\\")
	fmt.Fprintln(writer, "|\"^\"`    |")
	fmt.Fprintln(writer, "||_/=\\\\__|")
}
