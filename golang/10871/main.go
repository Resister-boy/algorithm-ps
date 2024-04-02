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
    
    var a, b, t int
    
    fmt.Fscanln(reader, &a, &b)
    for i := 0; i < a; i++ {
        fmt.Fscan(reader, &t)
        if t < b {
            fmt.Fprint(writer, t, " ")
        }
    }
}
