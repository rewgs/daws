package session

import (
	"log"
	"os"

	"github.com/rewgs/daws"
)

type Session struct {
	DAW      daws.DAW
	Path     string
	AudioDir string
}

func New(daw daws.DAW, path string) *Session {
	return &Session{
		DAW:  daw,
		Path: path,
	}
}

func (s *Session) Create() string {
	exists, _ := utils.PathExists(s.Path)
	if !exists {
		err := os.MkdirAll(s.Path, 755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// func (s *Session) GetSessionFiles() []string {
// }
