package cubase

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	// "github.com/biter777/processex"

	"github.com/rewgs/daws"
)

var platforms = []string{
	"darwin",
	"windows",
}

// type Cubase struct {
// 	Name    string
// 	Path    string
// 	Version int
// }

type Cubase struct {
	daws.Base
}

func getAll() (all []*Cubase) {
	getAllPaths := func() (installed []string) {
		allCubasePaths := func(dir string) (paths []string) {
			err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() && strings.HasPrefix(d.Name(), "Cubase") {
					path, err := filepath.Abs(filepath.Join(dir, d.Name()))
					if err != nil {
						log.Fatal(err)
					}
					paths = append(paths, path)
					return fs.SkipDir
				}
				return nil
			})

			if err != nil {
				fmt.Printf("Error walking the path %q: %v\n", dir, err)
			}

			if len(paths) == 0 {
				log.Fatal("Cubase does not appear to be installed!")
			}

			return
		}

		switch _os := runtime.GOOS; _os {
		case "darwin":
			installed = allCubasePaths("/Applications")
		case "linux":
			// TODO: WSL only!
		case "windows":
			installed = allCubasePaths("C:\\Program Files")
		default:
			log.Fatalf("%s not supported!", _os)
		}
		return
	}

	getVersion := func(name string) int {
		after, found := strings.CutPrefix(name, "Cubase")
		if !found {
			log.Fatalf("Not a Cubase path: %s", name)
		}
		version, err := strconv.Atoi(strings.TrimSpace(after))
		if err != nil {
			log.Fatal(err)
		}
		return version
	}

	getName := func(name string) string {
		switch _os := runtime.GOOS; _os {
		case "darwin":
			return strings.TrimSuffix(name, filepath.Ext(name))
		default:
			return name
		}
	}

	for _, path := range getAllPaths() {
		cubase := Cubase{}

		name := getName(filepath.Base(path))
		cubase.Name = name
		cubase.Path = path
		cubase.Version = getVersion(name)

		all = append(all, &cubase)
	}
	return
}

func getLatest(all []*Cubase) *Cubase {
	var latestVersion int = 0
	var latestCubase *Cubase
	for _, c := range all {
		if c.Version > latestVersion {
			latestVersion = c.Version
			latestCubase = c
		}
	}
	return latestCubase
}

func getVersion(version int, all []*Cubase) (bool, *Cubase) {
	for _, c := range all {
		if c.Version == version {
			return true, c
		}
	}
	return false, nil
}

func New() *Cubase {
	return getLatest(getAll())
}

func NewOfVersion(version int) *Cubase {
	exists, cubase := getVersion(version, getAll())
	if !exists {
		log.Fatalf("Cubase %d is not installed!", version)
	}
	return cubase
}

// TODO:
// I beieve newer versions also store Preferences in ~/Documents. Account for that.
func (c *Cubase) DefaultPrefsPath() (path string) {
	switch _os := runtime.GOOS; _os {
	case "darwin":
		path = fmt.Sprintf("/Library/Preferences/%s", c.Name)
	case "linux":
		// TODO: WSL only!
	case "windows":
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(home, "AppData", "Roaming", "Steinberg", fmt.Sprintf("%s_64", c.Name))
	default:
		log.Fatalf("%s not supported!", _os)
	}
	return
}

// YNGNI?
// func getPathOfVersion(version int) string {
// 	for _, path := range getAllPaths() {
// 		after, _ := strings.CutPrefix(filepath.Base(path), "Cubase")
// 		if strconv.Itoa(version) == strings.TrimSpace(after) {
// 			return path
// 		}
// 	}
// 	return ""
// }

// Checks if this particular Cubase installation is open.
// old Python:
// return True if len([proc for proc in psutil.process_iter(["pid", "name", "username"]) if f"Cubase {self.version}" in proc.name() and proc.is_running()]) > 0 else False
// func (app Cubase) GetProcess() []*os.Process {
// 	basename := filepath.Base(app.Path)
// 	name := strings.ReplaceAll(basename, ".app", "")
// 	process, _, err := processex.FindByName(name)
// 	if err == processex.ErrNotFound {
// 		fmt.Printf("Process %v not running", name)
// 		os.Exit(0)
// 	}
// 	if err != nil {
// 		fmt.Printf("Process %v find error: %v", name, err)
// 		os.Exit(1)
// 	}
// 	return process
// }
