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

	var s string
	var r int
	for i := 0; i < a; i++ {
		fmt.Fscanln(reader, &s)

		var pair = make(map[uint8]bool)
		var p uint8
		var isGroup = true

		for j := 0; j < len(s); j++ {
			_, exist := pair[s[j]]
			if !exist {
				pair[s[j]] = true
			} else if p != s[j] {
				isGroup = false
				break
			}
			p = s[j]
		}
		if isGroup {
			r++
		}
	}

	fmt.Println(writer, r)
}
