package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var s string
	a := 0
	fmt.Fscanln(reader, &s)
	d := strings.Split("ABC,DEF,GHI,JKL,MNO,PQRS,TUV,WXYZ", ",")

	for _, v := range s {
		for i, w := range d {
			if strings.Contains(w, string(v)) {
				a += i + 3
			}
		}
	}

	fmt.Fprintln(writer, a)
}
