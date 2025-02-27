package store

import (
	"testing"
	"time"
)

func TestStore_SetAndGet(t *testing.T) {
	s := NewStore[string, string]()
	key, value := "testKey", "testValue"
	s.Set(key, value, time.Now().Add(time.Minute))
	storedValue, found := s.Get(key)

	if !found {
		t.Errorf("Expected key to be found, but it was not")
	}
	if storedValue != value {
		t.Errorf("Expected value %v, got %v", value, storedValue)
	}
}

func TestStore_GetNonExistentKey(t *testing.T) {
	s := NewStore[string, string]()
	_, found := s.Get("nonExistentKey")

	if found {
		t.Errorf("Expected key to not be found, but it was")
	}
}

func TestStore_OverwriteKey(t *testing.T) {
	s := NewStore[string, string]()
	key := "testKey"
	firstValue, secondValue := "firstValue", "secondValue"

	s.Set(key, firstValue, time.Now().Add(time.Minute))
	s.Set(key, secondValue, time.Now().Add(time.Minute))

	storedValue, _ := s.Get(key)
	if storedValue != secondValue {
		t.Errorf("Expected value %v, got %v", secondValue, storedValue)
	}
}

func TestStore_TTLExpiration(t *testing.T) {
	s := NewStore[string, string]()
	key, value := "expiringKey", "value"
	s.Set(key, value, time.Now().Add(50*time.Millisecond))
	time.Sleep(100 * time.Millisecond) // Allow time for expiration

	_, found := s.Get(key)
	if found {
		t.Errorf("Expected key to expire, but it was found")
	}
}

func TestStore_LRU_Eviction(t *testing.T) {
	s := NewStore[string, string]()
	s.maxSize = 3
	s.Set("key1", "value1", time.Now().Add(time.Minute))
	s.Set("key2", "value2", time.Now().Add(time.Minute))
	s.Set("key3", "value3", time.Now().Add(time.Minute))
	s.Set("key4", "value4", time.Now().Add(time.Minute)) // This should trigger LRU eviction

	_, found := s.Get("key1")
	if found {
		t.Errorf("Expected LRU key 'key1' to be evicted, but it was found")
	}
}
