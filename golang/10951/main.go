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
    
    for true {
        v, _ := fmt.Fscanln(reader, &a, &b)
        if v != 2 {
            return ;
        }
        fmt.Fprintln(writer, a + b)
    }
}
