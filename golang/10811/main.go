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

	var a, b, f, t int

	fmt.Fscanln(reader, &a, &b)
	s := make([]int, a)

	for i := range s {
		s[i] = i + 1
	}

	for i := 0; i < b; i++ {
		fmt.Fscanln(reader, &f, &t)
		swap(&s, f-1, t-1)
	}

	for _, v := range s {
		fmt.Fprintf(writer, "%d ", v)
	}
}

func swap(s *[]int, f int, t int) {
	o := t
	for i := f; i <= f+(t-f)/2; i++ {
		a := (*s)[i]
		(*s)[i] = (*s)[o]
		(*s)[o] = a
		o--
	}
	fmt.Println(*s)
}
