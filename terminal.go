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

func Height() {
    windowSize, err := GetWindowSize()
    if err != nil {
        return -1
    }
    
    return int(windowSize.Row)
}

func Clear() {
    Output.WriteString(RESET)
}

func Flush() {
    for idx, str := range strings.Split(Screen.String(), "\n") {
        if idx > Height() {
            return
        }
        
        Output.WriteString(str + "\n")
    }
    Output.Flush()
    Screen.Reset()
}
