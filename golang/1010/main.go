package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp [][]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	dp = make([][]int, 31)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 31)
	}
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)
		fmt.Fprintln(writer, combi(m, n))
	}
}

func combi(m, n int) int {
	if m == n || n == 0 {
		return 1
	}
	if dp[m][n] != 0 {
		return dp[m][n]
	}
	dp[m][n] = combi(m-1, n) + combi(m-1, n-1) // 조합의 성질, 파스칼의 삼각형 활용.
	return dp[m][n]
}
