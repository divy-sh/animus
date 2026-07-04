package arrays_test

import (
	"testing"
	"time"

	"github.com/divy-sh/animus/store"
	"github.com/divy-sh/animus/types/arrays"
)

func TestArCount(t *testing.T) {
	// Test case 1: Array exists
	key := "testArray"
	store.Set(key, []any{1, 2, 3, 4, 5})

	count, err := arrays.ArCount(key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if count != 5 {
		t.Fatalf("Expected count to be 5, got %d", count)
	}

	// Test case 2: Array does not exist
	nonExistentKey := "nonExistentArray"
	count, err = arrays.ArCount(nonExistentKey)
	if err == nil {
		t.Fatalf("Expected an error for non-existent array, got nil")
	}
	if count != 0 {
		t.Fatalf("Expected count to be 0 for non-existent array, got %d", count)
	}
}

func TestArCount_EmptyArray(t *testing.T) {
	// Test case 3: Empty array
	key := "emptyArray"
	store.Set(key, []any{})

	count, err := arrays.ArCount(key)
	if err != nil {
		t.Fatalf("Expected no error for empty array, got %v", err)
	}
	if count != 0 {
		t.Fatalf("Expected count to be 0 for empty array, got %d", count)
	}
}

func TestArCount_InvalidType(t *testing.T) {
	// Test case 5: Key exists but is not an array
	key := "notAnArray"
	store.Set(key, "this is a string")

	count, err := arrays.ArCount(key)
	if err == nil {
		t.Fatalf("Expected an error for key that is not an array, got nil")
	}
	if count != 0 {
		t.Fatalf("Expected count to be 0 for key that is not an array, got %d", count)
	}
}

func TestArCount_ExpiredKey(t *testing.T) {
	// Test case 6: Key exists but is expired
	key := "expiredArray"
	store.SetWithTTL(key, []any{1, 2, 3}, 1) // Set expiration to 1 second

	// Wait for the key to expire
	time.Sleep(2 * time.Second)

	count, err := arrays.ArCount(key)
	if err == nil {
		t.Fatalf("Expected an error for expired key, got nil")
	}
	if count != 0 {
		t.Fatalf("Expected count to be 0 for expired key, got %d", count)
	}
}

func TestArDel(t *testing.T) {
	// Test case 1: Array exists and index is valid
	key := "testArrayDel"
	store.Set(key, []any{1, 2, 3, 4, 5})

	err := arrays.ArDel(key, 2) // Delete the element at index 2 (value 3)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	count, _ := arrays.ArCount(key)
	if count != 4 {
		t.Fatalf("Expected count to be 4 after deletion, got %d", count)
	}

	// Test case 2: Array does not exist
	nonExistentKey := "nonExistentArrayDel"
	err = arrays.ArDel(nonExistentKey, 0)
	if err == nil {
		t.Fatalf("Expected an error for non-existent array, got nil")
	}

	// Test case 3: Index out of bounds
	err = arrays.ArDel(key, 10) // Index 10 is out of bounds for an array of size 4
	if err == nil {
		t.Fatalf("Expected an error for index out of bounds, got nil")
	}
}

func TestArDel_EmptyArray(t *testing.T) {
	// Test case 4: Empty array
	key := "emptyArrayDel"
	store.Set(key, []any{})

	err := arrays.ArDel(key, 0) // Attempt to delete from an empty array
	if err == nil {
		t.Fatalf("Expected an error for deleting from an empty array, got nil")
	}
}

func TestArDel_InvalidType(t *testing.T) {
	// Test case 5: Key exists but is not an array
	key := "notAnArrayDel"
	store.Set(key, "this is a string")

	err := arrays.ArDel(key, 0)
	if err == nil {
		t.Fatalf("Expected an error for key that is not an array, got nil")
	}
}

func TestArDel_ExpiredKey(t *testing.T) {
	// Test case 6: Key exists but is expired
	key := "expiredArrayDel"
	store.SetWithTTL(key, []any{1, 2, 3}, 1) // Set expiration to 1 second

	// Wait for the key to expire
	time.Sleep(2 * time.Second)

	err := arrays.ArDel(key, 0)
	if err == nil {
		t.Fatalf("Expected an error for expired key, got nil")
	}
}
