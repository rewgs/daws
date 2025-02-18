package cubase_test

import (
	"testing"

	"github.com/rewgs/daws/cubase"
)

func TestCubaseNew(t *testing.T) {
	cubase := cubase.New()
	want := "Cubase 14"
	got := cubase.GetName()
	if want != got {
		t.Errorf("Cubase.GetName(): Want: %s; Got: %s", want, got)
	}
}

func TestCubaseVersion(t *testing.T) {
	cubase := cubase.NewOfVersion(14)
	want := 14
	got := cubase.Version
	if want != got {
		t.Errorf("Cubase.NewOfVersion(): Want: %d; Got: %d", want, got)
	}
}
