package commands

import (
	"fmt"
	"testing"

	"github.com/divy-sh/animus/resp"
)

func TestAppend(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: " world"}}
	result := append(args)
	if result.Typ != "string" || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestAppendInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: " world"}, {Typ: "bulk", Bulk: " world"}}
	result := append(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'append' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'append' command, got %v", result)
	}
}

func TestDecr(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "counter"}}
	result := decr(args)
	if result.Typ != "string" || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: " world"}}
	result := decr(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'decr' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'decr' command, got %v", result)
	}
}

func TestDecrInvalidValueType(t *testing.T) {
	set([]resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: "world"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}}
	result := decr(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR value is not an integer or out of range" {
		t.Errorf("Expected ERR value is not an integer or out of range, got %v", result)
	}
}

func TestDecrBy(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "counter"}, {Typ: "bulk", Bulk: "5"}}
	result := decrby(args)
	if result.Typ != "string" || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrByInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}}
	result := decrby(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'decrby' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'decrby' command, got %v", result)
	}
}

func TestDecrByInvalidValueType(t *testing.T) {
	set([]resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: "world"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "hello"}, {Typ: "bulk", Bulk: "hello"}}
	result := decrby(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR invalid decrement value" {
		t.Errorf("Expected error ERR invalid decrement value, got %v", result)
	}
}

func TestGet(t *testing.T) {
	set([]resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "value"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}}
	result := get(args)
	if result.Typ != "bulk" || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "non_existing"}}
	result := get(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "key"}}
	result := get(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'get' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'get' command, got %v", result)
	}
}

func TestGetDel(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}}
	result := getdel(args)
	if result.Typ != "bulk" && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetDelNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "non_existing"}}
	result := getdel(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetDelInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "key"}}
	result := getdel(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'getdel' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'getdel' command, got %v", result)
	}
}

func TestGetEx(t *testing.T) {
	set([]resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "value"}})
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "10"}}
	result := getex(args)
	if result.Typ != "bulk" {
		t.Errorf("Expected bulk, got %v", result)
	}
}

func TestGetExNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "non_existing"}}
	result := getex(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetExInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "key"}}
	result := getex(args)
	fmt.Println(result)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'getdel' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'getdel' command, got %v", result)
	}
}

func TestGetRange(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "0"}, {Typ: "bulk", Bulk: "4"}}
	result := getrange(args)
	if result.Typ != "bulk" && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetRangeInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "0"}}
	result := getrange(args)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'getrange' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'getrange' command, got %v", result)
	}
}

func TestGetRangeError(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "non_existing"}, {Typ: "bulk", Bulk: "0"}, {Typ: "bulk", Bulk: "4"}}
	result := getrange(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestSet(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}, {Typ: "bulk", Bulk: "value"}}
	result := set(args)
	if result.Typ != "string" || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestSetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: "bulk", Bulk: "key"}}
	result := set(args)
	if result.Typ != "error" || result.Str != "ERR wrong number of arguments for 'set' command" {
		t.Errorf("Expected ERR wrong number of arguments for 'set' command, got %v", result)
	}
}
