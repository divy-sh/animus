package generics_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/types/generics"
	"github.com/divy-sh/animus/types/hashes"
	"github.com/divy-sh/animus/types/lists"
	"github.com/divy-sh/animus/types/strings"
)

func TestStringCopy(t *testing.T) {
	strings.Set("TestStringCopy", "expected")
	generics.Copy("TestStringCopy", "TestStringCopy2")
	val, err := strings.Get("TestStringCopy2")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestHashCopy(t *testing.T) {
	hashes.HSet("TestHashCopy", "pizza", "expected")
	generics.Copy("TestHashCopy", "TestHashCopy2")
	val, err := hashes.HGet("TestHashCopy2", "pizza")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestListCopy(t *testing.T) {
	lists.RPush("TestListCopy", &[]string{"expected"})
	generics.Copy("TestListCopy", "TestListCopy2")
	val, err := lists.RPop("TestListCopy2", "1")
	if err != nil || val[0] != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestInvalidKeyCopy(t *testing.T) {
	val, err := generics.Copy("TestInvalidKeyCopy", "TestInvalidKeyCopy2")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("%v, %v", val, err)
	}
}

func TestStringDelete(t *testing.T) {
	strings.Set("TestStringDelete", "expected")
	generics.Delete(&[]string{"TestStringDelete"})
	val, err := strings.Get("TestStringDelete")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestHashDelete(t *testing.T) {
	hashes.HSet("TestHashDelete", "pizza", "expected")
	generics.Delete(&[]string{"TestHashDelete"})
	val, err := hashes.HGet("TestHashDelete", "pizza")
	if err == nil || err.Error() != common.ERR_HASH_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestListDelete(t *testing.T) {
	lists.RPush("TestListDelete", &[]string{"expected"})
	generics.Delete(&[]string{"TestListDelete"})
	val, err := lists.RPop("TestListDelete", "1")
	if err == nil || err.Error() != common.ERR_LIST_NOT_FOUND {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestExists(t *testing.T) {
	strings.Set("TestExists", "expected")
	validKeyCount := generics.Exists(&[]string{"TestExists"})
	if validKeyCount != 1 {
		t.Errorf("Expected count to be %d, got: %v", 1, validKeyCount)
	}
}

func TestExistsInvalidKey(t *testing.T) {
	validKeyCount := generics.Exists(&[]string{"TestExistsInvalidKey"})
	if validKeyCount != 0 {
		t.Errorf("Expected count to be %d, got: %v", 0, validKeyCount)
	}
}

func TestGenerics_Expire(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		setupExpiry   string
		expiry        string
		flag          string
		initialValue  string
		expectedError string
		checkExists   bool
	}{
		// No flag tests
		{
			name:          "NoFlag_KeyWithNoExpiry",
			key:           "NoFlag_KeyWithNoExpiry",
			expiry:        "0",
			flag:          "",
			initialValue:  "value",
			expectedError: "",
			checkExists:   true,
		},
		{
			name:          "NoFlag_KeyWithExpiry",
			key:           "NoFlag_KeyWithExpiry",
			setupExpiry:   "100",
			expiry:        "0",
			flag:          "",
			initialValue:  "value",
			expectedError: "",
			checkExists:   true,
		},
		{
			name:          "NoFlag_InvalidKey",
			key:           "NoFlag_InvalidKey",
			expiry:        "10",
			flag:          "",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},

		// NX tests
		{
			name:          "NX_KeyWithNoExpiry",
			key:           "NX_KeyWithNoExpiry",
			expiry:        "10",
			flag:          "NX",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "NX_KeyWithExpiry",
			key:           "NX_KeyWithExpiry",
			setupExpiry:   "100",
			expiry:        "10",
			flag:          "NX",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "NX_InvalidKey",
			key:           "NX_InvalidKey",
			expiry:        "10",
			flag:          "NX",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},

		// XX tests
		{
			name:          "XX_KeyWithNoExpiry",
			key:           "XX_KeyWithNoExpiry",
			expiry:        "10",
			flag:          "XX",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "XX_KeyWithExpiry",
			key:           "XX_KeyWithExpiry",
			setupExpiry:   "100",
			expiry:        "10",
			flag:          "XX",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "XX_InvalidKey",
			key:           "XX_InvalidKey",
			expiry:        "10",
			flag:          "XX",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},

		// GT tests
		{
			name:          "GT_KeyWithNoExpiry",
			key:           "GT_KeyWithNoExpiry",
			expiry:        "10",
			flag:          "GT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "GT_KeyWithExpiryNewTimeSmaller",
			key:           "GT_KeyWithExpiryNewTimeSmaller",
			setupExpiry:   "100",
			expiry:        "10",
			flag:          "GT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "GT_KeyWithExpiryNewTimeGreater",
			key:           "GT_KeyWithExpiryNewTimeGreater",
			setupExpiry:   "100",
			expiry:        "200",
			flag:          "GT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "GT_InvalidKey",
			key:           "GT_InvalidKey",
			expiry:        "10",
			flag:          "GT",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},

		// LT tests
		{
			name:          "LT_KeyWithNoExpiry",
			key:           "LT_KeyWithNoExpiry",
			expiry:        "10",
			flag:          "LT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "LT_KeyWithExpiryNewTimeSmaller",
			key:           "LT_KeyWithExpiryNewTimeSmaller",
			setupExpiry:   "100",
			expiry:        "10",
			flag:          "LT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "LT_KeyWithExpiryNewTimeGreater",
			key:           "LT_KeyWithExpiryNewTimeGreater",
			setupExpiry:   "100",
			expiry:        "200",
			flag:          "LT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "LT_InvalidKey",
			key:           "LT_InvalidKey",
			expiry:        "10",
			flag:          "LT",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.initialValue != "" {
				strings.Set(tt.key, tt.initialValue)
				if tt.setupExpiry != "" {
					err := generics.Expire(tt.key, tt.setupExpiry, "")
					if err != nil {
						t.Errorf("Failed to setup expiry: %s", err.Error())
						return
					}
				}
			}

			err := generics.Expire(tt.key, tt.expiry, tt.flag)

			if tt.expectedError == "" && err != nil {
				t.Errorf("Expected no error, got: %s", err.Error())
				return
			}
			if tt.expectedError != "" && (err == nil || err.Error() != tt.expectedError) {
				t.Errorf("Expected error: %s, got: %v", tt.expectedError, err)
				return
			}

			if tt.checkExists {
				val, err := strings.Get(tt.key)
				if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
					t.Errorf("Expected error: %s, got value: %v", common.ERR_STRING_NOT_FOUND, val)
				}
			}
		})
	}
}

func TestGenerics_ExpireAt(t *testing.T) {
	currentTime := time.Now().Unix()
	tests := []struct {
		name          string
		key           string
		setupExpiry   string
		expiry        string
		flag          string
		initialValue  string
		expectedError string
		checkExists   bool
	}{
		// No flag tests
		{
			name:          "NoFlag_KeyWithNoExpiry",
			key:           "TestGenerics_ExpireAt_NoFlag_KeyWithNoExpiry",
			expiry:        "0",
			flag:          "",
			initialValue:  "value",
			expectedError: "",
			checkExists:   true,
		},
		{
			name:          "NoFlag_KeyWithExpiry",
			key:           "TestGenerics_ExpireAt_NoFlag_KeyWithExpiry",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        "0",
			flag:          "",
			initialValue:  "value",
			expectedError: "",
			checkExists:   true,
		},
		{
			name:          "NoFlag_InvalidKey",
			key:           "TestGenerics_ExpireAt_NoFlag_InvalidKey",
			expiry:        "10",
			flag:          "",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
		// NX tests
		{
			name:          "NX_KeyWithNoExpiry",
			key:           "TestGenerics_ExpireAt_NX_KeyWithNoExpiry",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "NX",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "NX_KeyWithExpiry",
			key:           "TestGenerics_ExpireAt_NX_KeyWithExpiry",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "NX",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "NX_InvalidKey",
			key:           "TestGenerics_ExpireAt_NX_InvalidKey",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "NX",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
		// XX tests
		{
			name:          "XX_KeyWithNoExpiry",
			key:           "TestGenerics_ExpireAt_XX_KeyWithNoExpiry",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "XX",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "XX_KeyWithExpiry",
			key:           "TestGenerics_ExpireAt_XX_KeyWithExpiry",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "XX",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "XX_InvalidKey",
			key:           "TestGenerics_ExpireAt_XX_InvalidKey",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "XX",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
		// GT tests
		{
			name:          "GT_KeyWithNoExpiry",
			key:           "TestGenerics_ExpireAt_GT_KeyWithNoExpiry",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "GT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "GT_KeyWithExpiryNewTimeSmaller",
			key:           "TestGenerics_ExpireAt_GT_KeyWithExpiryNewTimeSmaller",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "GT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "GT_KeyWithExpiryNewTimeGreater",
			key:           "TestGenerics_ExpireAt_GT_KeyWithExpiryNewTimeGreater",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 200),
			flag:          "GT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "GT_InvalidKey",
			key:           "TestGenerics_ExpireAt_GT_InvalidKey",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "GT",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
		// LT tests
		{
			name:          "LT_KeyWithNoExpiry",
			key:           "TestGenerics_ExpireAt_LT_KeyWithNoExpiry",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "LT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "LT_KeyWithExpiryNewTimeSmaller",
			key:           "TestGenerics_ExpireAt_LT_KeyWithExpiryNewTimeSmaller",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "LT",
			initialValue:  "value",
			expectedError: "",
		},
		{
			name:          "LT_KeyWithExpiryNewTimeGreater",
			key:           "TestGenerics_ExpireAt_LT_KeyWithExpiryNewTimeGreater",
			setupExpiry:   fmt.Sprint(currentTime + 100),
			expiry:        fmt.Sprint(currentTime + 200),
			flag:          "LT",
			initialValue:  "value",
			expectedError: common.ERR_EXPIRY_TYPE,
		},
		{
			name:          "LT_InvalidKey",
			key:           "TestGenerics_ExpireAt_LT_InvalidKey",
			expiry:        fmt.Sprint(currentTime + 10),
			flag:          "LT",
			expectedError: common.ERR_SOURCE_KEY_NOT_FOUND,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.initialValue != "" {
				strings.Set(tt.key, tt.initialValue)
				if tt.setupExpiry != "" {
					err := generics.ExpireAt(tt.key, tt.setupExpiry, "")
					if err != nil {
						t.Errorf("Failed to setup expiry: %s", err.Error())
						return
					}
				}
			}

			err := generics.ExpireAt(tt.key, tt.expiry, tt.flag)

			if tt.expectedError == "" && err != nil {
				t.Errorf("Expected no error, got: %s", err.Error())
				return
			}
			if tt.expectedError != "" && (err == nil || err.Error() != tt.expectedError) {
				t.Errorf("Expected error: %s, got: %v", tt.expectedError, err)
				return
			}

			if tt.checkExists {
				val, err := strings.Get(tt.key)
				if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
					t.Errorf("Expected error: %s, got value: %v", common.ERR_STRING_NOT_FOUND, val)
				}
			}
		})
	}
}

func TestGenerics_KeysNoKeys(t *testing.T) {
	keys, err := generics.Keys("nonExisting")
	if err != nil || len(*keys) > 0 {
		t.Errorf("expected no keys, got keys: %v, error: %v", keys, err)
	}
}

func TestGenerics_Keys(t *testing.T) {
	strings.Set("TestGenerics_Keys", "value")
	hashes.HSet("TestGenerics_Keys1", "a", "b")
	lists.RPush("non_matching_key", &[]string{"a"})
	keys, err := generics.Keys("TestGenerics_Key")
	if err != nil || len(*keys) != 2 {
		t.Errorf("expected multiple keys, got: %v, error: %v", keys, err)
	}
}

func TestGenerics_Keys_invalidRegex(t *testing.T) {
	_, err := generics.Keys("[a-b")
	if err == nil || err.Error() != common.ERR_INVALID_REGEX {
		t.Errorf("expected error: %v, got: %v", common.ERR_INVALID_REGEX, err)
	}
}

func Test_Generics_ExpireTime(t *testing.T) {
	strings.Set("Test_Generics_ExpireTime", "value")
	val, err := generics.ExpireTime("Test_Generics_ExpireTime")
	if err != nil || val != -1 {
		t.Errorf("expected %d, got value: %d, error: %v", -1, val, err)
	}
}

func Test_Generics_ExpireTime_InvalidKey(t *testing.T) {
	val, err := generics.ExpireTime("Test_Generics_ExpireTime_InvalidKey")
	if err == nil || err.Error() != common.ERR_SOURCE_KEY_NOT_FOUND {
		t.Errorf("expected error: %s, got value: %d, error: %v", common.ERR_SOURCE_KEY_NOT_FOUND, val, err)
	}
}
