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
	fmt.Fscanln(reader, &n)

	arr := make([]int, 1000010)

	arr[1] = 0

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + 1
		if i%2 == 0 {
			arr[i] = min(arr[i], arr[i/2]+1)
		}
		if i%3 == 0 {
			arr[i] = min(arr[i], arr[i/3]+1)
		}
	}
	fmt.Fprintln(writer, arr[n])
}
