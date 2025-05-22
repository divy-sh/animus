package store

import (
	"testing"
	"time"
)

func TestExpiryCleanerRemovesExpiredKeys(t *testing.T) {
	SetWithTTL("test-key", "test-value", 0)

	// Wait enough time for it to expire and cleaner to run
	time.Sleep(1 * time.Millisecond)

	// Try to get the value again
	_, ok := Get[string, string]("test-key")
	if ok {
		t.Fatal("expected value to be expired and cleaned, but it was still present")
	}
}

func TestSetAndGet(t *testing.T) {
	key := "foo"
	val := "bar"

	Set(key, val)
	result, ok := Get[string, string](key)
	if !ok {
		t.Errorf("expected key %v to exist", key)
	}
	if result != val {
		t.Errorf("expected value %v, got %v", val, result)
	}
}

func TestSetWithTTLAndGet(t *testing.T) {
	key := "expiringKey"
	val := "willExpire"

	SetWithTTL(key, val, 0)

	time.Sleep(2 * time.Millisecond)

	result, ok := Get[string, string](key)
	if ok {
		t.Errorf("expected key %v to be expired, got value %v", key, result)
	}
}

func TestSetWithTTLAsUnixTimeStampAndGet(t *testing.T) {
	key := "expiringKey"
	val := "willExpire"

	SetWithTTLAsUnixTimeStamp(key, val, time.Now().Unix())

	time.Sleep(2 * time.Millisecond)

	result, ok := Get[string, string](key)
	if ok {
		t.Errorf("expected key %v to be expired, got value %v", key, result)
	}
}

func TestGetWithTTL(t *testing.T) {
	key := "ttlKey"
	val := "someVal"
	ttl := int64(5) // 5 seconds TTL

	SetWithTTL(key, val, ttl)
	result, storedTTL, ok := GetWithTTL[string, string](key)
	if !ok {
		t.Errorf("expected key %v to exist", key)
	}
	if result != val {
		t.Errorf("expected value %v, got %v", val, result)
	}
	if storedTTL <= 0 {
		t.Errorf("expected TTL to be > 0, got %v", storedTTL)
	}
}

func TestGetWithTTLExpiredKey(t *testing.T) {
	key := "ttlKeyExpired"
	val := "someVal"
	ttl := int64(0)

	SetWithTTL(key, val, ttl)
	result, storedTTL, ok := GetWithTTL[string, string](key)
	if ok || result != "" || storedTTL != -1 {
		t.Errorf("expected key %v to exist", key)
	}
}

func TestGetWithTTLInvalidKey(t *testing.T) {
	key := "invalidKey"
	result, storedTTL, ok := GetWithTTL[string, string](key)
	if ok || result != "" || storedTTL != -1 {
		t.Errorf("expected key %v to exist", key)
	}
}

func TestDelete(t *testing.T) {
	key := "tempKey"
	val := "tempVal"

	Set(key, val)
	Delete(key)

	_, ok := Get[string, string](key)
	if ok {
		t.Errorf("expected key %v to be deleted", key)
	}
}

func TestOverwriteValue(t *testing.T) {
	key := "overwrite"
	val1 := "initial"
	val2 := "updated"

	Set(key, val1)
	Set(key, val2)

	result, ok := Get[string, string](key)
	if !ok || result != val2 {
		t.Errorf("expected updated value %v, got %v", val2, result)
	}
}

func TestGetKeys(t *testing.T) {
	key := "TestGetKeys"
	val := "TestGetKeysVal"

	Set(key, val)
	result := GetKeys[string]()
	if len(*result) < 1 {
		t.Errorf("expected atleast 1 key but got no keys")
	}
}
