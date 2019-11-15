package models

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestNewVersion(t *testing.T) {
	ver := NewVersion("1.0.0")
	assertEqual(t, ver, Version{1, 0, 0})
}
