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
	aa := 0
	fmt.Fscanln(reader, &a)

	for a >= 0 {
		if a%5 == 0 {
			aa += a / 5
			fmt.Fprintln(writer, aa)
			return
		}
		a -= 3
		aa++
	}
	fmt.Fprintln(writer, -1)
}
