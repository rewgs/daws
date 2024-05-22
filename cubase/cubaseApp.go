package cubase

import (
    // "io/fs"
	"fmt"
    // "log"
	"os"
	"path/filepath"
    "strconv"
    "strings"

    "github.com/biter777/processex"
)

// A single Cubase application installation (i.e. one for Cubase 12, another 
// for Cubase 13, etc).
type CubaseApp struct {
    Path string
    Version int
}


func getVersion(s string) int {
    noCubase := strings.ReplaceAll(s, "Cubase", "")
    noApp := strings.ReplaceAll(noCubase, ".app", "")
    version, err := strconv.Atoi(strings.TrimSpace(noApp))
    if err != nil {
        panic(err)
    }
    return version
}


func GetAsCubaseApp(path string) CubaseApp {
    name := filepath.Base(path)
    version := getVersion(name)
    return CubaseApp{path, version}
}


// Checks if this particular Cubase installation is open. 
// old Python: 
// return True if len([proc for proc in psutil.process_iter(["pid", "name", "username"]) if f"Cubase {self.version}" in proc.name() and proc.is_running()]) > 0 else False
func (app CubaseApp) GetProcess() []*os.Process {
    basename := filepath.Base(app.Path)
    name := strings.ReplaceAll(basename, ".app", "")
    process, _, err := processex.FindByName(name)
	if err == processex.ErrNotFound {
		fmt.Printf("Process %v not running", name)
		os.Exit(0)
    }
	if err != nil {
		fmt.Printf("Process %v find error: %v", name, err)
		os.Exit(1)
	}
    return process
}
