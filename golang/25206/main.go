package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var subject string
	var score string
	var rate string

	a := 0.0
	h := 0.0
	m := map[string]float64{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0.0,
	}

	for i := 0; i < 20; i++ {
		fmt.Fscanln(reader, &subject, &score, &rate)

		q, _ := strconv.ParseFloat(score, 64)

		if rate != "P" {
			a += m[rate] * float64(q)
			h += float64(q)
		}
	}

	fmt.Fprintln(writer, a/h)
}
