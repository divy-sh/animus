package store

import (
	"testing"
	"time"
)

func TestStore_GetLock(t *testing.T) {
	Set("GetLock", "GetLock")
	lock := GetLock("GetLock")
	lock.Lock()
	Get[string, string]("GetLock")
	lock.Unlock()
}

func TestStore_GetLocks(t *testing.T) {
	Set("GetLocks", "GetLocks")
	lock := GetLocks(&[]any{"GetLocks"})
	lock[0].Lock()
	Get[string, string]("GetLocks")
	lock[0].Unlock()
}

func TestStore_SetAndGet(t *testing.T) {
	key, value := "testKey", "testValue"
	SetWithTTL(key, value, 60)
	storedValue, found := Get[string, string](key)

	if !found {
		t.Errorf("Expected key to be found, but it was not")
	}
	if storedValue != value {
		t.Errorf("Expected value %v, got %v", value, storedValue)
	}
}

func TestStore_GetNonExistentKey(t *testing.T) {
	_, found := Get[string, string]("nonExistentKey")

	if found {
		t.Errorf("Expected key to not be found, but it was")
	}
}

func TestStore_OverwriteKey(t *testing.T) {
	key := "testKey"
	firstValue, secondValue := "firstValue", "secondValue"

	Set(key, firstValue)
	Set(key, secondValue)

	storedValue, _ := Get[string, string](key)
	if storedValue != secondValue {
		t.Errorf("Expected value %v, got %v", secondValue, storedValue)
	}
}

func TestStore_OverwriteKeyWithTTL(t *testing.T) {
	key := "testKey"
	firstValue, secondValue := "firstValue", "secondValue"

	SetWithTTL(key, firstValue, 120)
	SetWithTTL(key, secondValue, 120)

	storedValue, _ := Get[string, string](key)
	if storedValue != secondValue {
		t.Errorf("Expected value %v, got %v", secondValue, storedValue)
	}
}

func TestStore_TTLExpirationOnGet(t *testing.T) {
	SetWithTTL("expiringKey", "value", 0)
	time.Sleep(1 * time.Microsecond) // Allow time for expiration
	Get[string, string]("expiringKey")
	_, found := Get[string, string]("expiringKey")
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_TTLExpirationOnSet(t *testing.T) {
	SetWithTTL("expiringKey", "value", 0)
	time.Sleep(1 * time.Microsecond) // Allow time for expiration
	Set("newKey", "value")
	_, found := Get[string, string]("expiringKey")
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_DeleteWithKey(t *testing.T) {
	Set("expiringKey", "value")
	_, found := Get[string, string]("expiringKey")
	if !found {
		t.Errorf("Expected key to be found but was not found")
	}
	DeleteWithKey("expiringKey")
	_, found = Get[string, string]("expiringKey")
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_TTL_Eviction(t *testing.T) {
	store := GetSharedStore()
	store.maxSize = 3
	SetWithTTL("key1", "value1", 60)
	SetWithTTL("key2", "value2", 60)
	SetWithTTL("key3", "value3", 0)
	// This should evict key3, even though LRU should've evicted key1
	SetWithTTL("key4", "value4", 60)
	store.maxSize = 100_000 // Revert size
	_, found := Get[string, string]("key3")
	if found {
		t.Errorf("Expected LRU key 'key3' to be evicted, but it was found")
	}
}

func TestStore_LRU_Eviction(t *testing.T) {
	store := GetSharedStore()
	store.maxSize = 3
	Set("LruEviction1", "value1")
	Set("LruEviction2", "value2")
	Set("LruEviction3", "value3")
	Set("LruEviction4", "value4") // This should trigger LRU eviction
	store.maxSize = 100_000       // Revert size

	_, found := Get[string, string]("key1")
	if found {
		t.Errorf("Expected LRU key 'key1' to be evicted, but it was found")
	}
}
