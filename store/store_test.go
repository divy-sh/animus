package store

import (
	"testing"
	"time"
)

func TestStore_SetAndGet(t *testing.T) {
	key, value := "testKey", "testValue"
	Set(key, value, time.Now().Add(time.Minute))
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

	Set(key, firstValue, time.Now().Add(time.Minute))
	Set(key, secondValue, time.Now().Add(time.Minute))

	storedValue, _ := Get[string, string](key)
	if storedValue != secondValue {
		t.Errorf("Expected value %v, got %v", secondValue, storedValue)
	}
}

func TestStore_TTLExpirationOnGet(t *testing.T) {
	Set("expiringKey", "value", time.Now().Add(50*time.Millisecond))
	time.Sleep(100 * time.Millisecond) // Allow time for expiration
	Get[string, string]("expiringKey")
	_, found := Get[string, string]("expiringKey")
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_TTLExpirationOnSet(t *testing.T) {
	Set("expiringKey", "value", time.Now().Add(50*time.Millisecond))
	time.Sleep(100 * time.Millisecond) // Allow time for expiration
	Set("newKey", "value", time.Now().AddDate(99, 0, 0))
	_, found := Get[string, string]("expiringKey")
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_DeleteWithKey(t *testing.T) {
	Set("expiringKey", "value", time.Now().AddDate(99, 0, 0))
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

// Need to come up with a way to mock a store creation with limited size to tes this.
// func TestStore_LRU_Eviction(t *testing.T) {
// 	Set("key1", "value1", time.Now().Add(time.Minute))
// 	Set("key2", "value2", time.Now().Add(time.Minute))
// 	Set("key3", "value3", time.Now().Add(time.Minute))
// 	Set("key4", "value4", time.Now().Add(time.Minute)) // This should trigger LRU eviction

// 	_, found := Get[string, string]("key1")
// 	if found {
// 		t.Errorf("Expected LRU key 'key1' to be evicted, but it was found")
// 	}
// }
