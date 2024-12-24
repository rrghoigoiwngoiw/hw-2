package pool

import (
	"testing"
)

func TestPool(t *testing.T) {
	result := Pool()

	want := 5
	if result != want {
		t.Errorf("Expected counter value to be %d, got %d", want, result)
	}
}
