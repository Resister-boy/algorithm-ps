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
    if (a <= 100 && a >= 90) {
        fmt.Fprintln(writer, "A")
    } else if (a <= 89 && a >= 80) {
        fmt.Fprintln(writer, "B")
    } else if (a <= 79 && a >= 70) {
        fmt.Fprintln(writer, "C")
    } else if (a <= 69 && a >= 60) {
        fmt.Fprintln(writer, "D")
    } else {
        fmt.Fprintln(writer, "F")
    }
}
