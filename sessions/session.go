package session

type Session struct {
	Path     string
	AudioDir string
}

func (s *Session) GetSessionFiles() []string {
}
