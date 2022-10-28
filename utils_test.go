package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"testing/fstest"
)

func TestGetSha(t *testing.T) {
	// Arrange
	fs := fstest.MapFS{
		"git_hash": {
			// Include a newline at the end of the hash since this actually mirrors the behaviour of the docker command specified in the Dockerfile
			Data: []byte("7cb10c4122663100eb60d2ab3a256297df963a07\n"),
		},
	}
	data, err := fs.Open("git_hash")
	if err != nil {
		panic(err)
	}

	// Act
	output, err := GetSha(data)

	// Assert
	assert.Empty(t, err)
	assert.NotEmpty(t, output)
	assert.Equal(t, "7cb10c4122663100eb60d2ab3a256297df963a07", output)
}

func TestGetMetadata(t *testing.T) {
	// Arrange
	fs := fstest.MapFS{
		"meta.json": {
			Data: []byte("{\"description\": \"Test description\",\"version\": \"1.2\"}"),
		},
	}
	data, err := fs.Open("meta.json")
	if err != nil {
		panic(err)
	}

	// Act
	output, err := GetMetadata(data)

	// Assert
	assert.Empty(t, err)
	assert.NotEmpty(t, output)
	assert.NotEmpty(t, output.Description)
	assert.NotEmpty(t, output.Version)
	assert.Equal(t, "Test description", output.Description)
	assert.Equal(t, "1.2", output.Version)
}

func TestGetMetadataFailure(t *testing.T) {
	// Arrange
	fs := fstest.MapFS{
		"meta.json": {
			Data: []byte("{\"description\": \"Test description\",\"version\": \"1.2\"}"),
		},
	}
	data, err := fs.Open("meta.json")
	if err != nil {
		panic(err)
	}

	// Read all bytes from the file-handle, close it, then trigger the function which reads from it
	_, _ = io.ReadAll(data)
	err = data.Close()
	if err != nil {
		panic(err)
	}

	// Act
	output, err := GetMetadata(data)

	// Assert
	assert.NotEmpty(t, err)
	assert.Empty(t, output)
}
