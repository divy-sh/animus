package hashes_test

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/types/hashes"
	"github.com/divy-sh/animus/types/strings"
)

func TestHashEssentia_HSetAndHGet(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashes.HSet(hash, key, value)

	// Retrieve the value
	got, err := hashes.HGet(hash, key)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got != value {
		t.Errorf("expected %s, got %s", value, got)
	}
}

func TestHashEssentia_HGet_NotFound(t *testing.T) {

	// Try to get a non-existent hash
	_, err := hashes.HGet("non_existent_hash", "non_existent_key")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HGet_KeyNotFound(t *testing.T) {
	hash := "test_hash"
	key := "test_key"
	value := "test_value"

	hashes.HSet(hash, key, value)
	// Try to get a non-existent key
	_, err := hashes.HGet("test_hash", "non_existent_key")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func TestHashEssentia_HSet_KeyFound(t *testing.T) {
	hashes.HSet("test_hash", "test_key", "test_value")
	hashes.HSet("test_hash", "test_key", "new_value")

	val, err := hashes.HGet("test_hash", "test_key")
	if err != nil || val != "new_value" {
		t.Errorf("expected value 'new_value', got error %v", err)
	}
}

func Test_HashExists(t *testing.T) {
	hashes.HSet("Tesh_HashExists", "Tesh_HashExists", "test_value")
	exists, err := hashes.HExists("Tesh_HashExists", "Tesh_HashExists")
	if exists != 1 || err != nil {
		t.Errorf("Expected hash to exist, got: %d, %v", exists, err)
	}
}

func Test_HashExists_Nope(t *testing.T) {
	exists, err := hashes.HExists("Test_HashExists_Nope", "Test_HashExists_Nope")
	if exists == 1 || err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected hash to not exist, got: %d, %v", exists, err)
	}
}

func Test_Hashes_ExpireNoFlagKeyWithNoExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireNoFlagKeyWithNoExpiry", "Test_Hashes_ExpireNoFlagKeyWithNoExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireNoFlagKeyWithNoExpiry", "0", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	val, err := strings.Get("Test_Hashes_ExpireNoFlagKeyWithNoExpiry")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, val)
	}
}

func Test_Hashes_ExpireNoFlagKeyWithExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireNoFlagKeyWithExpiry", "Test_Hashes_ExpireNoFlagKeyWithExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireNoFlagKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireNoFlagKeyWithExpiry", "0", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	val, err := strings.Get("Test_Hashes_ExpireNoFlagKeyWithExpiry")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_STRING_NOT_FOUND, val)
	}
}

