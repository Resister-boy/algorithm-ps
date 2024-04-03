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
    
    var s string
    
    fmt.Fscanln(reader, &s)
    fmt.Fprintf(writer, "%d", len(s))
}
