package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ostojics/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	mockRequest, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(responseRecorder, mockRequest)
	result := responseRecorder.Result()

	assert.Equal(t, result.StatusCode, http.StatusOK)

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
