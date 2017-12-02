package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)
	main()
	log.SetOutput(os.Stderr)

	expected := "Hello John Smith!"
	actual := buf.String()

	if !strings.Contains(actual, expected) {
		t.Errorf("Output of main was incorrect, expected (contains): %s, actual: %s.", expected, actual)
	}
}
