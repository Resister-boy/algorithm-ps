package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a int
	var s string

	r := 0

	fmt.Fscanln(reader, &s, &a)

	for i := 0; i < len(s); i++ {
		t := s[len(s)-(i+1)]
		if t >= 48 && t <= 57 {
			t -= 48
		} else {
			t = t + 10 - 65
		}
		r += int(t) * int(math.Pow(float64(a), float64(i)))
	}

	fmt.Fprintln(writer, r)
}
