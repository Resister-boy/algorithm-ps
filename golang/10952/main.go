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
    
    for true {
        fmt.Fscanln(reader, &a, &b)
        if (a == 0 && b == 0) {
            return 
        }
        fmt.Fprintln(writer, a + b)
    }
}
