package osext

import "os"

// readlink /proc/self/exe
func getExePath() (exePath string, err error) {
	return os.Readlink(`/proc/self/exe`)
}
