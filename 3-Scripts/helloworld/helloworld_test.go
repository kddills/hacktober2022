package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	name := "Tester"
	expectedString := fmt.Sprintf("Hello %s!", name)

	if HelloWorld("Tester") != expectedString {
		t.Fatalf("Hello world test failed 😢")
	}
}
