package test

import "testing"

func TestIf(t *testing.T) {
	if n := 1 == 1; n {
		t.Log(n)
	}
}
