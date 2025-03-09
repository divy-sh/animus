package essentias_test

import (
	"testing"

	"github.com/divy-sh/animus/essentias"
)

func TestHashEssentia_HSetAndHGet(t *testing.T) {
	hashEssentia := essentias.NewHashEssentia()
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashEssentia.HSet(hash, key, value)

	// Retrieve the value
	got, err := hashEssentia.HGet(hash, key)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != value {
		t.Errorf("expected %s, got %s", value, got)
	}
}

func TestHashEssentia_HGet_NotFound(t *testing.T) {
	hashEssentia := essentias.NewHashEssentia()

	// Try to get a non-existent hash
	_, err := hashEssentia.HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR not found" {
		t.Errorf("expected error 'ERR not found', got %v", err)
	}
}

func TestHashEssentia_HGet_KeyNotFound(t *testing.T) {
	hashEssentia := essentias.NewHashEssentia()
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashEssentia.HSet(hash, key, value)
	// Try to get a non-existent key
	_, err := hashEssentia.HGet("test_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR not found" {
		t.Errorf("expected error 'ERR not found', got %v", err)
	}
}

func TestHashEssentia_HSet_KeyFound(t *testing.T) {
	hashEssentia := essentias.NewHashEssentia()

	hashEssentia.HSet("test_hash", "test_key", "test_value")
	hashEssentia.HSet("test_hash", "test_key", "new_value")

	val, err := hashEssentia.HGet("test_hash", "test_key")
	if err != nil || val != "new_value" {
		t.Errorf("expected value 'new_value', got error %v", err)
	}
}
