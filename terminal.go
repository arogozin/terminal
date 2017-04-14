package terminal

import (
    "fmt"
    "bytes"
)

const CLEAR = "\033[2J"

var Screen *bytes.Buffer = new(bytes.Buffer)

func Clear() {
    fmt.Println(CLEAR)
    MoveCursor(0, 0)
}

func MoveCursor(x int, y int) {
    fmt.Printf("\033[%d;%dH", x, y)
}

