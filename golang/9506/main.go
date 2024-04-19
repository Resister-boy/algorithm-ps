package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSum(a []int) int {
	sum := 0

	for i, v := range a {
		if len(a)-1 != i {
			sum += v
		}
	}
	return sum
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var a int
	m := make([][]int, 0)

	i := 0
	for {
		fmt.Fscanln(reader, &a)
		if a == -1 {
			break
		}
		t := make([]int, 0)
		m = append(m, t)
		for j := 1; j <= a/2; j++ {
			if a%j == 0 {
				m[i] = append(m[i], j)
			}
		}
		m[i] = append(m[i], a)
		i++
	}

	for _, v := range m {
		if v[len(v)-1] == getSum(v) {
			fmt.Fprintf(writer, "%d = ", v[len(v)-1])
			for i := 0; i < len(v)-2; i++ {
				fmt.Fprintf(writer, "%d + ", v[i])
			}
			fmt.Fprintf(writer, "%d\n", v[len(v)-2])
		} else {
			fmt.Fprintf(writer, "%d is NOT perfect.\n", v[len(v)-1])
		}
	}
}
