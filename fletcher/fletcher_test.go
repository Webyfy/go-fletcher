package fletcher

import "testing"

func TestFletcher16(t *testing.T) {
	data := []byte{1, 2, 3, 4}
	v := Fletcher16(data)
	if v != 5130 {
		t.Error("Expected 5130, got ", v)
	}
}
