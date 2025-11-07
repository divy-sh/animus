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
