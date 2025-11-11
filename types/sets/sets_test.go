package sets

import (
	"testing"
)

func TestSadd(t *testing.T) {
	key := "TestSadd"
	values := []string{"value1", "value2", "value3"}

	count := Sadd(key, values)
	if count != 3 {
		t.Errorf("Expected 3 new elements added, got %d", count)
	}

	values = append(values, "value4")
	count = Sadd(key, values)
	if count != 1 {
		t.Errorf("Expected 1 new element added, got %d", count)
	}
}

func TestScard(t *testing.T) {
	key := "TestScard"
	values := []string{"value1", "value2", "value3"}

	// Ensure the set is empty initially
	count := Scard(key)
	if count != 0 {
		t.Errorf("Expected set cardinality to be 0, got %d", count)
	}

	// Add elements to the set
	Sadd(key, values)

	// Check the cardinality again
	count = Scard(key)
	if count != 3 {
		t.Errorf("Expected set cardinality to be 3, got %d", count)
	}
}

func TestSdiff(t *testing.T) {
	key1 := "TestSdiff1"
	key2 := "TestSdiff2"
	values1 := []string{"value1", "value2", "value3"}
	values2 := []string{"value2", "value4"}

	// Add elements to both sets
	Sadd(key1, values1)
	Sadd(key2, values2)

	// Compute the difference
	diffValues := Sdiff([]string{key1, key2})
	expectedDiff := map[string]bool{"value1": true, "value3": true}

	if len(diffValues) != len(expectedDiff) {
		t.Errorf("Expected %d elements in difference, got %d", len(expectedDiff), len(diffValues))
	}

	for _, val := range diffValues {
		if !expectedDiff[val] {
			t.Errorf("Unexpected value in difference: %s", val)
		}
	}
}

func TestSdiffNonExistingSet(t *testing.T) {
	key1 := "NonExistingSet1"
	key2 := "NonExistingSet2"

	// Compute the difference between two non-existing sets
	diffValues := Sdiff([]string{key1, key2})
	if len(diffValues) != 0 {
		t.Errorf("Expected 0 elements in difference for non-existing sets, got %d", len(diffValues))
	}
}

func TestSdiffNoKeys(t *testing.T) {

	// Compute the difference between two non-existing sets
	diffValues := Sdiff([]string{})
	if len(diffValues) != 0 {
		t.Errorf("Expected 0 elements in difference for non-existing sets, got %d", len(diffValues))
	}
}

func TestSdiffWithEmptySet(t *testing.T) {
	key1 := "TestSdiffWithEmptySet1"
	values1 := []string{"value1", "value2", "value3"}

	// Add elements to the first set
	Sadd(key1, values1)

	// Compute the difference with an empty set
	diffValues := Sdiff([]string{key1, "EmptySet"})
	expectedDiff := map[string]bool{"value1": true, "value2": true, "value3": true}

	if len(diffValues) != len(expectedDiff) {
		t.Errorf("Expected %d elements in difference, got %d", len(expectedDiff), len(diffValues))
	}

	for _, val := range diffValues {
		if !expectedDiff[val] {
			t.Errorf("Unexpected value in difference: %s", val)
		}
	}
}

func TestSismember(t *testing.T) {
	key := "TestSismember"
	values := []string{"value1", "value2", "value3"}

	// Add elements to the set
	Sadd(key, values)

	// Test membership for existing elements
	for _, val := range values {
		isMember := Sismember(key, val)
		if !isMember {
			t.Errorf("Expected %s to be a member of the set", val)
		}
	}

	// Test membership for a non-existing element
	nonMember := "value4"
	isMember := Sismember(key, nonMember)
	if isMember {
		t.Errorf("Expected %s to not be a member of the set", nonMember)
	}
}

func TestSismemberEmptySet(t *testing.T) {
	key := "EmptySet"
	value := "somevalue"

	isMember := Sismember(key, value)
	if isMember {
		t.Errorf("Expected %s to not be a member of the empty set", value)
	}
}