func Test_Hashes_ExpireNoFlagInvalidKey(t *testing.T) {
	err := hashes.HExpire("Test_Hashes_ExpireNoFlagInvalidKey", "10", "")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func Test_Hashes_ExpireNXKeyWithNoExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireNXKeyWithNoExpiry", "Test_Hashes_ExpireNXKeyWithNoExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireNXKeyWithNoExpiry", "10", "NX")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func Test_Hashes_ExpireNXKeyWithExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireNXKeyWithExpiry", "Test_Hashes_ExpireNXKeyWithExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireNXKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireNXKeyWithExpiry", "10", "NX")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func Test_Hashes_ExpireNXInvalidKey(t *testing.T) {
	err := hashes.HExpire("Test_Hashes_ExpireNXInvalidKey", "10", "NX")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func Test_Hashes_ExpireXXKeyWithNoExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireXXKeyWithNoExpiry", "Test_Hashes_ExpireXXKeyWithNoExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireXXKeyWithNoExpiry", "10", "XX")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func Test_Hashes_ExpireXXKeyWithExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireXXKeyWithExpiry", "Test_Hashes_ExpireXXKeyWithExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireXXKeyWithExpiry", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireXXKeyWithExpiry", "10", "XX")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func Test_Hashes_ExpireXXInvalidKey(t *testing.T) {
	err := hashes.HExpire("Test_Hashes_ExpireNXInvalidKey", "10", "XX")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func Test_Hashes_ExpireGTKeyWithNoExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireGTKeyWithNoExpiry", "Test_Hashes_ExpireGTKeyWithNoExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireGTKeyWithNoExpiry", "10", "GT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func Test_Hashes_ExpireGTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireGTKeyWithExpiryNewTimeSmaller", "Test_Hashes_ExpireGTKeyWithExpiryNewTimeSmaller", "value")
	err := hashes.HExpire("Test_Hashes_ExpireGTKeyWithExpiryNewTimeSmaller", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireGTKeyWithExpiryNewTimeSmaller", "10", "GT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func Test_Hashes_ExpireGTKeyWithExpiryNewTimeGreater(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireGTKeyWithExpiryNewTimeGreater", "Test_Hashes_ExpireGTKeyWithExpiryNewTimeGreater", "value")
	err := hashes.HExpire("Test_Hashes_ExpireGTKeyWithExpiryNewTimeGreater", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireGTKeyWithExpiryNewTimeGreater", "200", "GT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func Test_Hashes_ExpireGTInvalidKey(t *testing.T) {
	err := hashes.HExpire("Test_Hashes_ExpireGTInvalidKey", "10", "GT")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func Test_Hashes_ExpireLTKeyWithNoExpiry(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireLTKeyWithNoExpiry", "Test_Hashes_ExpireLTKeyWithNoExpiry", "value")
	err := hashes.HExpire("Test_Hashes_ExpireLTKeyWithNoExpiry", "10", "LT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func Test_Hashes_ExpireLTKeyWithExpiryNewTimeSmaller(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireLTKeyWithExpiryNewTimeSmaller", "Test_Hashes_ExpireLTKeyWithExpiryNewTimeSmaller", "value")
	err := hashes.HExpire("Test_Hashes_ExpireLTKeyWithExpiryNewTimeSmaller", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireLTKeyWithExpiryNewTimeSmaller", "10", "LT")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
}

func Test_Hashes_ExpireLTKeyWithExpiryNewTimeGreater(t *testing.T) {
	hashes.HSet("Test_Hashes_ExpireLTKeyWithExpiryNewTimeGreater", "Test_Hashes_ExpireLTKeyWithExpiryNewTimeGreater", "value")
	err := hashes.HExpire("Test_Hashes_ExpireLTKeyWithExpiryNewTimeGreater", "100", "")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	err = hashes.HExpire("Test_Hashes_ExpireLTKeyWithExpiryNewTimeGreater", "200", "LT")
	if err == nil || err.Error() != common.ERR_EXPIRY_TYPE {
		t.Errorf("Expected error: %s, got %v", common.ERR_EXPIRY_TYPE, err)
	}
}

func Test_Hashes_ExpireLTInvalidKey(t *testing.T) {
	err := hashes.HExpire("Test_Hashes_ExpireLTInvalidKey", "10", "LT")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_SOURCE_KEY_NOT_FOUND, err)
	}
}

func Test_HashDelete(t *testing.T) {
	hashes.HSet("Test_HashDelete", "Test_HashDelete", "value")
	err := hashes.HDel("Test_HashDelete", "Test_HashDelete")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	_, err = hashes.HGet("Test_HashDelete", "Test_HashDelete")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_HASH_NOT_FOUND, err)
	}
}

func Test_HashDelete_MultipleValuesInHash(t *testing.T) {
	hashes.HSet("Test_HashDelete", "Test_HashDelete", "value")
	hashes.HSet("Test_HashDelete", "Test_HashDelete2", "value")
	err := hashes.HDel("Test_HashDelete", "Test_HashDelete")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	_, err = hashes.HGet("Test_HashDelete", "Test_HashDelete")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_HASH_NOT_FOUND, err)
	}
	val, err := hashes.HGet("Test_HashDelete", "Test_HashDelete2")
	if err != nil || val != "value" {
		t.Errorf("Expected value: %s, got: %v, %v", "value", val, err)
	}
}

func Test_HashDelete_NonExistentKey(t *testing.T) {
	err := hashes.HDel("Test_HashDelete_NonExistentKey", "Test_HashDelete_NonExistentKey")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_HASH_NOT_FOUND, err)
	}
}

func Test_HashDelete_NonExistentField(t *testing.T) {
	hashes.HSet("Test_HashDelete_NonExistentField", "existing_field", "value")
	err := hashes.HDel("Test_HashDelete_NonExistentField", "non_existent_field")
	if err == nil || err.Error() != common.ERR_KEY_NOT_FOUND {
		t.Errorf("Expected error: %s, got: %v", common.ERR_KEY_NOT_FOUND, err)
	}
}

func Test_HGetAll(t *testing.T) {
	hash := "Test_HGetAll"
	field1 := "field1"
	value1 := "value1"
	field2 := "field2"
	value2 := "value2"

	hashes.HSet(hash, field1, value1)
	hashes.HSet(hash, field2, value2)

	result, err := hashes.HGetAll(hash)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 fields, got %d", len(result))
	}
	if result[field1] != value1 {
		t.Errorf("expected %s for %s, got %s", value1, field1, result[field1])
	}
	if result[field2] != value2 {
		t.Errorf("expected %s for %s, got %s", value2, field2, result[field2])
	}
}

func Test_HGetAll_NonExistentHash(t *testing.T) {
	_, err := hashes.HGetAll("NonExistentHash")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
}

func Test_HGetAll_EmptyHash(t *testing.T) {
	hash := "EmptyHash"
	hashes.HSet(hash, "field", "value")
	hashes.HDel(hash, "field")

	result, err := hashes.HGetAll(hash)
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("expected error 'ERR hash does not exist', got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected 0 fields, got %d", len(result))
	}
}
