package command_test

import (
	"reflect"
	"testing"

	"github.com/divy-sh/animus/command"
	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/store"
)

func TestArCount(t *testing.T) {
	// Test case 1: Array exists
	key := "testArray"
	store.Set(key, []any{1, 2, 3, 4, 5})

	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
	}

	result := command.ArCount(args)
	if result.Typ == common.ERROR_TYPE {
		t.Fatalf("Expected no error, got %v", result.Str)
	}
	if result.Num != 5 {
		t.Fatalf("Expected count to be 5, got %d", result.Num)
	}

	// Test case 2: Array does not exist
	nonExistentKey := "nonExistentArray"
	args = []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: nonExistentKey},
	}

	result = command.ArCount(args)
	if result.Typ != common.ERROR_TYPE {
		t.Fatalf("Expected an error for non-existent array, got nil")
	}
	if result.Num != 0 {
		t.Fatalf("Expected count to be 0 for non-existent array, got %d", result.Num)
	}
}

func TestArCount_EmptyArray(t *testing.T) {
	// Test case 3: Empty array
	key := "emptyArray"
	store.Set(key, []any{})

	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
	}

	result := command.ArCount(args)
	if result.Typ == common.ERROR_TYPE {
		t.Fatalf("Expected no error for empty array, got %v", result.Str)
	}
	if result.Num != 0 {
		t.Fatalf("Expected count to be 0 for empty array, got %d", result.Num)
	}
}

func TestArCount_WrongArgumentCount(t *testing.T) {
	// Test case 4: Wrong argument count
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "someKey"},
		{Typ: common.BULK_TYPE, Bulk: "extraArg"},
	}

	result := command.ArCount(args)
	if result.Typ != common.ERROR_TYPE {
		t.Fatalf("Expected an error for wrong argument count, got nil")
	}
	if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Fatalf("Expected error message '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}

func TestArDel(t *testing.T) {
	// Test case 1: Array exists and index is valid
	key := "testArray"
	store.Set(key, []any{1, 2, 3, 4, 5})

	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
		{Typ: common.INTEGER_TYPE, Num: 2},
	}

	result := command.ArDel(args)
	if result.Typ == common.ERROR_TYPE {
		t.Fatalf("Expected no error, got %v", result.Str)
	}
	if result.Bulk != "OK" {
		t.Fatalf("Expected response to be 'OK', got '%s'", result.Bulk)
	}

	// Verify that the element at index 2 was deleted
	arr, ok := store.Get[string, []any](key)
	if !ok {
		t.Fatalf("Expected array to exist after deletion")
	}
	if len(arr) != 4 || arr[2] != 4 {
		t.Fatalf("Expected array to be [1, 2, 4, 5], got %v", arr)
	}

	// Test case 2: Array does not exist
	nonExistentKey := "nonExistentArray"
	args = []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: nonExistentKey},
		{Typ: common.INTEGER_TYPE, Num: 0},
	}

	result = command.ArDel(args)
	if result.Typ != common.ERROR_TYPE {
		t.Fatalf("Expected an error for non-existent array, got nil")
	}
	if result.Str != common.ERR_ARRAY_NOT_FOUND {
		t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, result.Str)
	}

	// Test case 3: Index out of bounds
	args = []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
		{Typ: common.INTEGER_TYPE, Num: 10},
	}

	result = command.ArDel(args)
	if result.Typ != common.ERROR_TYPE {
		t.Fatalf("Expected an error for index out of bounds, got nil")
	}
	if result.Str != common.ERR_INDEX_OUT_OF_BOUNDS {
		t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, result.Str)
	}
}

