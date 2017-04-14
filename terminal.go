package terminal

import (
    "fmt"
)

const RESET = "\033[2J"

func Clear() {
    fmt.Println(RESET)
}
