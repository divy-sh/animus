package command

import (
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/resp"
)

func TestSadd(t *testing.T) {
	key := "TestSadd"
	values := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
		{Typ: common.BULK_TYPE, Bulk: "elem3"},
	}

	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: key}}
	args = append(args, values...)

	result := Sadd(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 3 {
		t.Errorf("Expected 3 new elements added, got %d", result.Num)
	}
	args = append(args, resp.Value{Typ: common.BULK_TYPE, Bulk: "elem4"})

	// Test adding the same elements again
	result = Sadd(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 1 {
		t.Errorf("Expected 1 new element added, got %d", result.Num)
	}
}

func TestSaddInvalidArgs(t *testing.T) {
	// Test with insufficient arguments
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "onlykey"}}
	result := Sadd(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for insufficient arguments, got %v", result.Typ)
	}
}

func TestScard(t *testing.T) {
	key := "TestScard"

	// Test cardinality of an empty set
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: key}}
	result := Scard(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 0 {
		t.Errorf("Expected cardinality 0, got %d", result.Num)
	}

	// Add elements to the set
	addArgs := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs)

	// Test cardinality after adding elements
	result = Scard(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 2 {
		t.Errorf("Expected cardinality 2, got %d", result.Num)
	}
}

func TestScardInvalidArgs(t *testing.T) {
	// Test with insufficient arguments
	args := []resp.Value{}
	result := Scard(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for insufficient arguments, got %v", result.Typ)
	}
}
