package essentias_test

import (
	"testing"

	"github.com/divy-sh/animus/essentias"
)

func TestHashEssentia_HSetAndHGet(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	essentias.HSet(hash, key, value)

	// Retrieve the value
	got, err := essentias.HGet(hash, key)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != value {
		t.Errorf("expected %s, got %s", value, got)
	}
}

func TestHashEssentia_HGet_NotFound(t *testing.T) {

	// Try to get a non-existent hash
	_, err := essentias.HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR hash does not exist" {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HGet_KeyNotFound(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	essentias.HSet(hash, key, value)
	// Try to get a non-existent key
	_, err := essentias.HGet("test_hash", "non_existent_key")
	if err == nil || err.Error() != "ERR hash does not exist" {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HSet_KeyFound(t *testing.T) {

	essentias.HSet("test_hash", "test_key", "test_value")
	essentias.HSet("test_hash", "test_key", "new_value")

	val, err := essentias.HGet("test_hash", "test_key")
	if err != nil || val != "new_value" {
		t.Errorf("expected value 'new_value', got error %v", err)
	}
}
