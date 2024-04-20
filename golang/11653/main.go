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

	fmt.Fscanln(reader, &a)

	for i := 2; i <= a; i++ {
		for a%i == 0 {
			fmt.Fprintf(writer, "%d\n", i)
			a /= i
		}
		if a == 1 {
			break
		}
	}
}
