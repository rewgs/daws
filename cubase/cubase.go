package cubase

import (
    "errors"
	"fmt"
    // "io/fs"
    "log"
	"os"
	"path/filepath"
    // "slices"
    "strings"
)

// Concerned with an abstract notion of Cubase -- the various application 
// versions installed, whether any of said applications are open, etc.
type Cubase struct {
    Apps []CubaseApp
    Latest CubaseApp
}


func checkIfVersionInstalled(cb Cubase, v int) (CubaseApp, bool) {
    installed := false
    var cubaseApp CubaseApp
    for i := range len(cb.Apps) {
    // for i := 0; i < len(cb.Apps); i++ {
        app := cb.Apps[i]
        if app.Version == v {
            installed = true
            cubaseApp = app
        }
    }
    return cubaseApp, installed
}


func (cb Cubase) GetByVersion(v int) (CubaseApp, error) {
    app, installed := checkIfVersionInstalled(cb, v)
    if installed {
        return app, nil
    } else {
        err := errors.New(fmt.Sprint("Cubase ", v, " is not installed!"))
        return app, err
    }
}


func (cb Cubase) GetNumApps() int {
    return len(cb.Apps)
}


func (cb Cubase) PrintNumApps() {
    fmt.Println(len(cb.Apps))
}


func (cb Cubase) PrintLatestApp() {
    fmt.Println(cb.Latest.Path)
}


func (cb Cubase) PrintApps() {
    for i := range len(cb.Apps) {
    // for i := 0; i < len(cb.Apps); i++ {
        app := cb.Apps[i]
        fmt.Println(app.Path)
    }
}


func Init() Cubase {
    apps := getCubaseApps()
    latest := getLatestApp(apps)
    return Cubase{apps, latest}
}


func getCubaseApps() []CubaseApp {
    var apps []CubaseApp
    installed := getInstalled()
    for i := 0; i < len(installed); i++ {
        app := GetAsCubaseApp(installed[i])
        apps = append(apps, app)
    }
    return apps
}


func getHighestVersion(apps []CubaseApp) int {
    var highest int
    for i := range len(apps) {
    // for i := 0; i < len(apps); i++ {
        app := apps[i]
        if app.Version > highest {
            highest = app.Version
        }
    }
    return highest
}


func getLatestApp(apps []CubaseApp) CubaseApp {
    var latestApp CubaseApp
    highest := getHighestVersion(apps)
    for i := range len(apps) {
    // for i := 0; i < len(apps); i++ {
        app := apps[i]
        if app.Version == highest {
            latestApp = app
        }
    }
    return latestApp
}


// Returns a slice containing all fs.DirEntry objects located at 
// GetDefaultPath() containing the word "Cubase".
func getInstalled() []string {
    defaultPath := GetDefaultPath()
    entries, err := os.ReadDir(defaultPath)
    if err != nil {
        log.Fatal(err)
    }
 
    var installed []string
    for _, entry := range entries {
        if strings.Contains(entry.Name(), "Cubase") {
            appPath := filepath.Join(defaultPath, entry.Name())
            if filepath.IsAbs(appPath) {
                installed = append(installed, appPath)
            }
        }
    }
    return installed
}


func PrintInstalled() {
    installed := getInstalled()
    for i := range len(installed) {
    // for i := 0; i < len(installed); i++ {
        fmt.Println(installed[i])
    }
}
