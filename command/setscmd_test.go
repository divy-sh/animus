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

func TestSDiffStore(t *testing.T) {
	destKey := "TestSdiffStoreDest"
	key1 := "TestSdiffStore1"
	key2 := "TestSdiffStore2"

	// Add elements to the first set
	addArgs1 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs1)

	// Add elements to the second set
	addArgs2 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key2},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
		{Typ: common.BULK_TYPE, Bulk: "elem3"},
	}
	Sadd(addArgs2)

	// Test set difference store
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: destKey},
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	}
	result := SdiffStore(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 1 {
		t.Errorf("Expected 1 new element in destination set, got %d", result.Num)
	}

	// Verify the contents of the destination set
	diffResult := Sdiff([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	})
	if diffResult.Typ != common.ARRAY_TYPE {
		t.Errorf("Expected ARRAY_TYPE, got %v", diffResult.Typ)
	}
	if len(diffResult.Array) != 1 {
		t.Errorf("Expected 0 elements in destination set, got %d", len(diffResult.Array))
	}
}

func TestSDiffStoreInvalidArgs(t *testing.T) {
	// Test with insufficient arguments
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "onlydest"}}
	result := SdiffStore(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for insufficient arguments, got %v", result.Typ)
	}
}

func TestSDiffStoreNoDifference(t *testing.T) {
	destKey := "TestSdiffStoreNoDiffDest"
	key1 := "TestSdiffStoreNoDiff1"
	key2 := "TestSdiffStoreNoDiff2"

	// Add identical elements to both sets
	addArgs1 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs1)

	addArgs2 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key2},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs2)

	// Test set difference store
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: destKey},
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	}
	result := SdiffStore(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 0 {
		t.Errorf("Expected 0 new elements in destination set, got %d", result.Num)
	}

	// Verify the contents of the destination set
	diffResult := Sdiff([]resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	})
	if diffResult.Typ != common.ARRAY_TYPE {
		t.Errorf("Expected ARRAY_TYPE, got %v", diffResult.Typ)
	}
	if len(diffResult.Array) != 0 {
		t.Errorf("Expected 0 elements in destination set, got %d", len(diffResult.Array))
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

func TestSdiff(t *testing.T) {
	key1 := "TestSdiff1"
	key2 := "TestSdiff2"

	// Add elements to the first set
	addArgs1 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs1)

	// Add elements to the second set
	addArgs2 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key2},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
		{Typ: common.BULK_TYPE, Bulk: "elem3"},
	}
	Sadd(addArgs2)

	// Test set difference
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	}
	result := Sdiff(args)
	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("Expected ARRAY_TYPE, got %v", result.Typ)
	}
	expectedDiff := map[string]bool{"elem1": true}
	if len(result.Array) != len(expectedDiff) {
		t.Errorf("Expected %d elements in difference, got %d", len(expectedDiff), len(result.Array))
	}
	for _, val := range result.Array {
		if !expectedDiff[val.Bulk] {
			t.Errorf("Unexpected element in difference: %s", val.Bulk)
		}
	}
}

func TestSdiffInvalidArgs(t *testing.T) {
	// Test with insufficient arguments
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "onlykey"}}
	result := Sdiff(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for insufficient arguments, got %v", result.Typ)
	}
}

func TestSdiffNoDifference(t *testing.T) {
	key1 := "TestSdiffNoDiff1"
	key2 := "TestSdiffNoDiff2"

	// Add identical elements to both sets
	addArgs1 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs1)

	addArgs2 := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key2},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs2)

	// Test set difference
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key1},
		{Typ: common.BULK_TYPE, Bulk: key2},
	}
	result := Sdiff(args)
	if result.Typ != common.ARRAY_TYPE {
		t.Errorf("Expected ARRAY_TYPE, got %v", result.Typ)
	}
	if len(result.Array) != 0 {
		t.Errorf("Expected 0 elements in difference, got %d", len(result.Array))
	}
}

func TestSismember(t *testing.T) {
	key := "TestSismember"

	// Add elements to the set
	addArgs := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
		{Typ: common.BULK_TYPE, Bulk: "elem2"},
	}
	Sadd(addArgs)

	// Test membership for an existing element
	args := []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: key},
		{Typ: common.BULK_TYPE, Bulk: "elem1"},
	}
	result := Sismember(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 1 {
		t.Errorf("Expected membership 1, got %d", result.Num)
	}

	// Test membership for a non-existing element
	args[1] = resp.Value{Typ: common.BULK_TYPE, Bulk: "elem3"}
	result = Sismember(args)
	if result.Typ != common.INTEGER_TYPE {
		t.Errorf("Expected INTEGER_TYPE, got %v", result.Typ)
	}
	if result.Num != 0 {
		t.Errorf("Expected membership 0, got %d", result.Num)
	}
}

func TestSismemberInvalidArgs(t *testing.T) {
	// Test with insufficient arguments
	args := []resp.Value{{Typ: common.BULK_TYPE, Bulk: "onlykey"}}
	result := Sismember(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for insufficient arguments, got %v", result.Typ)
	}
	// Test with too many arguments
	args = []resp.Value{
		{Typ: common.BULK_TYPE, Bulk: "key"},
		{Typ: common.BULK_TYPE, Bulk: "value1"},
		{Typ: common.BULK_TYPE, Bulk: "value2"},
	}
	result = Sismember(args)
	if result.Typ != common.ERROR_TYPE {
		t.Errorf("Expected ERROR_TYPE for too many arguments, got %v", result.Typ)
	}
}
