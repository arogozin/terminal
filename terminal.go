package terminal

import (
    "os"
    "bufio"
    "bytes"
)

const RESET = "\033[2J"

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)


func Clear() {
    Output.WriteString(RESET)
}
