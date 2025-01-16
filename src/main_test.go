package main

import "testing"

func TestAddition(t *testing.T) {
	result := 1 + 1
	expected := 2

	if result != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, result)
	}
}

func TestMainFunction(t *testing.T) {
	err := Run()
	if err != nil {
		t.Errorf("Failed to run main function: %v", err)
	}
}
