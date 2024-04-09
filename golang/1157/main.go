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

	ss := strings.ToUpper(s)
	l := len(ss)
	ts := ""

	highest := ""
	t := 0

	for i := 0; i < l; i++ {
		c := 0
		if !strings.Contains(ts, string(ss[i])) {
			for j := i; j < l; j++ {
				if ss[i] == ss[j] {
					c++
				}
			}
		} else {
			ts = ts + string(ss[i])
		}
		if c > t {
			highest = string(ss[i])
			t = c
		} else if c == t {
			highest = "?"
		}
	}
	fmt.Fprintln(writer, highest);
}
