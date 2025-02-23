package world

import (
	"testing"
)

func TestSayWorld(t *testing.T) {
	expected := "world!"
	actual := SayWorld()

	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
