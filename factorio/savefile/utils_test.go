package savefile

import (
	"bytes"
	"testing"
)

func TestReadVersionHeader(t *testing.T) {
	v, err := ReadVersionHeader(bytes.NewBuffer([]byte{0x0, 0x0, 0x11, 0x0, 0x3, 0x0, 0x0, 0x0}))
	if err != nil {
		t.Fatalf("error reading test buffer: %s", err)
	}
	if v.Major != 0 {
		t.Errorf("wrong major value: %d != %d", v.Major, 0)
	}
	if v.Minor != 17 {
		t.Errorf("wrong minor value: %d != %d", v.Minor, 17)
	}
	if v.Patch != 3 {
		t.Errorf("wrong patch value: %d != %d", v.Patch, 3)
	}
	if v.Devel != 0 {
		t.Errorf("wrong devel value: %d != %d", v.Devel, 0)
	}
}
