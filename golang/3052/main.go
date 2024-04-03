package main

import (
	"bufio"
	"fmt"
	"os"
)

func isContain(s []int, v int, a int) bool {
	for i := 0; i < a; i++ {
		if s[i] == v {
			return true
		}
	}
	return false
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	t := 0
	s := make([]int, 10)
    var a int
    
	for i := range s {
		s[i] = -1
	}

	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &a)
        s[i] = a % 42
		if !isContain(s, s[i], i) {
			t += 1
		}
	}
	fmt.Fprintln(writer, t)
}

