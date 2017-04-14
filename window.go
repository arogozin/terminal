package terminal

import (
    "syscall"
    "unsafe"
)

type size struct {
    Row     uint16
    Col     uint16
    Xpixel  uint16
    Ypixel  uint16
}


func Size() size {
    size := &size{}
    
    returnCode, _, error := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(size)),
    )
    
    if int(returnCode) == -1 {
        panic(error)
    }
    
    return *size
}

func Height() uint {
    size := Size()
    
    return uint(size.Col)
}

func Width() uint {
    size := Size()
    
    return uint(size.Row)
}
