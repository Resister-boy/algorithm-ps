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
    
    var a, t, max, min int
    
    fmt.Fscanln(reader, &a)
    for i := 0; i < a; i++ {
        fmt.Fscan(reader, &t)
        if (i == 0) {
            max = t
            min = t
        } else {
            if (t > max) {
                max = t
            } 
            if (t < min) {
                min = t
            }
        }
    }
    fmt.Fprintln(writer, min, max)
}
