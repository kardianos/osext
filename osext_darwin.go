package osext

import "errors"

// _NSGetExecutablePath()
func getExePath() (exePath string, err error) {
	return "", errors.New("Unimplemented")
}
