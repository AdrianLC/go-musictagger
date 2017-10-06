package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const mp3File = "testdata/Pat Benatar - Fire and Ice.mp3"

func TestGetGenreIntegration(t *testing.T) {
	genre := GetGenre(mp3File)
	assert.Contains(t, genre, "Rock")
	assert.Contains(t, genre, "Classic Rock")
}
