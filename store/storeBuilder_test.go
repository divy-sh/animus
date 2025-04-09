package store

import (
	"sync"
	"testing"
)

// Reset global variables between tests
func resetGlobals() {
	once = sync.Once{}
	storeConfig = StoreConfig{}
	sharedInstance = nil
}

// Test that a new builder has default config
func TestNewStoreBuilder(t *testing.T) {
	resetGlobals()

	builder := NewStoreBuilder()
	if builder == nil {
		t.Fatal("Expected builder instance, got nil")
	}

	if builder.config.MaxSize != 0 {
		t.Errorf("Expected default MaxSize of 0, got %d", builder.config.MaxSize)
	}
}

// Test SetMaxSize updates the config
func TestSetMaxSize(t *testing.T) {
	resetGlobals()

	builder := NewStoreBuilder().SetMaxSize(42)
	if builder.config.MaxSize != 42 {
		t.Errorf("Expected MaxSize to be 42, got %d", builder.config.MaxSize)
	}
}

// Test that Build applies config and initializes Store only once
func TestBuildInitializesStore(t *testing.T) {
	resetGlobals()

	builder := NewStoreBuilder().SetMaxSize(100)
	builder.Build()

	if sharedInstance == nil {
		t.Fatal("Expected sharedInstance to be initialized")
	}

	if sharedInstance.maxSize != 100 {
		t.Errorf("Expected maxSize to be 100, got %d", sharedInstance.maxSize)
	}

	// Try to re-build with different config — should not change
	builder2 := NewStoreBuilder().SetMaxSize(200)
	builder2.Build()

	if sharedInstance.maxSize != 100 {
		t.Errorf("Expected maxSize to remain 100 after second Build, got %d", sharedInstance.maxSize)
	}
}
