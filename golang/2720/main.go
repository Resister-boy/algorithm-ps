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

	fmt.Fscanln(reader, &a)

	s := make([]map[int]int, a)

	for i := 0; i < a; i++ {
		s[i] = map[int]int{}
	}

	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &b)
		for b != 0 {
			if b-25 >= 0 {
				b -= 25
				s[i][25] += 1
			} else if b-10 >= 0 {
				b -= 10
				s[i][10] += 1
			} else if b-5 >= 0 {
				b -= 5
				s[i][5] += 1
			} else if b-1 >= 0 {
				b -= 1
				s[i][1] += 1
			}
		}
	}
	for _, v := range s {
		fmt.Fprintf(writer, "%d %d %d %d\n", v[25], v[10], v[5], v[1])
	}
}
