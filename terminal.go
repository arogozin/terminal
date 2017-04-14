package terminal

import (
    "os"
    "bufio"
    "bytes"
    "strings"
)

const RESET = "\033[2J"

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)


func Clear() {
    Output.WriteString(RESET)
}

func Flush() {
    for i, str := range strings.Split(Screen.String(), "\n") {
        if i > Height() {
            return
        }
        
        Output.WriteString(str + "\n")
    }
    
    Output.Flush()
    Screen.Reset()
}