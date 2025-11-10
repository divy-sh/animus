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
