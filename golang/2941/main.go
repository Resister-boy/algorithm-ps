package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var s string

	fmt.Fscanln(reader, &s)

	c := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

	for i, cv := range c {
		s = strings.Replace(s, cv, string(rune(i)), -1)
	}

	fmt.Fprintln(writer, len(s))
}
