package main

import (
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, there!"

	wg.Add(2)
	go updateMessage("Hello, world!")
	go updateMessage("x")
	wg.Wait()

	if msg != "Hello, world!" {
		t.Error("incorrect value in msg")
	}

}
