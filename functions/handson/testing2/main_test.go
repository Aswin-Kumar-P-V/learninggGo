package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(10, 5)
	if total != 15 {
		t.Errorf("Add failed")
	}
}
func TestSubtract(t *testing.T) {
	diff := subtract(10, 5)
	if diff != 5 {
		t.Errorf("subtraction failed")
	}
}
