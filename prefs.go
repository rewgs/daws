package daws

type Prefs struct {
	DefaultPath string
	UserPath    string
}

type PrefFile struct {
	Default bool
	Name    string
	Path    string
}
