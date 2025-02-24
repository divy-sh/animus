package types

import (
	"testing"
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
