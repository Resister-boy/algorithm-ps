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
    for i := 1; i <= a; i++ {
        for j := 0; j < i; j++ {
            fmt.Fprint(writer, "*")
        }
        fmt.Fprintln(writer, "")
    }
}
