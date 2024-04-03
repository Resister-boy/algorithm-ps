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
    s := make([]int, 30)
    for i := range s {
        s[i] = 0
    }
    
    for i := 0; i < 28; i++ {
        fmt.Fscanln(reader, &a)
        s[a - 1] = 1
    }
    
    for i, v := range s {
        if v == 0 {
            fmt.Fprintln(writer, i + 1)
        }
    }
}
