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

	var n int
	var t [101][101]int
	var r int = 0

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				t[x+j][y+k] = 1
			}
		}
	}

	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			r += t[i][j]
		}
	}

	fmt.Fprintln(writer, r)
}
