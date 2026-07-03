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
