package command

import (
	"fmt"
	"testing"
	"time"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestStringCopy(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringCopy2"}}
	result := Get(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestHashCopy(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy2"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashCopy1"}}
	result := HGet(args)
	if result.Typ != common.BULK_TYPE || result.Bulk != "value" {
		t.Errorf("Expected bulk or null, got %v", result)
	}
}

func TestListCopy(t *testing.T) {
	RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	CopyVal([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy1"},
		{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListCopy2"}}
	result := RPop(args)
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
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestStringDelete1"}}
	result := Get(args)
	expected := common.ERR_STRING_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestHashDelete(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}})
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "TestHashDelete1"}}
	result := HGet(args)
	expected := common.ERR_HASH_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestListDelete(t *testing.T) {
	RPush([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"}})
	Del([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete1"}})
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestListDelete2"}}
	result := RPop(args)
	expected := common.ERR_LIST_NOT_FOUND
	if result.Typ != "error" || result.Str != expected {
		t.Errorf("Expected %s, got %v", expected, result)
	}
}

func TestGeneric_Delete_InvalidArgumentsCount(t *testing.T) {
	Del([]resp.Value{})
}

func TestGeneric_Exists(t *testing.T) {
	Set([]resp.Value{
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
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireNoFlagKeyWithExpiry(t *testing.T) {
	Set([]resp.Value{
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
	result = Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireNoFlagKeyWithExpiry"}})
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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
	Set([]resp.Value{
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

func TestGenerics_ExpireAt_InvalidArgumentCount1(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_ExpireAt_InvalidArgumentCount2(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestGenerics_ExpireAtNoFlagKeyWithNoExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtNoFlagKeyWithExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtNoFlagInvalidKey(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtNXKeyWithNoExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireAtNXKeyWithExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireAtNXInvalidKey(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtXXKeyWithNoExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireAtXXKeyWithExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireAtXXInvalidKey(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtGTKeyWithNoExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireAtGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireAtGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireAtGTInvalidKey(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtGTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtLTKeyWithNoExpiry(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireAtLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func TestGenerics_ExpireAtLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: fmt.Sprint(time.Now().Unix() + 100)}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = Expire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func TestGenerics_ExpireAtLTInvalidKey(t *testing.T) {
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtLTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireTime(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"}})
	if result.Typ != common.INTEGER_TYPE || result.Num != -1 {
		t.Errorf("Expected value %d, got: %v", -1, result)
	}
}

func TestGenerics_ExpireTime_InvalidKey(t *testing.T) {
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime_InvalidKey"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireTime_InvaliArgCount(t *testing.T) {
	result := ExpireTime([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"},
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireTime"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func TestKeys(t *testing.T) {
	Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestKeys"},
		{Typ: common.BULK_TYPE, Bulk: "value"},
	})

	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "Test.eys"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}
	if len(result.Array) != 1 || result.Array[0].Bulk != "TestKeys" {
		t.Errorf("expected ['TestKeys'], got %v", result.Array)
	}
}

func TestKeysWrongArgCount(t *testing.T) {
	result := Keys([]resp.Value{})

	if result.Typ != common.ERROR_TYPE {
		t.Errorf("expected ERROR_TYPE, got %v", result.Typ)
	}
	if result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("expected error '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}

func TestKeysInvalidRegex(t *testing.T) {
	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "[a-b"}})

	if result.Typ != common.ERROR_TYPE {
		t.Errorf("expected ERROR_TYPE, got %v", result.Typ)
	}
	if result.Str != common.ERR_INVALID_REGEX {
		t.Errorf("expected error '%s', got '%s'", common.ERR_WRONG_ARGUMENT_COUNT, result.Str)
	}
}

func TestKeysMultipleMatches(t *testing.T) {
	Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:1"}, {Typ: common.BULK_TYPE, Bulk: "red"}})
	Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:2"}, {Typ: common.BULK_TYPE, Bulk: "green"}})
	Set([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "banana:1"}, {Typ: common.BULK_TYPE, Bulk: "yellow"}})

	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "apple:*"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}

	got := map[string]bool{}
	for _, v := range result.Array {
		got[v.Bulk] = true
	}

	if !(got["apple:1"] && got["apple:2"]) || len(result.Array) != 2 {
		t.Errorf("expected ['apple:1', 'apple:2'], got %v", result.Array)
	}
}

func TestKeysNoMatch(t *testing.T) {
	result := Keys([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "nonexistent*"}})

	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("expected ARRAY_TYPE, got %v", result.Typ)
	}

	if len(result.Array) != 0 {
		t.Errorf("expected empty result, got %v", result.Array)
	}
}
