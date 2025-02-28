package types_test

import (
	"testing"

	"github.com/divy-sh/animus/types"
)

func TestHashType_HSetAndHGet(t *testing.T) {
	hashType := types.NewHashType()
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
	hashType := types.NewHashType()

	// Try to get a non-existent hash
	_, err := hashType.HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR not found" {
		t.Errorf("expected error 'ERR not found', got %v", err)
	}
}

func TestHashType_HGet_KeyNotFound(t *testing.T) {
	hashType := types.NewHashType()
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashType.HSet(hash, key, value)
	// Try to get a non-existent key
	_, err := hashType.HGet("test_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR not found" {
		t.Errorf("expected error 'ERR not found', got %v", err)
	}
}

func TestHashType_HSet_KeyFound(t *testing.T) {
	hashType := types.NewHashType()

	hashType.HSet("test_hash", "test_key", "test_value")
	hashType.HSet("test_hash", "test_key", "new_value")

	val, err := hashType.HGet("test_hash", "test_key")
	if err != nil || val != "new_value" {
		t.Errorf("expected value 'new_value', got error %v", err)
	}
}
