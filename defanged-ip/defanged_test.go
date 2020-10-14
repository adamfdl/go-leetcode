package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefangIP(t *testing.T) {
	defangedIP := defangIPaddr("1.1.1.1")
	assert.Equal(t, "1[.]1[.]1[.]1", defangedIP)
}
