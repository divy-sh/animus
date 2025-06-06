package hashcmd

import (
	"testing"

	"github.com/divy-sh/animus/internal/command/stringcmd"
	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/resp"
)

func Test_Hashes_HExpire_InvalidArgumentCount1(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func Test_Hashes_HExpire_InvalidArgumentCount2(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_WRONG_ARGUMENT_COUNT {
		t.Errorf("Expected error: %s, got: %v", common.ERR_WRONG_ARGUMENT_COUNT, result)
	}
}

func Test_Hashes_HExpireNoFlagKeyWithNoExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithNoExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireNoFlagKeyWithExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "0"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = stringcmd.Get([]resp.Value{{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagKeyWithExpiry"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireNoFlagInvalidKey(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNoFlagInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireNXKeyWithNoExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func Test_Hashes_HExpireNXKeyWithExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func Test_Hashes_HExpireNXInvalidKey(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "NX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireXXKeyWithNoExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func Test_Hashes_HExpireXXKeyWithExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireXXKeyWithExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func Test_Hashes_HExpireXXInvalidKey(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireNXInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "XX"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireGTKeyWithNoExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func Test_Hashes_HExpireGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func Test_Hashes_HExpireGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func Test_Hashes_HExpireGTInvalidKey(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireGTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "GT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}

func Test_Hashes_HExpireLTKeyWithNoExpiry(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithNoExpiry"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func Test_Hashes_HExpireLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeSmaller"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
}

func Test_Hashes_HExpireLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	HSet([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "value"}})
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "100"}})
	if result.Typ == common.ERROR_TYPE {
		t.Errorf("Expected no error, got: %s", result.Str)
	}
	result = HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTKeyWithExpiryNewTimeGreater"},
		{Typ: common.BULK_TYPE, Bulk: "200"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, result)
	}
}

func Test_Hashes_HExpireLTInvalidKey(t *testing.T) {
	result := HExpire([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "Test_Hashes_HExpireLTInvalidKey"},
		{Typ: common.BULK_TYPE, Bulk: "10"},
		{Typ: common.BULK_TYPE, Bulk: "LT"}})
	if result.Typ != common.ERROR_TYPE || result.Str != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, result)
	}
}
