package generics

import (
	"testing"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/types/hashes"
	"github.com/divy-sh/animus/internal/types/lists"
	"github.com/divy-sh/animus/internal/types/strings"
)

func TestStringCopy(t *testing.T) {
	strings.Set("TestStringCopy", "expected")
	Copy("TestStringCopy", "TestStringCopy2")
	val, err := strings.Get("TestStringCopy2")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestHashCopy(t *testing.T) {
	hashes.HSet("TestHashCopy", "pizza", "expected")
	Copy("TestHashCopy", "TestHashCopy2")
	val, err := hashes.HGet("TestHashCopy2", "pizza")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestListCopy(t *testing.T) {
	lists.RPush("TestListCopy", &[]string{"expected"})
	Copy("TestListCopy", "TestListCopy2")
	val, err := lists.RPop("TestListCopy2", "1")
	if err != nil || val[0] != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestInvalidKeyCopy(t *testing.T) {
	val, err := Copy("TestInvalidKeyCopy", "TestInvalidKeyCopy2")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("%v, %v", val, err)
	}
}

func TestStringDelete(t *testing.T) {
	strings.Set("TestStringDelete", "expected")
	Delete(&[]string{"TestStringDelete"})
	val, err := strings.Get("TestStringDelete")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestHashDelete(t *testing.T) {
	hashes.HSet("TestHashDelete", "pizza", "expected")
	Delete(&[]string{"TestHashDelete"})
	val, err := hashes.HGet("TestHashDelete", "pizza")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestListDelete(t *testing.T) {
	lists.RPush("TestListDelete", &[]string{"expected"})
	Delete(&[]string{"TestListDelete"})
	val, err := lists.RPop("TestListDelete", "1")
	if err == nil || err.Error() != common.ERR_LIST_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestExists(t *testing.T) {
	strings.Set("TestExists", "expected")
	validKeyCount := Exists(&[]string{"TestExists"})
	if validKeyCount != 1 {
		t.Errorf("Expected count to be %d, got: %v", 1, validKeyCount)
	}
}

func TestExistsInvalidKey(t *testing.T) {
	validKeyCount := Exists(&[]string{"TestExistsInvalidKey"})
	if validKeyCount != 0 {
		t.Errorf("Expected count to be %d, got: %v", 0, validKeyCount)
	}
}

func TestGenerics_ExpireNoFlagKeyWithNoExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireNoFlagKeyWithNoExpiry", "value")
	err := Expire("TestGenerics_ExpireNoFlagKeyWithNoExpiry", "0", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	val, err := strings.Get("TestGenerics_ExpireNoFlagKeyWithNoExpiry")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, val)
	}
}

func TestGenerics_ExpireNoFlagKeyWithExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireNoFlagKeyWithExpiry", "value")
	err := Expire("TestGenerics_ExpireNoFlagKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireNoFlagKeyWithExpiry", "0", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	val, err := strings.Get("TestGenerics_ExpireNoFlagKeyWithExpiry")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, val)
	}
}

func TestGenerics_ExpireNoFlagInvalidKey(t *testing.T) {
	err := Expire("TestGenerics_ExpireNoFlagInvalidKey", "10", "")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func TestGenerics_ExpireNXKeyWithNoExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireNXKeyWithNoExpiry", "value")
	err := Expire("TestGenerics_ExpireNXKeyWithNoExpiry", "10", "NX")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestGenerics_ExpireNXKeyWithExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireNXKeyWithExpiry", "value")
	err := Expire("TestGenerics_ExpireNXKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireNXKeyWithExpiry", "10", "NX")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func TestGenerics_ExpireNXInvalidKey(t *testing.T) {
	err := Expire("TestGenerics_ExpireNXInvalidKey", "10", "NX")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func TestGenerics_ExpireXXKeyWithNoExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireXXKeyWithNoExpiry", "value")
	err := Expire("TestGenerics_ExpireXXKeyWithNoExpiry", "10", "XX")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func TestGenerics_ExpireXXKeyWithExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireXXKeyWithExpiry", "value")
	err := Expire("TestGenerics_ExpireXXKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireXXKeyWithExpiry", "10", "XX")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestGenerics_ExpireXXInvalidKey(t *testing.T) {
	err := Expire("TestGenerics_ExpireNXInvalidKey", "10", "XX")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func TestGenerics_ExpireGTKeyWithNoExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireGTKeyWithNoExpiry", "value")
	err := Expire("TestGenerics_ExpireGTKeyWithNoExpiry", "10", "GT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	strings.Set("TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller", "value")
	err := Expire("TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireGTKeyWithExpiryNewTimeSmaller", "10", "GT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	strings.Set("TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater", "value")
	err := Expire("TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireGTKeyWithExpiryNewTimeGreater", "200", "GT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestGenerics_ExpireGTInvalidKey(t *testing.T) {
	err := Expire("TestGenerics_ExpireGTInvalidKey", "10", "GT")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func TestGenerics_ExpireLTKeyWithNoExpiry(t *testing.T) {
	strings.Set("TestGenerics_ExpireLTKeyWithNoExpiry", "value")
	err := Expire("TestGenerics_ExpireLTKeyWithNoExpiry", "10", "LT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	strings.Set("TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller", "value")
	err := Expire("TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireLTKeyWithExpiryNewTimeSmaller", "10", "LT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	strings.Set("TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater", "value")
	err := Expire("TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = Expire("TestGenerics_ExpireLTKeyWithExpiryNewTimeGreater", "200", "LT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func TestGenerics_ExpireLTInvalidKey(t *testing.T) {
	err := Expire("TestGenerics_ExpireLTInvalidKey", "10", "LT")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func TestGenerics_KeysNoKeys(t *testing.T) {
	keys, err := Keys("nonExisting")
	if err != nil || len(*keys) > 0 {
		t.Errorf("expected no keys, got keys: %v, error: %v", keys, err)
	}
}

func TestGenerics_Keys(t *testing.T) {
	strings.Set("TestGenerics_Keys", "value")
	hashes.HSet("TestGenerics_Keys1", "a", "b")
	lists.RPush("non_matching_key", &[]string{"a"})
	keys, err := Keys("TestGenerics_Key")
	if err != nil || len(*keys) != 2 {
		t.Errorf("expected multiple keys, got: %v, error: %v", keys, err)
	}
}

func TestGenerics_Keys_invalidRegex(t *testing.T) {
	_, err := Keys("[a-b")
	if err == nil || err.Error() != common.ERR_INVALID_REGEX {
		t.Errorf("expected error: %v, got: %v", common.ERR_INVALID_REGEX, err)
	}
}
