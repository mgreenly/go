package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello"
	actual := SayHello()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
