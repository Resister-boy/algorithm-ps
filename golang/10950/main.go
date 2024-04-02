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
    
    var s, a, b int

    fmt.Fscanln(reader, &s)
    for i := 0; i < s; i++ {
        fmt.Fscanln(reader, &a, &b)
        fmt.Fprintln(writer, a + b)
    }
}
