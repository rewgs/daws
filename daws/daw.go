package daws

import "fmt"

type DAW interface {
	DefaultPreferencesPath() string
	IsOpen() bool
}

type Base struct {
	Name    string
	Path    string
	Version int
}

func (daw *Base) GetName() string {
	return daw.Name
}

func (daw *Base) PrintName() {
	fmt.Println(daw.Name)
}
