package genericcmd

import (
	"testing"
	"time"

	"fmt"

	"github.com/divy-sh/animus/internal/command/stringcmd"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

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
	stringcmd.Set([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := ExpireAt([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func TestGenerics_ExpireAtNoFlagKeyWithExpiry(t *testing.T) {
	stringcmd.Set([]resp.Value{
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
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "TestGenerics_ExpireAtNoFlagKeyWithExpiry"}})
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
	stringcmd.Set([]resp.Value{
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
