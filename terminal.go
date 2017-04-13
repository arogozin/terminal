package terminal

import (
    "os"
    "bufio"
)

const RESET = "\033[2J"

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)

func Clear() {
    Output.WriteString(RESET)
}
