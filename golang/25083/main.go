package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fprintf(w, "         ,r'\"7\n")
	fmt.Fprintf(w, "r`-_   ,'  ,/\n")
	fmt.Fprintf(w, " \\. \". L_r'\n")
	fmt.Fprintf(w, "   `~\\/\n")
	fmt.Fprintf(w, "      |\n")
	fmt.Fprintf(w, "      |\n")
}
