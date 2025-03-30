package daws

type DAW interface {
	DefaultPreferencesPath() (path string)
	IsOpen() bool
}
