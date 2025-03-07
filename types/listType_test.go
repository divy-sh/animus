package types_test

import (
	"testing"

	"github.com/divy-sh/animus/types"
)

func TestRPush(t *testing.T) {
	lt := types.NewListType()
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	lt.RPush(key, &values)

	popped, err := lt.RPop(key, "4")
	if err != nil {
		t.Errorf("Expected no error for valid RPop")
	}
	if len(popped) != len(values) {
		t.Errorf("Expected popped values to match pushed values")
	}
}

func TestRPopValid(t *testing.T) {
	lt := types.NewListType()
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	lt.RPush(key, &values)

	popped, err := lt.RPop(key, "2")
	if err != nil {
		t.Errorf("Expected no error for valid RPop")
	}
	expected := []string{"b", "c"}
	if len(popped) != len(expected) {
		t.Errorf("Expected popped values to match")
	}
}

func TestRPopInvalidCountHigh(t *testing.T) {
	lt := types.NewListType()
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	lt.RPush(key, &values)

	_, err := lt.RPop(key, "10")
	if err == nil {
		t.Errorf("Expected error for invalid count")
	}
}

func TestRPopInvalidCountNegative(t *testing.T) {
	lt := types.NewListType()
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	lt.RPush(key, &values)

	_, err := lt.RPop(key, "-1")
	if err == nil {
		t.Errorf("Expected error for negative count")
	}
}

func TestRPopNonExistentKey(t *testing.T) {
	lt := types.NewListType()
	_, err := lt.RPop("nonExistentKey", "1")
	if err == nil {
		t.Errorf("Expected error for non-existent key")
	}
}
