package commands

import (
	"testing"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestStringCopy(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	copyVal([]resp.Value{
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
	copyVal([]resp.Value{
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
	copyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}}
	result := rpop(args)
	if result.Typ != "array" || result.Array[0].Bulk != "value2" {
		t.Errorf("Expected array or null, got %v", result)
	}
}

func TestGeneric_Copy_InvalidArgumentsCount(t *testing.T) {
	copyVal([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "invalidArguemntCount"}})
}

func TestGeneric_Copy_InvalidSourceKey(t *testing.T) {
	copyVal([]resp.Value{
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

func TestGeneric_Exists(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestExistsKey"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestExistsKey"}}
	result := exists(args)
	if result.Typ != common.INTEGER_TYPE || result.Num != 1 {
		t.Errorf("Expected key to exist, got %v", result)
	}
}

func TestGeneric_Exists_InvalidKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestExistsInvalidKey"}}
	result := exists(args)
	if result.Typ != common.INTEGER_TYPE || result.Num != 0 {
		t.Errorf("Expected key to exist, got %v", result)
	}
}

func TestGeneric_Exists_InvalidArguments(t *testing.T) {
	args := []resp.Value{}
	result := exists(args)
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_Expire_InvalidArgumentCount1(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_Expire_InvalidArgumentCount2(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_ExpireNoFlagKeyWithNoExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNoFlagKeyWithExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNoFlagInvalidKey(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNXKeyWithNoExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireNXKeyWithExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireNXInvalidKey(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireXXKeyWithNoExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireXXKeyWithExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireXXInvalidKey(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireGTKeyWithNoExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireGTInvalidKey(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireLTKeyWithNoExpiry(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireLTInvalidKey(t *testing.T) {
	result := expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}
