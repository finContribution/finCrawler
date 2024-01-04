package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	test := assert.New(t)
	LoadEnv()
	testObject := os.Getenv("TOKEN")
	test.NotEqual(0, len(testObject))
}

func TestNewToken(t *testing.T) {
	test := assert.New(t)
	testObject := NewToken()
	test.NotEqual(0, len(testObject))
}
