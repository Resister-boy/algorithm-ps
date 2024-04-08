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

	var a, k int

	fmt.Fscanln(reader, &a)

	for i := 0; i < a*2-1; i++ {
		if i < a {
			for j := 0; j < a+i; j++ {
				if j+i+(a-(a-i)) >= a+i-1 {
					fmt.Fprint(writer, "*")
				} else {
					fmt.Fprint(writer, " ")
				}
			}
			k = i
		} else {
			for j := 0; j <= k*2-(i-k); j++ {
				if j > i-a {
					fmt.Fprint(writer, "*")
				} else {
					fmt.Fprint(writer, " ")
				}
			}
		}
		fmt.Fprintln(writer, "")
	}
}
