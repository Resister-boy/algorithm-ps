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

	var a, b, c int

	fmt.Fscanln(reader, &a)

	tx := 0
	ty := 0
	bx := 0
	by := 0

	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &b, &c)

		if tx == 0 || tx < b {
			tx = b
		}
		if bx == 0 || bx > b {
			bx = b
		}
		if ty == 0 || ty < c {
			ty = c
		}
		if by == 0 || by > c {
			by = c
		}
	}

	fmt.Fprintln(writer, (tx-bx)*(ty-by))
}
