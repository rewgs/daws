package daws

import "fmt"

type Base struct {
	Name    string
	Path    string
	Version int
	prefs   *Prefs
}

func (daw *Base) GetName() string {
	return daw.Name
}

func (daw *Base) PrintName() {
	fmt.Println(daw.Name)
}
