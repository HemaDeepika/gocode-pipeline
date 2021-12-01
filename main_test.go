package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!! How are you world??" {
		t.Fatal("Test fail")
	}
}
