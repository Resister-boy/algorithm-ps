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

	r := 0

	fmt.Fscan(reader, &a)

	s := make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &s[i])
	}

	for _, v := range s {
		if v != 1 {
			for i := 1; i <= v/2; i++ {
				if i != 1 && v%i == 0 {
					break
				}

				if i == v/2 {
					r++
				}
			}
		}
	}

	fmt.Fprintln(writer, r)
}
