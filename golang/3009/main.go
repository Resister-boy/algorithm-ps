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

	s := make([][]int, 3)

	for i := 0; i < 3; i++ {
		s[i] = make([]int, 2)
		fmt.Fscanln(reader, &s[i][0], &s[i][1])
	}

	var a, b int

	if s[0][0] == s[1][0] {
		a = s[2][0]
	} else if s[0][0] == s[2][0] {
		a = s[1][0]
	} else {
		a = s[0][0]
	}

	if s[0][1] == s[1][1] {
		b = s[2][1]
	} else if s[0][1] == s[2][1] {
		b = s[1][1]
	} else {
		b = s[0][1]
	}

	fmt.Fprintf(writer, "%d %d\n", a, b)
}
