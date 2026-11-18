package command

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
	"github.com/divy-sh/animus/store"
)

func TestBLMoveRegisteredAndMovesValue(t *testing.T) {
	sourceKey := "command-blmove-source"
	destinationKey := "command-blmove-destination"
	t.Cleanup(func() {
		store.Delete(sourceKey)
		store.Delete(destinationKey)
	})

	registered, ok := Handlers["BLMOVE"]
	if !ok {
		t.Fatal("BLMOVE is not registered")
	}
	if registered.Arity != 6 {
		t.Fatalf("BLMOVE registered arity = %d, want 6", registered.Arity)
	}

	RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: sourceKey},
		{Typ: common.BULK_TYPE, Bulk: "value"},
	})
	response := registered.Func([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: sourceKey},
		{Typ: common.BULK_TYPE, Bulk: destinationKey},
		{Typ: common.BULK_TYPE, Bulk: "LEFT"},
		{Typ: common.BULK_TYPE, Bulk: "RIGHT"},
		{Typ: common.BULK_TYPE, Bulk: "1"},
	})
	if response.Typ != common.BULK_TYPE || response.Bulk != "value" {
		t.Fatalf("BLMOVE response = %v, want bulk value", response)
	}
}

func TestBLMoveInvalidArgumentsAndTimeout(t *testing.T) {
	if response := BLMove(nil); response.Typ != common.ERROR_TYPE || response.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Fatalf("BLMOVE with missing arguments = %v, want wrong-argument-count error", response)
	}

	if response := BLMove([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "blmove-invalid-source"},
		{Typ: common.BULK_TYPE, Bulk: "blmove-invalid-destination"},
		{Typ: common.BULK_TYPE, Bulk: "LEFT"},
		{Typ: common.BULK_TYPE, Bulk: "RIGHT"},
		{Typ: common.BULK_TYPE, Bulk: "not-a-number"},
	}); response.Typ != common.ERROR_TYPE || response.Str != common.ERR_INVALID_TIME_SECONDS {
		t.Fatalf("BLMOVE with invalid timeout = %v, want invalid-time error", response)
	}

	response := BLMove([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "blmove-empty-source"},
		{Typ: common.BULK_TYPE, Bulk: "blmove-empty-destination"},
		{Typ: common.BULK_TYPE, Bulk: "LEFT"},
		{Typ: common.BULK_TYPE, Bulk: "RIGHT"},
		{Typ: common.BULK_TYPE, Bulk: "0.01"},
	})
	if response.Typ != common.NULL_TYPE {
		t.Fatalf("BLMOVE after timeout = %v, want null", response)
	}
}

func TestRPush(t *testing.T) {
	listKey := "mylist"
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: listKey},
		{Typ: common.BULK_TYPE, Bulk: "val1"},
		{Typ: common.BULK_TYPE, Bulk: "val2"}}
	result := RPush(args)

	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestRPushInvalidArgs(t *testing.T) {
	args := []resp.Value{}
	result := RPush(args)
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestRPopEmpty(t *testing.T) {
	listKey := "emptylist"
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: listKey}}
	result := RPop(args)

	expected := common.ERR_LIST_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestRPopSingle(t *testing.T) {
	listKey := "testlist"
	RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: listKey},
		{Typ: common.BULK_TYPE, Bulk: "val1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: listKey}}
	result := RPop(args)

	if result.Typ != "array" || len(result.Array) != 1 || result.Array[0].Bulk != "val1" {
		t.Errorf("Expected [val1], got %v", result)
	}
}

func TestRPopInvalidArgs(t *testing.T) {
	args := []resp.Value{}
	result := RPop(args)
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestRPopMultiple(t *testing.T) {
	listKey := "multilist"
	RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: listKey},
		{Typ: common.BULK_TYPE, Bulk: "val1"},
		{Typ: common.BULK_TYPE, Bulk: "val2"},
		{Typ: common.BULK_TYPE, Bulk: "val3"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: listKey}, {Typ: common.BULK_TYPE, Bulk: "2"}}
	result := RPop(args)

	if result.Typ != "array" ||
		len(result.Array) != 2 ||
		result.Array[0].Bulk != "val3" ||
		result.Array[1].Bulk != "val2" {
		t.Errorf("Expected [val3 val2], got %v", result)
	}
}
