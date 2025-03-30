package reaper

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/rewgs/daws"
)

var platforms = []string{
	"darwin",
	"linux",
	"windows",
}

type Reaper struct {
	daws.Base
}

func getAll() (all []*Reaper) {
	getAllPaths := func() (paths []string) {
		walkDir := func(dir string) (paths []string) {
			err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() && strings.HasPrefix(d.Name(), "Reaper") {
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
			paths = walkDir("/Applications")
		case "linux":
			// TODO:
		case "windows":
			paths = walkDir("C:\\Program Files")
		default:
			log.Fatalf("%s not supported!", _os)
		}
		return
	}

	getVersion := func(name string) int {
		after, found := strings.CutPrefix(name, "Reaper")
		if !found {
			log.Fatalf("Not a Reaper path: %s", name)
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
		reaper := Reaper{}

		name := getName(filepath.Base(path))
		reaper.Name = name
		reaper.Path = path
		reaper.Version = getVersion(name)

		all = append(all, &reaper)
	}
	return
}

// func getLatest(all []*Reaper) *Reaper {
// }

// TODO: via CLI
// func getVersion(version int, all []*Reaper) (bool, *Reaper) {
// }

// func New() *Reaper {
// 	return getLatest(getAll())
// }

// func NewOfVersion(version int) *Reaper {
// 	exists, cubase := getVersion(version, getAll())
// 	if !exists {
// 		log.Fatalf("Cubase %d is not installed!", version)
// 	}
// 	return cubase
// }
