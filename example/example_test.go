package example

import "testing"

func TestStuff(t *testing.T) {
	if !*dbTests {
		t.Skipf("not running database tests")
	}

	t.Log("write some database tests...")
}
