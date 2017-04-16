package terminal

import (
    "fmt"
    "bufio"
    "os"
    "bytes"
)

const CLEAR = "\033[H\033[2J\033[3J"
const RESET = "\033[0m"

const BOLD = "\033[1m%s\033[0m"

var Output *bufio.Writer = bufio.NewWriter(os.Stdout)
var Screen *bytes.Buffer = new(bytes.Buffer)

func Clear() {
    fmt.Println(CLEAR)
    MoveCursor(0,0)
}

func MoveCursor(x int, y int) {
    fmt.Printf("\033[%d;%dH", x, y)
}

func Print(str string) {
    fmt.Println(Output, str)
    Output.Flush()
}

func Bold(str string) {
    fmt.Printf(BOLD, str+"\n")
}
