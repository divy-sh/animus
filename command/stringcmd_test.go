package command

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestAppend(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: " world"}}
	result := Append(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestAppendInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: " world"},
		{Typ: common.BULK_TYPE, Bulk: " world"}}
	result := Append(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected %s, got %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestDecr(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}}
	result := Decr(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: " world"}}
	result := Decr(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'Decr' command, got %v", result)
	}
}

func TestDecrInvalidValueEssentia(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := Decr(args)
	if result.Typ != "error" || result.Str != "ERR value is not an integer or out of range" {
		t.Errorf("Expected ERR value is not an integer or out of range, got %v", result)
	}
}

func TestDecrBy(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "counter"},
		{Typ: common.BULK_TYPE, Bulk: "5"}}
	result := DecrBy(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestDecrByInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := DecrBy(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'DecrBy' command, got %v", result)
	}
}

func TestDecrByInvalidValueEssentia(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := DecrBy(args)
	if result.Typ != "error" || result.Str != "ERR invalid decrement value" {
		t.Errorf("Expected error ERR invalid Decrement value, got %v", result)
	}
}

func TestGet(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := Get(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}}
	result := Get(args)
	expected := common.ERR_STRING_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestGetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := Get(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'get' command, got %v", result)
	}
}

func TestGetDel(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := GetDel(args)
	if result.Typ != common.BULK_TYPE && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetDelNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}}
	result := GetDel(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetDelInvalidArgsCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := GetDel(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'GetDel' command, got %v", result)
	}
}

func TestGetEx(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "10"}}
	result := GetEx(args)
	if result.Typ != common.BULK_TYPE {
		t.Errorf("Expected bulk, got %v", result)
	}
}

func TestGetExNonExistingKey(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "non_existing"},
		{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := GetEx(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetExInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := GetEx(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'GetEx' command, got %v", result)
	}
}

func TestGetRange(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "0"},
		{Typ: common.BULK_TYPE, Bulk: "4"}}
	result := GetRange(args)
	if result.Typ != common.BULK_TYPE && result.Typ != "null" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestGetRangeInvalidArgsCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "0"}}
	result := GetRange(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'GetRange' command, got %v", result)
	}
}

func TestGetRangeError(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "non_existing"},
		{Typ: common.BULK_TYPE, Bulk: "0"},
		{Typ: common.BULK_TYPE, Bulk: "4"}}
	result := GetRange(args)
	if result.Typ != "null" {
		t.Errorf("Expected null, got %v", result)
	}
}

func TestGetSet(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "val"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "val1"}}
	result := GetSet(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "val" {
		t.Errorf("Expected val: val, got %v", result)
	}
	result = Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}})
	if result.Typ != common.BULK_TYPE || result.Bulk != "val1" {
		t.Errorf("Expected val: val1, got %v", result)
	}
}

func TestGetSetNonExistingKey(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "non_existing"},
		{Typ: common.BULK_TYPE, Bulk: "val"}}
	expected := common.ERR_STRING_NOT_FOUND
	result := GetSet(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestGetSetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := GetSet(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestSet(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}}
	result := Set(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestSetInvalidArgsCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := Set(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'Set' command, got %v", result)
	}
}

func TestIncr(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "counter"}}
	result := Incr(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: " world"}}
	result := Incr(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'Decr' command, got %v", result)
	}
}

func TestIncrInvalidValueEssentia(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	result := Incr(args)
	if result.Typ != "error" || result.Str != "ERR value is not an integer or out of range" {
		t.Errorf("Expected ERR value is not an integer or out of range, got %v", result)
	}
}

func TestIncrBy(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "counter"},
		{Typ: common.BULK_TYPE, Bulk: "5"}}
	result := IncrBy(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrByInvalidArgumentCount(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := IncrBy(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected ERR %s, got %v", expected, result)
	}
}

func TestIncrByInvalidValueEssentia(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := "ERR invalid increment value"
	result := IncrBy(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestIncrByFloat(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "counter"},
		{Typ: common.BULK_TYPE, Bulk: "5"}}
	result := IncrByFloat(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
}

func TestIncrByInvalidArgumentCountFloat(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := IncrByFloat(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected ERR %s, got %v", expected, result)
	}
}

func TestIncrByInvalidValueEssentiaFloat(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "world"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "hello"},
		{Typ: common.BULK_TYPE, Bulk: "hello"}}
	expected := "ERR invalid increment value"
	result := IncrByFloat(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error %s, got %v", expected, result)
	}
}

func TestLcs(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
		{Typ: common.BULK_TYPE, Bulk: "lasagna"}})
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "baigan"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
		{Typ: common.BULK_TYPE, Bulk: "key2"}}
	result := LCS(args)
	if result.Typ == "error" || result.Bulk != "aga" {
		t.Errorf("Expected value: aga, got %v", result)
	}
}

