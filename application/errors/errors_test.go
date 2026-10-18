package errors

import (
	"errors"
	"testing"
)

func TestDetailsSupportsWrappedApplicationErrors(t *testing.T) {
	sentinel := New(CodeNotFound, 404, "resource not found")
	wrapped := errors.Join(errors.New("operation failed"), sentinel)

	code, status, message, ok := Details(wrapped)
	if !ok {
		t.Fatal("Details() did not identify the application error")
	}
	if code != CodeNotFound || status != 404 || message != "resource not found" {
		t.Fatalf("Details() = %q, %d, %q; want not_found, 404, resource not found", code, status, message)
	}
}
