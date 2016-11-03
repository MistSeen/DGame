package rand

import (
	. "server/core/rand"
	"testing"
)

func TestRandGroup(t *testing.T) {
	i := RandGroup(0, 0, 50, 50)
	switch i {
	case 2, 3:
		t.Log("okj")
	}

	// Output:
	// ok
}

func TestRandInterval(t *testing.T) {
	v := RandInterval(-1, 1)
	switch v {
	case -1, 0, 1:
		t.Log("okj")
	}

	// Output:
	// ok
}

func TestRandIntervalN(t *testing.T) {
	r := RandIntervalN(-1, 0, 2)
	if r[0] == -1 && r[1] == 0 ||
		r[0] == 0 && r[1] == -1 {
		t.Log("okj")
	}

	// Output:
	// ok
}
