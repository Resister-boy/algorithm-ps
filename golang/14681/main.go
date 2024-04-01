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
    
    var a, b int
    
    fmt.Fscanln(reader, &a)
    fmt.Fscanln(reader, &b)
    
    if (a > 0 && b > 0) {
        fmt.Fprintln(writer, "1")
    } else if (a < 0 && b > 0) {
        fmt.Fprintln(writer, "2")
    } else if (a < 0 && b < 0) {
        fmt.Fprintln(writer, "3")
    } else if (a > 0 && b < 0) {
        fmt.Fprintln(writer, "4")
    }
}
