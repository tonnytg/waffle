package handler

import (
	"bytes"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello World"
	result := jsonError(msg)
	expected := []byte(`{"message":"Hello World"}`)

	if !bytes.Equal(result, expected) {
		t.Error("result must been equal to expected")
	}
}
