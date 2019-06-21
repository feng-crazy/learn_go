package main
import (
    "fmt"
    "bufio"
    "os"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {
    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please enter some input: ")
    input2, err = inputReader.ReadString('\n')
    if err == nil {
        fmt.Printf("The input was: %s\n", input2)
    }
}