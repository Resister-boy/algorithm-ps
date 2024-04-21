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
	s := make([][]int, 0)

	i := 0

	for {
		h := 0
		m := 0
		b := 0

		t := make([]int, 3)
		s = append(s, t)
		fmt.Fscan(reader, &s[i][0])
		h = s[i][0]
		fmt.Fscan(reader, &s[i][1])
		if s[i][1] > h {
			m = h
			h = s[i][1]
		} else {
			m = s[i][1]
		}
		fmt.Fscan(reader, &s[i][2])
		if s[i][2] > h {
			b = m
			m = h
			h = s[i][2]
		} else {
			if s[i][2] > m {
				b = m
				m = s[i][2]
			} else {
				b = s[i][2]
			}
		}

		if s[i][0] == 0 && s[i][1] == 0 && s[i][2] == 0 {
			s = s[:i]
			break
		}

		s[i][0] = h
		s[i][1] = m
		s[i][2] = b

		i++
	}

	for _, v := range s {

		if v[0] >= v[1]+v[2] {
			fmt.Fprintln(writer, "Invalid")
		} else {
			if v[0] == v[1] && v[1] == v[2] {
				fmt.Fprintln(writer, "Equilateral")
			} else {
				if v[0] != v[1] && v[0] != v[2] && v[1] != v[2] {
					fmt.Fprintln(writer, "Scalene")
				} else {
					fmt.Fprintln(writer, "Isosceles")
				}
			}

		}
	}
}
