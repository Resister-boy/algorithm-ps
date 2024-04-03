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

	var s string
	var t bool

	fmt.Fscanln(reader, &s)

	for i := 97; i <= 122; i++ {
		t = false
		for k, v := range s {
			if int(v) == i {
				fmt.Fprintf(writer, "%d ", k)
				t = true
				break
			}
		}
		if !t {
			fmt.Fprintf(writer, "%d ", -1)
		}
	}
}
