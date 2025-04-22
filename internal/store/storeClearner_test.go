package store

import (
	"fmt"
	"testing"
	"time"
)

func TestExpiryCleaner(t *testing.T) {
	// Clear store before testing
	store.mutex.Lock()
	store.LRUCache.Purge()
	store.mutex.Unlock()

	// Set values: 2 active, 3 expired
	SetWithTTL("active1", "value1", 60)   // TTL in 60 seconds
	SetWithTTL("active2", "value2", 120)  // TTL in 120 seconds
	SetWithTTL("expired1", "value3", -10) // Already expired
	SetWithTTL("expired2", "value4", -20)
	SetWithTTL("expired3", "value5", -30)

	// Give the cleaner some time to run
	time.Sleep(20 * time.Millisecond)

	// Check expired keys are removed
	if _, ok := Get[string, string]("expired1"); ok {
		t.Error("expired1 should have been removed")
	}
	if _, ok := Get[string, string]("expired2"); ok {
		t.Error("expired2 should have been removed")
	}
	if _, ok := Get[string, string]("expired3"); ok {
		t.Error("expired3 should have been removed")
	}

	// Check active keys are still there
	if val, ok := Get[string, string]("active1"); !ok || val != "value1" {
		t.Error("active1 should still be present")
	}
	if val, ok := Get[string, string]("active2"); !ok || val != "value2" {
		t.Error("active2 should still be present")
	}
}

func TestExpiryCleanerAutoWithSampling(t *testing.T) {
	// Clear store before testing
	store.mutex.Lock()
	store.LRUCache.Purge()
	store.mutex.Unlock()

	for i := 0; i < 100; i++ {
		SetWithTTL(fmt.Sprintf("active%d", i), "value", int64(i%2))
	}

	time.Sleep(2 * time.Second)
	cleanExpiredKeys()
	if store.LRUCache.Len() > 0 {
		t.Errorf("cleanExpiredKeys shoul've cleared the expired keys")
	}
}

func TestCleanerStartStop(t *testing.T) {
	StopExpiryCleaner()
	if store.isRunning {
		t.Error("Cleaner should have stopped")
	}

	StartExpiryCleaner()
	if !store.isRunning {
		t.Error("Cleaner should have started")
	}

	StopExpiryCleaner()
}
