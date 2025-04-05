package commands

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestAppend(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: " world"}}
	result := appendCmd(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestAppendInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: " world"}, {Typ: common.BULK_TYPE, Bulk: " world"}}
	result := appendCmd(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected %s, got %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestDecr(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}}
	result := decr(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: " world"}}
	result := decr(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'decr' command, got %v", result)
	}
}

func TestDecrInvalidValueEssentia(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := decr(args)
	if result.Typ != "error" || result.Str != "ERR value is not an integer or out of range" {
		t.Errorf("Expected ERR value is not an integer or out of range, got %v", result)
	}
}

func TestDecrBy(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}, {Typ: common.BULK_TYPE, Bulk: "5"}}
	result := decrby(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrByInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := decrby(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'decrby' command, got %v", result)
	}
}

func TestDecrByInvalidValueEssentia(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := decrby(args)
	if result.Typ != "error" || result.Str != "ERR invalid decrement value" {
		t.Errorf("Expected error ERR invalid decrement value, got %v", result)
	}
}

func TestGet(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := get(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}}
	result := get(args)
	expected := common.ERROR_STRING_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestGetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "key"}}
	result := get(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'get' command, got %v", result)
	}
}

func TestGetDel(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := getdel(args)
	if result.Typ != common.BULK_TYPE && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetDelNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}}
	result := getdel(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetDelInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "key"}}
	result := getdel(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'getdel' command, got %v", result)
	}
}

func TestGetEx(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "10"}}
	result := getex(args)
	if result.Typ != common.BULK_TYPE {
		t.Errorf("Expected bulk, got %v", result)
	}
}

func TestGetExNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}, {Typ: common.BULK_TYPE, Bulk: "key"}}
	result := getex(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetExInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := getex(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'getex' command, got %v", result)
	}
}

func TestGetRange(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "0"}, {Typ: common.BULK_TYPE, Bulk: "4"}}
	result := getrange(args)
	if result.Typ != common.BULK_TYPE && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetRangeInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "0"}}
	result := getrange(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'getrange' command, got %v", result)
	}
}

func TestGetRangeError(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}, {Typ: common.BULK_TYPE, Bulk: "0"}, {Typ: common.BULK_TYPE, Bulk: "4"}}
	result := getrange(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetSet(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "val"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "val1"}}
	result := getset(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "val" {
		t.Errorf("Expected val: val, got %v", result)
	}
	result = get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}})
	if result.Typ != common.BULK_TYPE || result.Bulk != "val1" {
		t.Errorf("Expected val: val1, got %v", result)
	}
}

func TestGetSetNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}, {Typ: common.BULK_TYPE, Bulk: "val"}}
	expected := common.ERROR_STRING_NOT_FOUND
	result := getset(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestGetSetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := getset(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestSet(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "value"}}
	result := set(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestSetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := set(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'set' command, got %v", result)
	}
}

func TestIncr(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}}
	result := incr(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: " world"}}
	result := incr(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'decr' command, got %v", result)
	}
}

func TestIncrInvalidValueEssentia(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := incr(args)
	if result.Typ != "error" || result.Str != "ERR value is not an integer or out of range" {
		t.Errorf("Expected ERR value is not an integer or out of range, got %v", result)
	}
}

func TestIncrBy(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}, {Typ: common.BULK_TYPE, Bulk: "5"}}
	result := incrby(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrByInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := incrby(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected ERR %s, got %v", expected, result)
	}
}

func TestIncrByInvalidValueEssentia(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := "ERR invalid increment value"
	result := incrby(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestIncrByFloat(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}, {Typ: common.BULK_TYPE, Bulk: "5"}}
	result := incrbyfloat(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrByInvalidArgumentCountFloat(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := incrbyfloat(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected ERR %s, got %v", expected, result)
	}
}

func TestIncrByInvalidValueEssentiaFloat(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}, {Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := "ERR invalid increment value"
	result := incrbyfloat(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestLcs(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}, {Typ: common.BULK_TYPE, Bulk: "lasagna"}})
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key2"}, {Typ: common.BULK_TYPE, Bulk: "baigan"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}, {Typ: common.BULK_TYPE, Bulk: "key2"}}
	result := lcs(args)
	if result.Typ == "error" || result.Bulk != "aga" {
		t.Errorf("Expected value: aga, got %v", result)
	}
}

func TestLcsLen(t *testing.T) {
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}, {Typ: common.BULK_TYPE, Bulk: "lasagna"}})
	set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key2"}, {Typ: common.BULK_TYPE, Bulk: "baigan"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}, {Typ: common.BULK_TYPE, Bulk: "key2"}, {Typ: common.BULK_TYPE, Bulk: "len"}}
	result := lcs(args)
	if result.Typ == "error" || result.Bulk != "3" {
		t.Errorf("Expected value: 3, got %v", result)
	}
}

func TestLcsInvalidFirstKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "invalid"}, {Typ: common.BULK_TYPE, Bulk: "key2"}}
	expected := common.ERROR_STRING_NOT_FOUND
	result := lcs(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestLcsInvalidSecondtKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}, {Typ: common.BULK_TYPE, Bulk: "invalid"}}
	expected := common.ERROR_STRING_NOT_FOUND
	result := lcs(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestLcsInvalidArguments(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := lcs(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestMGetAndMSet(t *testing.T) {
	mset([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "value2"},
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})

	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}, {Typ: common.BULK_TYPE, Bulk: "key2"}, {Typ: common.BULK_TYPE, Bulk: "invalid"}}
	expected := []string{"value", "value2", ""}
	result := mget(args)
	if result.Typ != "array" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
	for i, val := range result.Array {
		if val.Bulk != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], val.Bulk)
		}
	}
}
func TestMGetInvalidCommands(t *testing.T) {
	args := []resp.Value{}
	result := mget(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}

func TestMSetInvalidCommands(t *testing.T) {
	args := []resp.Value{}
	result := mset(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}

func TestMSetInvalidCommands2(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := mset(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}
