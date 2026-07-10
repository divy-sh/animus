package arrays_test

import (
	"reflect"
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

func TestArCount(t *testing.T) {
	t.Run("returns the length of an existing array", func(t *testing.T) {
		key := "testArrayCountType"
		store.Set(key, []any{1, 2, 3, 4})

		count, err := arrays.ArCount(key)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if count != 4 {
			t.Fatalf("Expected count to be 4, got %d", count)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArCount("missingCountArray")
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})
}

func TestArDel(t *testing.T) {
	t.Run("removes the value at a valid index", func(t *testing.T) {
		key := "testArrayDelType"
		store.Set(key, []any{1, 2, 3, 4})

		err := arrays.ArDel(key, 1)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		updated, ok := store.Get[string, []any](key)
		if !ok {
			t.Fatalf("Expected array to remain stored")
		}
		if !reflect.DeepEqual(updated, []any{1, 3, 4}) {
			t.Fatalf("Expected array to become [1 3 4], got %v", updated)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		err := arrays.ArDel("missingDeleteArray", 0)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an out-of-bounds index", func(t *testing.T) {
		key := "testArrayDelOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		err := arrays.ArDel(key, 5)
		if err == nil {
			t.Fatalf("Expected an error for out-of-bounds index, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArDelRange(t *testing.T) {
	t.Run("removes the values in a valid range", func(t *testing.T) {
		key := "testArrayDelRangeType"
		store.Set(key, []any{10, 20, 30, 40, 50})

		err := arrays.ArDelRange(key, 1, 3)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		updated, ok := store.Get[string, []any](key)
		if !ok {
			t.Fatalf("Expected array to remain stored")
		}
		if !reflect.DeepEqual(updated, []any{10, 50}) {
			t.Fatalf("Expected array to become [10 50], got %v", updated)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		err := arrays.ArDelRange("missingDeleteRangeArray", 0, 1)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an invalid range", func(t *testing.T) {
		key := "testArrayDelRangeOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		err := arrays.ArDelRange(key, 1, 5)
		if err == nil {
			t.Fatalf("Expected an error for invalid range, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArGetRange(t *testing.T) {
	t.Run("returns the values in a valid range", func(t *testing.T) {
		key := "testArrayGetRangeType"
		store.Set(key, []any{1, 2, 3, 4, 5})

		values, err := arrays.ArGetRange(key, 1, 3)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{2, 3, 4}) {
			t.Fatalf("Expected values [2 3 4], got %v", values)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArGetRange("missingArrayRange", 0, 1)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an invalid range", func(t *testing.T) {
		key := "testArrayGetRangeOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		_, err := arrays.ArGetRange(key, 1, 5)
		if err == nil {
			t.Fatalf("Expected an error for invalid range, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}
