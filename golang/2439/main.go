package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var reader *bufio.Reader = bufio.NewReader(os.Stdin)
    var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

    defer writer.Flush()
    
    var a int
    
    fmt.Fscanln(reader, &a)
    for i := a; 0 < i; i-- {
        for j := 1; j <= a; j++ {
            if (j >= i) {
                fmt.Fprint(writer, "*")
            } else {
                fmt.Fprint(writer, " ")
            }
        }
        fmt.Fprintln(writer, "")
    }
}
