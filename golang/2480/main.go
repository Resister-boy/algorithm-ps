package main

import (
	"bufio"
	"fmt"
	"os"
)

func biggest(a int, b int, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	}
	return c
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a, b, c int

	fmt.Fscanln(reader, &a, &b, &c)

	if a != b && a != c && b != c {
		fmt.Fprintln(writer, biggest(a, b, c)*100)
	} else if a == b && b == c {
		fmt.Fprintln(writer, 10000+a*1000)
	} else {
		if a == b {
			fmt.Fprintln(writer, 1000+a*100)
        } else if a == c {
			fmt.Fprintln(writer, 1000+a*100)
        } else {
			fmt.Fprintln(writer, 1000+b*100)
        }
	}
}

