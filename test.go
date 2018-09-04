package test

import "testing"

func assertEqual(t *testing.T, a interface{}, binterface{}) {
	if a != b {
		t.Fatalf("%s != %s", a,b)
	}
}
