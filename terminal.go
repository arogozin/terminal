package terminal

import (
    "os"
    "bufio"
    "bytes"
    "strings"
    "fmt"
)

const RESET = "\033[2J"

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)


func Clear() {
    Output.WriteString(RESET)
}

func Flush() {
    for i, str := range strings.Split(Screen.String(), "\n") {
        if i > int(Height()) {
            return
        }
        
        fmt.Println(i)
        fmt.Println(str)
    }
    
    Output.Flush()
    Screen.Reset()
}