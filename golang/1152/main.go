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

	s, _ = reader.ReadString('\n')

	ss := strings.Split(s, " ")

	r := 0
	for i := range ss {
		if ss[i] != "\n" && ss[i] != "" {
			r++
		}
	}
	fmt.Fprintln(writer, r)
}
