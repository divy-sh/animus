package generics

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/types/hashes"
	"github.com/divy-sh/animus/types/lists"
	"github.com/divy-sh/animus/types/strings"
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
