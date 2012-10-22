package osext

/*
#include <mach-o/dyld.h>
#include <string.h>

int
GetExecPath(char* path) {
	uint32_t size = 32*1024;
	if (_NSGetExecutablePath(path, &size) == 0) {
		// Despite Apple docs, size does NOT get set in call.
		return strlen(path);
	} else {
		return 0;
	}
}
*/
import "C"
import (
	"unsafe"
	"errors"
	"path/filepath"
)

// _NSGetExecutablePath()
func getExePath() (exePath string, err error) {
	buffer := make([]byte, 32*1024)
	size := C.GetExecPath((*C.char)(unsafe.Pointer(&buffer[0])))
	if size == 0 {
		return "", errors.New("Unable to get exec path.")
	}
	buffer = buffer[:size]
	ret := string(buffer)
	return filepath.Clean(ret), nil
}
