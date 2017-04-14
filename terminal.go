package terminal

import (
    "fmt"
)

const CLEAR = "\033[H\033[2J\033[3J"

func Clear() {
    fmt.Println(CLEAR)
    MoveCursor(0, 0)
}

func MoveCursor(x int, y int) {
    fmt.Printf("\033[%d;%dH", x, y)
}

