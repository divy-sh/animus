package types

import (
	"errors"
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	strType := NewStringType()
	strType.Set("key1", "value1")

	val, err := strType.Get("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
}

func TestAppend(t *testing.T) {
	strType := NewStringType()
	strType.Set("key1", "value1")
	strType.Append("key1", "_appended")

	val, err := strType.Get("key1")
	if err != nil || val != "value1_appended" {
		t.Errorf("Expected value1_appended, got %v, err: %v", val, err)
	}
}

func TestGetDel(t *testing.T) {
	strType := NewStringType()
	strType.Set("key1", "value1")

	val, err := strType.GetDel("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}

	_, err = strType.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetEx(t *testing.T) {
	strType := NewStringType()
	strType.Set("key1", "value1")

	val, err := strType.GetEx("key1", "0")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
	time.Sleep(50 * time.Millisecond)
	_, err = strType.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetRange(t *testing.T) {
	// Create a StringType instance and populate it
	strType := NewStringType()
	strType.Set("hello", "Hello, World!")
	strType.Set("empty", "")

	tests := []struct {
		key      string
		start    string
		end      string
		expected string
		err      error
	}{
		{"hello", "0", "4", "Hello", nil},                                      // Normal range
		{"hello", "-6", "-1", "World!", nil},                                   // Negative index wraparound
		{"hello", "7", "11", "World", nil},                                     // Extract "World"
		{"hello", "0", "50", "Hello, World", nil},                              // End index exceeds length
		{"hello", "-50", "4", "llo", nil},                                      // Start index negative out of bounds
		{"unknown", "0", "2", "", errors.New("ERR key not found, or expired")}, // Key not found
	}

	for _, tt := range tests {
		result, err := strType.GetRange(tt.key, tt.start, tt.end)

		// Check error equality
		if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("GetRange(%q, %q, %q) error = %v, expected %v", tt.key, tt.start, tt.end, err, tt.err)
		}

		// Check expected result
		if result != tt.expected {
			t.Errorf("GetRange(%q, %q, %q) = %q, expected %q", tt.key, tt.start, tt.end, result, tt.expected)
		}
	}
}

func TestDecr(t *testing.T) {
	strType := NewStringType()
	strType.Set("num", "10")

	err := strType.Decr("num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strType.Get("num")
	if val != "9" {
		t.Errorf("Expected 9, got %v", val)
	}
}

func TestDecrBy(t *testing.T) {
	strType := NewStringType()
	strType.Set("num", "10")

	err := strType.DecrBy("num", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strType.Get("num")
	if val != "7" {
		t.Errorf("Expected 7, got %v", val)
	}
}

func TestDecrByInvalid(t *testing.T) {
	strType := NewStringType()
	strType.Set("num", "10")

	err := strType.DecrBy("num", "invalid")
	if err == nil || err.Error() != "ERR invalid decrement value" {
		t.Errorf("Expected error for invalid decrement, got: %v", err)
	}
}
