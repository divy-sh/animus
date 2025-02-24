package types

import (
	"testing"
)

func TestHashType_HSetAndHGet(t *testing.T) {
	hashType := NewHashType()
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashType.HSet(hash, key, value)

	// Retrieve the value
	got, err := hashType.HGet(hash, key)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != value {
		t.Errorf("expected %s, got %s", value, got)
	}
}

func TestHashType_HGet_NotFound(t *testing.T) {
	hashType := NewHashType()

	// Try to get a non-existent key
	_, err := hashType.HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR not found" {
		t.Errorf("expected error 'ERR not found', got %v", err)
	}
}
