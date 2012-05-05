package osext

import "testing"

func TestGetExePath(t *testing.T) {
	p, err := GetExePath()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("IMG EXE PATH: %s", p)
}
