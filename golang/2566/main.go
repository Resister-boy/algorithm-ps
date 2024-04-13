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

	coor := 0
	max := 0
	var a int

	for i := 0; i < 81; i++ {
		fmt.Fscan(reader, &a)
		if i == 0 {
			max = a
			coor = 0
		} else {
			if a > max {
				max = a
				coor = i
			}
		}
	}

	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, coor/9+1, coor%9+1)
}
