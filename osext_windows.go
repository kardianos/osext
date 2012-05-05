package osext

import (
	"unsafe"
	"utf16"
)

var (
	kernel = syscall.MustLoadDLL("kernel32.dll")
	getModuleFileNameProc = kernel.MustFindProc("GetModuleFileNameW")
)

// GetModuleFileName() with hModule = NULL
func getExePath() (exePath string, err error) {
	return getModuleFileName()
}

func getModuleFileName() (string, error) {
        var n uint32
        b := make([]uint16, syscall.MAX_PATH)
        size := uint32(len(b))

        r0, _, e1 := getModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&b[0])), uintptr(size))
        n = uint32(r0)
        if n == 0 {
                return "", e1
        }
        return string(utf16.Decode(b[0:n])), nil
}
