package commands

import (
	"testing"

	"github.com/divy-sh/animus/resp"
)

func TestStringCopy(t *testing.T) {
	set([]resp.Value{{Typ: "bulk", Bulk: "TestStringCopy1"}, {Typ: "bulk", Bulk: "value"}})
	copy([]resp.Value{{Typ: "bulk", Bulk: "TestStringCopy1"}, {Typ: "bulk", Bulk: "TestStringCopy2"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "TestStringCopy2"}}
	result := get(args)
	if result.Typ != "bulk" || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestHashCopy(t *testing.T) {
	hset([]resp.Value{{Typ: "bulk", Bulk: "TestHashCopy1"}, {Typ: "bulk", Bulk: "TestHashCopy1"}, {Typ: "bulk", Bulk: "value"}})
	copy([]resp.Value{{Typ: "bulk", Bulk: "TestHashCopy1"}, {Typ: "bulk", Bulk: "TestHashCopy2"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "TestHashCopy2"}, {Typ: "bulk", Bulk: "TestHashCopy1"}}
	result := hget(args)
	if result.Typ != "bulk" || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestListCopy(t *testing.T) {
	rpush([]resp.Value{{Typ: "bulk", Bulk: "TestListCopy1"}, {Typ: "bulk", Bulk: "value1"}, {Typ: "bulk", Bulk: "value2"}})
	copy([]resp.Value{{Typ: "bulk", Bulk: "TestListCopy1"}, {Typ: "bulk", Bulk: "TestListCopy2"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "TestListCopy2"}}
	result := rpop(args)
	if result.Typ != "array" || result.Array[0].Bulk != "value2" {
		t.Errorf("Expected array or null, got %v", result)
	}
}
