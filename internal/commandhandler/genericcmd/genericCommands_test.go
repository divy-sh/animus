package genericcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/commandhandler/hashcmd"
	"github.com/divy-sh/animus/internal/commandhandler/listcmd"
	"github.com/divy-sh/animus/internal/commandhandler/stringcmd"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func TestStringCopy(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}}
	result := stringcmd.Get(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestHashCopy(t *testing.T) {
	hashcmd.HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"}}
	result := hashcmd.HGet(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestListCopy(t *testing.T) {
	listcmd.RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}}
	result := listcmd.RPop(args)
	if result.Typ != "array" || result.Array[0].Bulk != "value2" {
		t.Errorf("Expected array or null, got %v", result)
	}
}

func TestGeneric_Copy_InvalidArgumentsCount(t *testing.T) {
	CopyVal([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "invalidArguemntCount"}})
}

func TestGeneric_Copy_InvalidSourceKey(t *testing.T) {
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "InvalidSourceKey"},
		{Typ: common.BULK_TYPE, Bulk: "InvalidDestinationKey"},
	})
}

func TestStringDelete(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}}
	result := stringcmd.Get(args)
	expected := common.ERR_STRING_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestHashDelete(t *testing.T) {
	hashcmd.HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}}
	result := hashcmd.HGet(args)
	expected := common.ERR_HASH_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestListDelete(t *testing.T) {
	listcmd.RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete2"}}
	result := listcmd.RPop(args)
	expected := common.ERR_LIST_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestGeneric_Delete_InvalidArgumentsCount(t *testing.T) {
	Del([]resp.Value{})
}

func TestGeneric_Exists(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestExistsKey"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestExistsKey"}}
	result := Exists(args)
	if result.Typ != common.INTEGER_TYPE || result.Num != 1 {
		t.Errorf("Expected key to exist, got %v", result)
	}
}

func TestGeneric_Exists_InvalidKey(t *testing.T) {
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestExistsInvalidKey"}}
	result := Exists(args)
	if result.Typ != common.INTEGER_TYPE || result.Num != 0 {
		t.Errorf("Expected key to exist, got %v", result)
	}
}

func TestGeneric_Exists_InvalidArguments(t *testing.T) {
	args := []resp.Value{}
	result := Exists(args)
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_Expire_InvalidArgumentCount1(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_Expire_InvalidArgumentCount2(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_ExpireNoFlagKeyWithNoExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNoFlagKeyWithExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNoFlagInvalidKey(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNXKeyWithNoExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireNXKeyWithExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireNXInvalidKey(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireXXKeyWithNoExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireXXKeyWithExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireXXInvalidKey(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireGTKeyWithNoExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireGTInvalidKey(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireGTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireLTKeyWithNoExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireLTInvalidKey(t *testing.T) {
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireLTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}
