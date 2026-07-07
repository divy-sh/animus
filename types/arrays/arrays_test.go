package arrays_test

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
	"github.com/divy-sh/animus/types/arrays"
)

func TestArGet(t *testing.T) {
	t.Run("returns the value at a valid index", func(t *testing.T) {
		key := "testArrayGetType"
		store.Set(key, []any{1, 2, 3, 4, 5})

		value, err := arrays.ArGet(key, 2)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if value != 3 {
			t.Fatalf("Expected value to be 3, got %v", value)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArGet("missingArray", 0)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an out-of-bounds index", func(t *testing.T) {
		key := "outOfBoundsArrayType"
		store.Set(key, []any{1, 2, 3})

		_, err := arrays.ArGet(key, 5)
		if err == nil {
			t.Fatalf("Expected an error for out-of-bounds index, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}
