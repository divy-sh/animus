package essentias_test

import (
	"testing"

	"github.com/divy-sh/animus/essentias"
)

func TestRPush(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	essentias.RPush(key, &values)

	popped, err := essentias.RPop(key, "4")
	if err != nil {
		t.Errorf("Expected no error for valid RPop")
	}
	if len(popped) != len(values) {
		t.Errorf("Expected popped values to match pushed values")
	}
}

func TestRPopValid(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	essentias.RPush(key, &values)

	popped, err := essentias.RPop(key, "2")
	if err != nil {
		t.Errorf("Expected no error for valid RPop")
	}
	expected := []string{"b", "c"}
	if len(popped) != len(expected) {
		t.Errorf("Expected popped values to match")
	}
}

func TestRPopInvalidCountHigh(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	essentias.RPush(key, &values)

	_, err := essentias.RPop(key, "10")
	if err == nil {
		t.Errorf("Expected error for invalid count")
	}
}

func TestRPopInvalidCountNegative(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	essentias.RPush(key, &values)

	_, err := essentias.RPop(key, "-1")
	if err == nil {
		t.Errorf("Expected error for negative count")
	}
}

func TestRPopNonExistentKey(t *testing.T) {
	_, err := essentias.RPop("nonExistentKey", "1")
	if err == nil {
		t.Errorf("Expected error for non-existent key")
	}
}
