package hashes

import (
	"testing"

	"github.com/divy-sh/animus/internal/common"
)

func TestHashEssentia_HSetAndHGet(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	HSet(hash, key, value)

	// Retrieve the value
	got, err := HGet(hash, key)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != value {
		t.Errorf("expected %s, got %s", value, got)
	}
}

func TestHashEssentia_HGet_NotFound(t *testing.T) {

	// Try to get a non-existent hash
	_, err := HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HGet_KeyNotFound(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	HSet(hash, key, value)
	// Try to get a non-existent key
	_, err := HGet("test_hash", "non_existent_key")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HSet_KeyFound(t *testing.T) {
	HSet("test_hash", "test_key", "test_value")
	HSet("test_hash", "test_key", "new_value")

	val, err := HGet("test_hash", "test_key")
	if err != nil || val != "new_value" {
		t.Errorf("expected value 'new_value', got error %v", err)
	}
}

func Test_HashExists(t *testing.T) {
	HSet("Tesh_HashExists", "Tesh_HashExists", "test_value")
	exists, err := HExists("Tesh_HashExists", "Tesh_HashExists")
	if exists != 1 || err != nil {
		t.Errorf("Expected hash to exist, got: %d, %v", exists, err)
	}
}

func Test_HashExists_Nope(t *testing.T) {
	exists, err := HExists("Test_HashExists_Nope", "Test_HashExists_Nope")
	if exists == 1 || err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected hash to not exist, got: %d, %v", exists, err)
	}
}
