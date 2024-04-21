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

	var a, b, c int

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	if a+b+c != 180 {
		fmt.Fprintln(writer, "Error")
		return
	}

	if a == b && a == c {
		fmt.Fprintln(writer, "Equilateral")
		return
	}
	if a != b && a != c && b != c {
		fmt.Fprintln(writer, "Scalene")
	} else {
		fmt.Fprintln(writer, "Isosceles")
	}
}
