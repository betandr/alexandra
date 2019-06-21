package main

import "testing"

func TestNothingPlaceholder(t *testing.T) {
	actualResult := "Placeholder"
	var expectedResult = "Placeholder"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
