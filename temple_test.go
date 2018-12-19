package temple

import (
	"testing"
)

func TestExecute(t *testing.T) {
	var result string
	result, _ = Execute("hello, {{.target}}!", map[string]interface{}{
		"target": "world",
	})
	if result != "hello, world!" {
		t.Fatal("failed test")
	}

	result, _ = Execute("{{.one}} {{.two}} {{.three}}", map[string]interface{}{
		"one": 1,
		"two": 2,
		"three": 3,
	})
	if result != "1 2 3" {
		t.Fatal("failed test")
	}
}