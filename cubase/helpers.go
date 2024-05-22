package cubase

import (
    // "fs"
    "os"
	"path/filepath"
	"runtime"
)


func GetSystemRoot() string {
    return os.Getenv("SystemDrive") + string(os.PathSeparator)
}


// func GetFileSystem() fs.FS {
//     return os.DirFS(GetSystemRoot())
// }


// Returns the parent directory of the default Cubase installation location of 
// the runtime OS.
func GetDefaultPath() string {
	var defaultPath string
	if runtime.GOOS == "darwin" {
        systemRoot := GetSystemRoot()
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
