package osext

// Returns the full path of the running executable
// as reported by the system. Includes the executable
// image name.
func GetExePath() (exePath string, err error) {
	return getExePath()
}
