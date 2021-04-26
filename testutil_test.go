package ondotori

import "testing"

func TestEquals(t *testing.T) {
	testEquals(t, "1123", "1123")
	testEquals(t, 17, 17)
}
