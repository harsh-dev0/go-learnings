package basics

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	result := Hello() // Assuming Hello() is the function in hello.go that returns the greeting

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}