func TestArDel_WrongArgumentCount(t *testing.T) {
	// Test case 4: Wrong argument count
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "someKey"},
	}

	result := command.ArDel(args)
	if result.Typ != common.ERROR_TYPE {
		t.Fatalf("Expected an error for wrong argument count, got nil")
	}
	if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Fatalf("Expected error message '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}
func TestArGetRange(t *testing.T) {
	t.Run("returns the values in the requested range", func(t *testing.T) {
		key := "testArrayGetRange"
		store.Set(key, []any{1, 2, 3, 4, 5})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 1},
			{Typ: common.INTEGER_TYPE, Num: 3},
		}

		result := command.ArGetRange(args)
		if result.Typ == common.ERROR_TYPE {
			t.Fatalf("Expected no error, got %v", result.Str)
		}
		if result.Typ != common.ARRAY_TYPE {
			t.Fatalf("Expected result type ARRAY_TYPE, got %v", result.Typ)
		}
		if len(result.Array) != 3 {
			t.Fatalf("Expected 3 values, got %d", len(result.Array))
		}
		got := []string{result.Array[0].Bulk, result.Array[1].Bulk, result.Array[2].Bulk}
		expected := []string{"2", "3", "4"}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("Expected values %v, got %v", expected, got)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: "missingArray"},
			{Typ: common.INTEGER_TYPE, Num: 0},
			{Typ: common.INTEGER_TYPE, Num: 1},
		}

		result := command.ArGetRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if result.Str != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, result.Str)
		}
	})

	t.Run("returns an error for out-of-bounds range", func(t *testing.T) {
		key := "outOfBoundsRangeArray"
		store.Set(key, []any{1, 2, 3})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 1},
			{Typ: common.INTEGER_TYPE, Num: 5},
		}

		result := command.ArGetRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for out-of-bounds range, got nil")
		}
		if result.Str != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, result.Str)
		}
	})

	t.Run("returns an error for wrong argument count", func(t *testing.T) {
		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: "someKey"},
			{Typ: common.INTEGER_TYPE, Num: 0},
		}

		result := command.ArGetRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for wrong argument count, got nil")
		}
		if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
		}
	})
}
func TestArDelRange(t *testing.T) {
	t.Run("deletes a range successfully", func(t *testing.T) {
		key := "testArrayRange"
		store.Set(key, []any{1, 2, 3, 4, 5})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 1},
			{Typ: common.INTEGER_TYPE, Num: 3},
		}

		result := command.ArDelRange(args)
		if result.Typ == common.ERROR_TYPE {
			t.Fatalf("Expected no error, got %v", result.Str)
		}
		if result.Bulk != "OK" {
			t.Fatalf("Expected response to be 'OK', got '%s'", result.Bulk)
		}

		arr, ok := store.Get[string, []any](key)
		if !ok {
			t.Fatalf("Expected array to exist after deletion")
		}
		if !reflect.DeepEqual(arr, []any{1, 5}) {
			t.Fatalf("Expected array to be [1 5], got %v", arr)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: "missingArray"},
			{Typ: common.INTEGER_TYPE, Num: 0},
			{Typ: common.INTEGER_TYPE, Num: 1},
		}

		result := command.ArDelRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if result.Str != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, result.Str)
		}
	})

	t.Run("returns an error when start is greater than end", func(t *testing.T) {
		key := "invalidRangeArray"
		store.Set(key, []any{1, 2, 3})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 2},
			{Typ: common.INTEGER_TYPE, Num: 1},
		}

		result := command.ArDelRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for invalid range, got nil")
		}
		if result.Str != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, result.Str)
		}
	})

	t.Run("returns an error when end is out of bounds", func(t *testing.T) {
		key := "outOfBoundsArray"
		store.Set(key, []any{1, 2, 3})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 1},
			{Typ: common.INTEGER_TYPE, Num: 5},
		}

		result := command.ArDelRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for out-of-bounds range, got nil")
		}
		if result.Str != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, result.Str)
		}
	})

	t.Run("returns an error for wrong argument count", func(t *testing.T) {
		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: "someKey"},
			{Typ: common.INTEGER_TYPE, Num: 0},
		}

		result := command.ArDelRange(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for wrong argument count, got nil")
		}
		if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
		}
	})
}

func TestArGet(t *testing.T) {
	t.Run("returns the value at a valid index", func(t *testing.T) {
		key := "testArrayGet"
		store.Set(key, []any{1, 2, 3, 4, 5})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 2},
		}

		result := command.ArGet(args)
		if result.Typ == common.ERROR_TYPE {
			t.Fatalf("Expected no error, got %v", result.Str)
		}
		if result.Bulk != "3" {
			t.Fatalf("Expected value to be '3', got '%s'", result.Bulk)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: "missingArray"},
			{Typ: common.INTEGER_TYPE, Num: 0},
		}

		result := command.ArGet(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if result.Str != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, result.Str)
		}
	})

	t.Run("returns an error for an out-of-bounds index", func(t *testing.T) {
		key := "outOfBoundsArray"
		store.Set(key, []any{1, 2, 3})

		args := []resp.Value{
			{Typ: common.BULK_TYPE, Bulk: key},
			{Typ: common.INTEGER_TYPE, Num: 5},
		}

		result := command.ArGet(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for out-of-bounds index, got nil")
		}
		if result.Str != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, result.Str)
		}
	})

	t.Run("returns an error for wrong argument count", func(t *testing.T) {
		args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "someKey"}}

		result := command.ArGet(args)
		if result.Typ != common.ERROR_TYPE {
			t.Fatalf("Expected an error for wrong argument count, got nil")
		}
		if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
		}
	})
}
