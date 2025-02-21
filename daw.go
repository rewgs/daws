package daws

// import "fmt"

type DAW interface {
	DefaultPreferencesPath() string
	IsOpen() bool
}
