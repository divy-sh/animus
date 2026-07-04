package command_test

import (
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
