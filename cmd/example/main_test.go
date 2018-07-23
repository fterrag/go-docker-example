package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)
	main()
	log.SetOutput(os.Stderr)

	contains := "Hello John Smith!"
	out := buf.String()

	assert.Contains(t, out, contains)
}
