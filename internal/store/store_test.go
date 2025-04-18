package store

import (
	"testing"
	"time"
)

func TestExpiryCleanerRemovesExpiredKeys(t *testing.T) {
	// Set key with short TTL (e.g., 200ms)
	SetWithTTL("test-key", "test-value", 1) // 1 second TTL

	// Make sure value is retrievable immediately
	val, ok := Get[string, string]("test-key")
	if !ok || val != "test-value" {
		t.Fatal("expected value to be present immediately after setting")
	}

	// Wait enough time for it to expire and cleaner to run
	time.Sleep(1 * time.Second) // Wait for TTL to expire + cleaner to run

	// Try to get the value again
	_, ok = Get[string, string]("test-key")
	if ok {
		t.Fatal("expected value to be expired and cleaned, but it was still present")
	}
}
