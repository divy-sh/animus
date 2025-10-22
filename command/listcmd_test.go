package command

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

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
		result.Array[0].Bulk != "val2" ||
		result.Array[1].Bulk != "val3" {
		t.Errorf("Expected [val3 val2], got %v", result)
	}
}
