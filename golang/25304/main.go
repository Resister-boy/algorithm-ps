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
    
    var t, c, p, a int
    
    fmt.Fscanln(reader, &t)
    fmt.Fscanln(reader, &c)
    for i := 0; i < c; i++ {
        fmt.Fscanln(reader, &p, &a)
        t -= (p * a)
    }
    if (t == 0) {
        fmt.Fprintln(writer, "Yes")
        return 
    }
    fmt.Fprintln(writer, "No")
}
