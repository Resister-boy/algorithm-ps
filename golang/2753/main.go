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
    if ((a % 4 == 0) && ((a % 100 != 0) || (a % 400 == 0))) {
        fmt.Fprintln(writer, "1")
        return ;
    } 
    fmt.Fprintln(writer, "0")
}
