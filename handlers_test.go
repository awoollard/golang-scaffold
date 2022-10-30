package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test for the root endpoint
func TestRootHandler(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRoot)
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello World", rr.Body.String())
}
