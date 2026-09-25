package sprint

import "testing"

func TestAbacus(t *testing.T) {
	if Abacus(8, 3) != 2 {
		t.Error("Abacus(8, 3) should return 2")
	}

	if Abacus(9, 2) != 4 {
		t.Error("Abacus(9, 2) should return 4")
	}
}