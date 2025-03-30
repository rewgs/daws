package daws

import (
	"errors"
	"fmt"
	"log"
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

func pathExists(path string) (bool, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	return true, nil
}

// func findFileInDir(dir string, query string) string {
// 	var matches = []string{}
// 	err := filepath.WalkDir(dir, func(path string, file fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !file.IsDir() && file.Name() == query {
// 			abs, err := filepath.Abs(filepath.Join(path, file.Name()))
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			matches = append(matches, abs)
// 			return fs.SkipDir
// 		}
// 		return nil
// 	})
//
// 	if err != nil {
// 		fmt.Printf("Error walking the path %q: %v\n", dir, err)
// 	}
//
// 	if len(matches) == 0 {
// 		log.Fatalf("Could not find %s in %s", query, dir)
// 	}
//
// 	if len(matches) > 1 {
// 		fmt.Printf("Found more than one match for %s in %s:\n", query, dir)
// 		for _, match := range matches {
// 			fmt.Println(match)
// 		}
// 		os.Exit(1)
// 	}
//
// 	return matches[0]
// }

func findFileInDir(dir string, file string) string {
	path := filepath.Join(dir, file)
	matches, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	if len(matches) == 0 {
		log.Fatalf("Could not find %s in %s", file, dir)
	}

	if len(matches) > 1 {
		fmt.Printf("Found more than one match for %s in %s:\n", file, dir)
		for _, match := range matches {
			fmt.Println(match)
		}
		os.Exit(1)
	}

	fmt.Printf("Found %s in %s", file, dir)
	return matches[0]
}
