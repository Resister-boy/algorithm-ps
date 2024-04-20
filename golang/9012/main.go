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

	var a int

	fmt.Fscanln(reader, &a)

	s := make([]string, a)

	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &s[i])
	}

	for _, v := range s {
		stat := 0

		for i := 0; i < len(v); i++ {
			if v[i] == '(' {
				stat += 1
			} else {
				if stat <= 0 {
					fmt.Fprintln(writer, "NO")
					break
				} else {
					stat -= 1
				}
			}

			if i == len(v)-1 {
				if stat == 0 {
					fmt.Fprintln(writer, "YES")
				} else {
					fmt.Fprintln(writer, "NO")
				}
			}
		}
	}
}
