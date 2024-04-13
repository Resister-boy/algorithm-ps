package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var r string

	ss := make([][]string, 0, 5)
	sm := make(map[int]int)

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &r)
		sss := strings.Split(r, "")
		sm[i] = len(sss)
		ss = append(ss, sss)
	}

	for i := 0; i < 15; i++ {
		for j := 0; j < 5; j++ {
			if i < sm[j] {
				fmt.Fprint(writer, ss[j][i])
			}
		}
	}
}
