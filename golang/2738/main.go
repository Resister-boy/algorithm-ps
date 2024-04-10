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

	var x, y int

	fmt.Fscanln(reader, &x, &y)

	a := make([]int, x*y)
	b := make([]int, x*y)

	for i := 0; i < x*y; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < x*y; i++ {
		fmt.Fscan(reader, &b[i])
	}

	for i := 0; i < x*y; i++ {
		if i != 0 && (i+1)%x == 0 {
			fmt.Fprintf(writer, "%d\n", a[i]+b[i])
		} else {
			fmt.Fprintf(writer, "%d ", a[i]+b[i])
		}
	}
}
