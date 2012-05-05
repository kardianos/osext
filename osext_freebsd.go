package osext

// FreeBSD: sysctl CTL_KERN KERN_PROC KERN_PROC_PATHNAME -1
// BSD with procfs: readlink /proc/curproc/file
func getExePath() (exePath string, err error) {
	return os.Readlink(`/proc/curproc/file`)
}
