package terminal

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
    "runtime"
)

type windowSize struct {
    Row     uint16
    Col     uint16
    Xpixel  uint16
    Ypixel  uint16
}


func GetWindowSize(*windowSize, error) {
    windowSize := new(windowSize)
    
    var _TIOCGWINSZ int64
    
    switch runtime.GOOS {
        case "linux":
            _TIOCGWINSZ = 0x5413
        case "darwin":
            _TIOCGWINSZ = 1074295912
    }
    
    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(_TIOCGWINSZ),
        uintptr(unsafe.Pointer(windowSize)),
    )
    
    if int(r1) == -1 {
        fmt.Println("Error: ", os.NewSyscallError("getWindowSize", errno))
        return nil, os.NewSyscallError("getWindowSize", errno)
    }
    
    return windowSize, nil
}