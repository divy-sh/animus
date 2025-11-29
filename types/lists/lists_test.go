package lists

import (
	"testing"

	"github.com/divy-sh/animus/common"
)

func TestRPush(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	popped, err := RPop(key, "4")
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
	RPush(key, &values)

	popped, err := RPop(key, "2")
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
	RPush(key, &values)

	_, err := RPop(key, "10")
	if err == nil {
		t.Errorf("Expected error for invalid count")
	}
}

func TestRPopInvalidCountNegative(t *testing.T) {
	key := "testList"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	_, err := RPop(key, "-1")
	if err == nil {
		t.Errorf("Expected error for negative count")
	}
}

func TestRPopNonExistentKey(t *testing.T) {
	_, err := RPop("nonExistentKey", "1")
	if err == nil {
		t.Errorf("Expected error for non-existent key")
	}
}

func TestLindex(t *testing.T) {
	key := "testListIndex"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	val, err := Lindex(key, 2)
	if err != nil {
		t.Errorf("Expected no error for valid LIndex")
	}
	if val != "c" {
		t.Errorf("Expected 'c', got '%s'", val)
	}
}

func Test_LindexUnderflow(t *testing.T) {
	key := "testListIndexUnderflow"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	val, err := Lindex(key, -1)
	if err != nil {
		t.Errorf("Expected no error for valid LIndex")
	}
	if val != "d" {
		t.Errorf("Expected 'd', got '%s'", val)
	}
}

func Test_LindexOverflow(t *testing.T) {
	key := "testListIndexOverflow"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	_, err := Lindex(key, 10)
	if err == nil {
		t.Errorf("Expected error for out-of-bounds index")
	}
}

func Test_Lindex_NoKey(t *testing.T) {
	_, err := Lindex("nonExistentKey", 0)
	if err == nil {
		t.Errorf("Expected error for non-existent key")
	}
}

func Test_Llen(t *testing.T) {
	key := "testListLen"
	values := []string{"a", "b", "c", "d"}
	RPush(key, &values)

	length, err := LLen(key)
	if err != nil {
		t.Errorf("Expected no error for valid LLen")
	}
	if length != 4 {
		t.Errorf("Expected length 4, got %d", length)
	}
}

func Test_Llen_NoKey(t *testing.T) {
	_, err := LLen("nonExistentKey")
	if err == nil || err.Error() != common.ERR_LIST_NOT_FOUND {
		t.Errorf("Expected error for non-existent key, got %v", err)
	}
}

func Test_LInsert(t *testing.T) {
	tests := map[string]struct {
		key        string
		pivot      string
		value      string
		position   string
		expectErr  bool
		finalIndex int
		finalValue string
	}{
		"valid before": {
			key:        "testList1",
			pivot:      "b",
			value:      "x",
			position:   "BEFORE",
			expectErr:  false,
			finalIndex: 1,
			finalValue: "x",
		},
		"valid after": {
			key:        "testList2",
			pivot:      "b",
			value:      "y",
			position:   "AFTER",
			expectErr:  false,
			finalIndex: 2,
			finalValue: "y",
		},
		// "pivot not found": {
		// 	key:       "testList3",
		// 	pivot:     "z",
		// 	value:     "w",
		// 	position:  "BEFORE",
		// 	expectErr: false,
		// },
		"invalid position": {
			key:       "testList4",
			pivot:     "b",
			value:     "v",
			position:  "MIDDLE",
			expectErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			values := []string{"a", "b", "c", "d"}
			RPush(tc.key, &values)

			_, err := Linsert(tc.key, tc.position, tc.pivot, tc.value)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			if tc.finalIndex >= 0 {
				val, err := Lindex(tc.key, int64(tc.finalIndex))
				if err != nil {
					t.Errorf("Expected no error for final LIndex")
				}
				if val != tc.finalValue {
					t.Errorf("Expected '%s' at index %d, got '%s'", tc.finalValue, tc.finalIndex, val)
				}
			}
		})
	}
}
