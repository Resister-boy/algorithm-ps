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
	m := -1
	t := 0.0
	fmt.Fscanln(reader, &a)

	s := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &s[i])
	}

	for _, v := range s {
		if m < v {
			m = v
		}
	}

	for i := 0; i < a; i++ {
		t += (float64(s[i]) / float64(m) * 100)
	}
	fmt.Fprintf(writer, "%f", t/float64(a))
}
