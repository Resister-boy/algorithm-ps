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

	a := make([]int, 6)
	d := []int{1, 1, 2, 2, 2, 8}

	for i := 0; i < len(a); i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i, v := range d {
		fmt.Fprintf(writer, "%d ", v-a[i])
	}
}
