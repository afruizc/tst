package tst

import "testing"

// If err is not nil, calls t.Fatal(err)
func Ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// If expr is false, calls t.Error()
func True(t *testing.T, expr bool, fmtMsg string, vals ...interface{}) {
	t.Helper()
	if !expr {
		t.Errorf(fmtMsg, vals...)
	}
}
