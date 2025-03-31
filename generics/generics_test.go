package generics

import (
	"testing"

	"github.com/divy-sh/animus/essentias"
)

func TestStringCopy(t *testing.T) {
	essentias.Set("TestStringCopy", "expected")
	Copy("TestStringCopy", "TestStringCopy2")
	val, err := essentias.Get("TestStringCopy2")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestHashCopy(t *testing.T) {
	essentias.HSet("TestStringCopy", "pizza", "expected")
	Copy("TestStringCopy", "TestStringCopy2")
	val, err := essentias.HGet("TestStringCopy2", "pizza")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestListCopy(t *testing.T) {
	essentias.RPush("TestStringCopy", &[]string{"expected"})
	Copy("TestStringCopy", "TestStringCopy2")
	val, err := essentias.RPop("TestStringCopy2", "1")
	if err != nil || val[0] != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}
