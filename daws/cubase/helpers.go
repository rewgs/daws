package cubase

import (
	// "fs"
	"os"
	"path/filepath"
	"runtime"
)

func getSystemRoot() string {
	return os.Getenv("SystemDrive") + string(os.PathSeparator)
}

// func getFileSystem() fs.FS {
//     return os.DirFS(GetSystemRoot())
// }

// Returns the parent directory of the default Cubase installation location of
// the runtime OS.
func getDefaultPath() string {
	var defaultPath string
	if runtime.GOOS == "darwin" {
		systemRoot := getSystemRoot()
		defaultPath = filepath.Join(systemRoot, "Applications")
		// TODO:
		// } else if runtime.GOOS == "linux" {
		//     func isWsl() {
		//     }
		// TODO:
		// } else if runtime.GOOS == "windows" {
		// } else {
		//     return nil
	}
	return defaultPath
}
