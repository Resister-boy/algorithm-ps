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
    
    var a, b int 
    
    fmt.Fscanln(reader, &a, &b)
    
    if (a == 0 && b >= 45) {
        fmt.Fprintln(writer, 0, b - 45)
    } else if (a == 0 && b < 45) {
        fmt.Fprintln(writer, 23, 60 + (b - 45))
    } else if (a != 0 && b >= 45) {
        fmt.Fprintln(writer, a, b - 45)
    } else if (a != 0 && b < 45) {
        fmt.Fprintln(writer, a - 1, 60 + (b - 45))
    }
}
