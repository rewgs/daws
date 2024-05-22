package daws

import(
    "path/filepath"
)

type dawApp struct {
    path string
    version int
}

func (d dawApp) isOpen() {
}


type daw struct {
    app dawApp
    pref string
    name string
    developer string
    operatingSystems []string
}
