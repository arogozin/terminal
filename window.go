package terminal

import (
    "syscall"
    "unsafe"
)

type windowSize struct {
    Row     uint16
    Col     uint16
    Xpixel  uint16
    Ypixel  uint16
}


func GetWindowSize() uint {
    windowSize := &windowSize{}
    
    returnCode, _, error := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(windowSize)),
    )
    
    if int(returnCode) == -1 {
        panic(error)
    }
    
    return uint(windowSize.Col)
}