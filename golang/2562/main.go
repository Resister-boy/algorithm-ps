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
    s := make([]int, 2, 2)
    for i := 1; i <= 9; i++ {
        fmt.Fscanln(reader, &a)
        if (i == 0) {
            s[0] = a
            s[1] = i
        } else {
            if s[0] < a {
                s[0] = a
                s[1] = i
            }
        }
    }
    fmt.Fprintln(writer, s[0], s[1])
}
