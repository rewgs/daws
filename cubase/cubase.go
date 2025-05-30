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

	"github.com/rewgs/daws/session"
)

const (
	developer string = "Steinberg"
	extension string = ".cpr"
)

var platforms = []string{
	"darwin",
	"windows",
}

type Cubase struct {
	Name    string
	Path    string
	Version int
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
func (c *Cubase) DefaultPreferencesPath() (path string) {
	switch _os := runtime.GOOS; _os {
	case "darwin":
		path = fmt.Sprintf("/Library/Preferences/%s", c.Name)
	case "linux":
		// TODO:
		// - WSL
		// - Wine
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

// TODO:
func (c *Cubase) IsOpen() bool {
	// ...
	return true
}

func (c *Cubase) NewSession(path string) *session.Session {
	s := session.New(c, path)
	// s.Create() // TODO:
	return s
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
