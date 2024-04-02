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
    
    var t, a, b int
    
    fmt.Fscanln(reader, &t)
    for i := 1; i <= t; i++ {
        fmt.Fscanln(reader, &a, &b)
        fmt.Fprint(writer, "Case #")
        fmt.Fprint(writer, i)
        fmt.Fprint(writer, ": ")
        fmt.Fprint(writer, a)
        fmt.Fprint(writer, " + ")
        fmt.Fprint(writer, b)
        fmt.Fprint(writer, " = ")
        fmt.Fprintln(writer, a + b)
    }
}
