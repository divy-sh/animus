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
	essentias.HSet("TestHashCopy", "pizza", "expected")
	Copy("TestHashCopy", "TestHashCopy2")
	val, err := essentias.HGet("TestHashCopy2", "pizza")
	if err != nil || val != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestListCopy(t *testing.T) {
	essentias.RPush("TestListCopy", &[]string{"expected"})
	Copy("TestListCopy", "TestListCopy2")
	val, err := essentias.RPop("TestListCopy2", "1")
	if err != nil || val[0] != "expected" {
		t.Errorf("Expected value: expected, got: %v", val)
	}
}

func TestStringDelete(t *testing.T) {
	essentias.Set("TestStringDelete", "expected")
	Delete([]string{"TestStringDelete"})
	val, err := essentias.Get("TestStringDelete")
	if err == nil || err.Error() != "ERR string does not exist" {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestHashDelete(t *testing.T) {
	essentias.HSet("TestHashDelete", "pizza", "expected")
	Delete([]string{"TestHashDelete"})
	val, err := essentias.HGet("TestHashDelete", "pizza")
	if err == nil || err.Error() != "ERR hash does not exist" {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}

func TestListDelete(t *testing.T) {
	essentias.RPush("TestListDelete", &[]string{"expected"})
	Delete([]string{"TestListDelete"})
	val, err := essentias.RPop("TestListDelete", "1")
	if err == nil || err.Error() != "ERR list does not exist" {
		t.Errorf("Expected the key to be deleted, got: %v, %v", val, err)
	}
}