func TestLcsLen(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
		{Typ: common.BULK_TYPE, Bulk: "lasagna"}})
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "baigan"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "len"}}
	result := LCS(args)
	if result.Typ == "error" || result.Bulk != "3" {
		t.Errorf("Expected value: 3, got %v", result)
	}
}

func TestLcsInvalidFirstKey(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "invalid"},
		{Typ: common.BULK_TYPE, Bulk: "key2"}}
	expected := common.ERR_STRING_NOT_FOUND
	result := LCS(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestLcsInvalidSecondtKey(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
		{Typ: common.BULK_TYPE, Bulk: "invalid"}}
	expected := common.ERR_STRING_NOT_FOUND
	result := LCS(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestLcsInvalidArguments(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key1"}}
	expected := common.ERR_WRONG_ARGUMENT_COUNT
	result := LCS(args)
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected error: %s, got %v", expected, result)
	}
}

func TestMGetAndMSet(t *testing.T) {
	MSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "value2"},
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})

	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "key2"},
		{Typ: common.BULK_TYPE, Bulk: "invalid"}}
	expected := []string{"value", "value2", ""}
	result := MGet(args)
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
	result := MGet(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}

func TestMSetInvalidCommands(t *testing.T) {
	args := []resp.Value{}
	result := MSet(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}

func TestMSetInvalidCommands2(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := MSet(args)
	if result.Typ != "error" {
		t.Errorf("Expected error got %v", result)
	}
}

func TestSetEx(t *testing.T) {
	tests := []struct {
		name     string
		args     []resp.Value
		wantType string
		wantStr  string
		wantNum  int
	}{
		{
			name: "valid setex call",
			args: []resp.Value{
				{Typ: common.BULK_TYPE, Bulk: "key1"},
				{Typ: common.BULK_TYPE, Bulk: "value1"},
				{Typ: common.BULK_TYPE, Bulk: "10"},
			},
			wantType: common.STRING_TYPE,
			wantStr:  "OK",
		},
		{
			name: "missing arguments",
			args: []resp.Value{
				{Typ: common.BULK_TYPE, Bulk: "key1"},
			},
			wantType: common.ERROR_TYPE,
			wantStr:  common.ERR_WRONG_ARGUMENT_COUNT,
		},
		{
			name: "non-numeric expiry",
			args: []resp.Value{
				{Typ: common.BULK_TYPE, Bulk: "key2"},
				{Typ: common.BULK_TYPE, Bulk: "value2"},
				{Typ: common.BULK_TYPE, Bulk: "not-a-number"},
			},
			wantType: common.ERROR_TYPE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetEx(tt.args)

			if result.Typ != tt.wantType {
				t.Errorf("%s: expected type %s, got %s", tt.name, tt.wantType, result.Typ)
			}
			if tt.wantStr != "" && result.Str != tt.wantStr {
				t.Errorf("%s: expected message '%s', got '%s'", tt.name, tt.wantStr, result.Str)
			}
		})
	}

	// Separate check for ExpireTime after valid SetEx
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key1"},
	}
	result := ExpireTime(args)
	if result.Typ != common.INTEGER_TYPE || result.Num <= 0 {
		t.Errorf("Expected expiry time to be set, got %v", result)
	}
}

func TestStrLen(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := StrLen(args)
	if result.Typ != common.INTEGER_TYPE || result.Num != 5 {
		t.Errorf("Expected length 5, got %v", result)
	}
}

func TestStrLenNonExistingKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "non_existing"}}
	result := StrLen(args)
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected length 0, got %v", result)
	}
}

func TestStrLenInvalidArgsCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "key"}}
	result := StrLen(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'StrLen' command, got %v", result)
	}
}

func TestSetRange(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "Hello World"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "6"},
		{Typ: common.BULK_TYPE, Bulk: "Animus"}}
	result := SetRange(args)
	if result.Typ != common.STRING_TYPE || result.Str != "OK" {
		t.Errorf("Expected OK, got %v", result)
	}
	val := Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "key"}})
	if val.Bulk != "Hello Animus" {
		t.Errorf("Expected 'Hello Animus', got %v", val.Bulk)
	}
}

func TestSetRangeInvalidArgsCount(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "6"}}
	result := SetRange(args)
	if result.Typ != "error" || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected ERR wrong number of arguments for 'SetRange' command, got %v", result)
	}
}

func TestSetRangeInvalidRange(t *testing.T) {
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "-6"},
		{Typ: common.BULK_TYPE, Bulk: "Animus"}}
	result := SetRange(args)
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_OUT_OF_RANGE {
		t.Errorf("Expected ERR wrong number of arguments for 'SetRange' command, got %v", result)
	}
}
