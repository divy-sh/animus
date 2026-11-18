package lists

import (
	"testing"
	"time"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
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

func Test_LMove(t *testing.T) {
	sourceKey := "sourceList"
	destKey := "destList"
	values := []string{"a", "b", "c", "d"}
	RPush(sourceKey, &values)

	movedValue, err := Lmove(sourceKey, destKey, "RIGHT")
	if err != nil {
		t.Errorf("Expected no error for valid LMove")
	}
	if movedValue != "d" {
		t.Errorf("Expected moved value 'd', got '%s'", movedValue)
	}

	newSourceLen, _ := LLen(sourceKey)
	if newSourceLen != 3 {
		t.Errorf("Expected source list length 3, got %d", newSourceLen)
	}

	newDestLen, _ := LLen(destKey)
	if newDestLen != 1 {
		t.Errorf("Expected destination list length 1, got %d", newDestLen)
	}

	destValue, _ := Lindex(destKey, 0)
	if destValue != "d" {
		t.Errorf("Expected destination list first element 'd', got '%s'", destValue)
	}
}

func Test_LMove_EmptySource(t *testing.T) {
	sourceKey := "emptySourceList"
	destKey := "destList"

	_, err := Lmove(sourceKey, destKey, "RIGHT")
	if err == nil {
		t.Errorf("Expected error for LMove from empty source list")
	}
}

func Test_RPushX(t *testing.T) {
	key := "testRPushX"
	values := []string{"a", "b"}
	RPush(key, &values)

	err := RPushx(key, &[]string{"d", "e"})
	if err != nil {
		t.Errorf("Expected no error for valid RPushX")
	}

	length, _ := LLen(key)
	if length != 4 {
		t.Errorf("Expected list length 3 after RPushX, got %d", length)
	}

	err = RPushx("nonExistentKey", &[]string{"d", "e"})
	if err == nil {
		t.Errorf("Expected error for RPushX on non-existent key")
	}
}

func Test_BLmoveDirections(t *testing.T) {
	tests := []struct {
		name                 string
		sourceDirection      string
		destinationDirection string
		moved                string
		expectedDestination  []string
	}{
		{"left to left", "LEFT", "LEFT", "a", []string{"a", "x"}},
		{"left to right", "LEFT", "RIGHT", "a", []string{"x", "a"}},
		{"right to left", "RIGHT", "LEFT", "b", []string{"b", "x"}},
		{"right to right", "RIGHT", "RIGHT", "b", []string{"x", "b"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sourceKey := "blmove-source-" + t.Name()
			destinationKey := "blmove-destination-" + t.Name()
			t.Cleanup(func() {
				store.Delete(sourceKey)
				store.Delete(destinationKey)
			})

			RPush(sourceKey, &[]string{"a", "b"})
			RPush(destinationKey, &[]string{"x"})

			moved, ok, err := BLmove(sourceKey, destinationKey, tc.sourceDirection, tc.destinationDirection, 1)
			if err != nil {
				t.Fatalf("BLmove returned unexpected error: %v", err)
			}
			if !ok || moved != tc.moved {
				t.Fatalf("BLmove returned (%q, %t), want (%q, true)", moved, ok, tc.moved)
			}

			destination, err := Lrange(destinationKey, 0, -1)
			if err != nil {
				t.Fatalf("Lrange destination: %v", err)
			}
			if len(destination) != len(tc.expectedDestination) {
				t.Fatalf("destination = %v, want %v", destination, tc.expectedDestination)
			}
			for i := range destination {
				if destination[i] != tc.expectedDestination[i] {
					t.Fatalf("destination = %v, want %v", destination, tc.expectedDestination)
				}
			}
		})
	}
}

func Test_BLmoveBlocksUntilSourceHasValue(t *testing.T) {
	sourceKey := "blmove-blocking-source"
	destinationKey := "blmove-blocking-destination"
	t.Cleanup(func() {
		store.Delete(sourceKey)
		store.Delete(destinationKey)
	})

	result := make(chan struct {
		value string
		ok    bool
		err   error
	}, 1)
	go func() {
		value, ok, err := BLmove(sourceKey, destinationKey, "LEFT", "RIGHT", 1)
		result <- struct {
			value string
			ok    bool
			err   error
		}{value, ok, err}
	}()

	time.Sleep(20 * time.Millisecond)
	RPush(sourceKey, &[]string{"ready"})

	select {
	case got := <-result:
		if got.err != nil || !got.ok || got.value != "ready" {
			t.Fatalf("BLmove returned (%q, %t, %v), want (ready, true, nil)", got.value, got.ok, got.err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("BLmove did not unblock after a value was pushed")
	}
}

func Test_BLmoveAllowsSameSourceAndDestination(t *testing.T) {
	key := "blmove-same-key"
	t.Cleanup(func() {
		store.Delete(key)
	})
	RPush(key, &[]string{"a", "b"})

	moved, ok, err := BLmove(key, key, "LEFT", "RIGHT", 1)
	if err != nil || !ok || moved != "a" {
		t.Fatalf("BLmove returned (%q, %t, %v), want (a, true, nil)", moved, ok, err)
	}
	values, err := Lrange(key, 0, -1)
	if err != nil {
		t.Fatalf("Lrange: %v", err)
	}
	if len(values) != 2 || values[0] != "b" || values[1] != "a" {
		t.Fatalf("list after same-key move = %v, want [b a]", values)
	}
}

func Test_BLmoveTimeoutAndInvalidArguments(t *testing.T) {
	sourceKey := "blmove-timeout-source"
	destinationKey := "blmove-timeout-destination"
	t.Cleanup(func() {
		store.Delete(sourceKey)
		store.Delete(destinationKey)
	})

	_, moved, err := BLmove(sourceKey, destinationKey, "LEFT", "RIGHT", 0.02)
	if err != nil || moved {
		t.Fatalf("BLmove on an empty source returned moved=%t, err=%v; want false, nil", moved, err)
	}
	if _, err := LLen(destinationKey); err == nil {
		t.Fatal("BLmove created the destination after timing out")
	}

	if _, _, err := BLmove(sourceKey, destinationKey, "UP", "RIGHT", 0); err == nil {
		t.Fatal("BLmove accepted an invalid source direction")
	}
	if _, _, err := BLmove(sourceKey, destinationKey, "LEFT", "RIGHT", -1); err == nil {
		t.Fatal("BLmove accepted a negative timeout")
	}
}
