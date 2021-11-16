package util

import (
	"os"
	"testing"
)

func TestConsole(t *testing.T) {
	args := []string{os.Args[0], "do", "-name=lucy"}
	console := NewConsole(args)

	if !console.HasCommand("do") {
		t.Error("command ?")
	}

	if console.HasCommand("do1") {
		t.Error("why do1 ?")
	}

	if !console.HasFlag("name") {
		t.Error("flag ?")
	}

	if console.HasFlag("name1") {
		t.Error("why name1 ??")
	}

	if console.GetFlag("name") != "lucy" {
		t.Error("name != lucy ??")
	}

	if console.GetFlagAuto("age", "20") != "20" {
		t.Error("age != 20 ??")
	}
}
