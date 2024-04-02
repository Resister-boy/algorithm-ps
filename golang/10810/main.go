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

	var a, b int
	fmt.Fscanln(reader, &a, &b)
	s := make([]int, a)
	t := make([]int, 3)

	for i := range s {
		s[i] = -1
	}

	for i := 0; i < b; i++ {
		fmt.Fscan(reader, &t[0], &t[1], &t[2])
		fmt.Println(t)
		for j := t[0]; j <= t[1]; j++ {
			s[j-1] = t[2]
		}
	}

	fmt.Println(s)

	for _, v := range s {
		if v == -1 {
			fmt.Fprint(writer, 0, " ")
		} else {
			fmt.Fprint(writer, v, " ")
		}
	}
}
