package bloom

import (
	"errors"
	"testing"
)

// TestNewRejectsZeroFPRate verifies that New returns ErrInvalidFPRate
// when called with fpRate=0, since log(0) is -Inf and would cause
// the bit array size to overflow.
func TestNewRejectsZeroFPRate(t *testing.T) {
	_, err := New(1000, 0)
	if err == nil {
		t.Fatal("expected error for fpRate=0, got nil")
	}
	if !errors.Is(err, ErrInvalidFPRate) {
		t.Fatalf("expected ErrInvalidFPRate, got: %v", err)
	}
}
