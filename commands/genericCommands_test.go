package commands

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestStringCopy(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	copy([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}}
	result := get(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestHashCopy(t *testing.T) {
	hset([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	copy([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"}}
	result := hget(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestListCopy(t *testing.T) {
	rpush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	copy([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}}
	result := rpop(args)
	if result.Typ != "array" || result.Array[0].Bulk != "value2" {
		t.Errorf("Expected array or null, got %v", result)
	}
}

func TestGeneric_Copy_InvalidArgumentsCount(t *testing.T) {
	copy([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "invalidArguemntCount"}})
}

func TestGeneric_Copy_InvalidSourceKey(t *testing.T) {
	copy([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "InvalidSourceKey"},
		{Typ: common.BULK_TYPE, Bulk: "InvalidDestinationKey"},
	})
}

func TestStringDelete(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}}
	result := get(args)
	expected := common.ERR_STRING_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestHashDelete(t *testing.T) {
	hset([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}}
	result := hget(args)
	expected := common.ERR_HASH_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestListDelete(t *testing.T) {
	rpush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete2"}}
	result := rpop(args)
	expected := common.ERR_LIST_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestGeneric_Delete_InvalidArgumentsCount(t *testing.T) {
	del([]resp.Value{})
}
