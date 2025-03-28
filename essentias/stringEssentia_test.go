package essentias_test

import (
	"errors"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/divy-sh/animus/essentias"
)

const float64EqualityThreshold = 1e-9

func floatsAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestSetAndGet(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")

	val, err := strEssentia.Get("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
}

func TestAppend(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")
	strEssentia.Append("key1", "_appended")

	val, err := strEssentia.Get("key1")
	if err != nil || val != "value1_appended" {
		t.Errorf("Expected value1_appended, got %v, err: %v", val, err)
	}
}

func TestAppendNewKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Append("TestAppendNewKey", "appended")

	val, err := strEssentia.Get("TestAppendNewKey")
	if err != nil || val != "appended" {
		t.Errorf("Expected appended, got %v, err: %v", val, err)
	}
}

func TestGetDel(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")

	val, err := strEssentia.GetDel("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}

	_, err = strEssentia.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetDelInvalidKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()

	_, err := strEssentia.GetDel("invalid_key")
	if err == nil || err.Error() != "ERR key not found, or expired" {
		t.Errorf("Expected error: %v, got: %v", "ERR key not found, or expired", err)
	}
}

func TestGetEx(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")

	val, err := strEssentia.GetEx("key1", "0")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
	time.Sleep(50 * time.Millisecond)
	_, err = strEssentia.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetExInvalidKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()

	_, err := strEssentia.GetEx("invalid_key", "1")
	if err == nil || err.Error() != "ERR key not found, or expired" {
		t.Errorf("Expected error: %v, got: %v", "ERR key not found, or expired", err)
	}
}

func TestGetExInvalidTime(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")

	_, err := strEssentia.GetEx("key1", "invalid")
	if err == nil || err.Error() != "ERR invalid expire time" {
		t.Errorf("Expected error: %v, got: %v", "ERR invalid expire time", err)
	}
}

func TestGetRange(t *testing.T) {
	// Create a StringEssentia instance and populate it
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("hello", "Hello, World!")
	strEssentia.Set("empty", "")

	tests := []struct {
		key      string
		start    string
		end      string
		expected string
		err      error
	}{
		{"hello", "0", "4", "Hello", nil},                                      // Normal range
		{"hello", "-6", "-1", "World!", nil},                                   // Negative index wraparound
		{"hello", "7", "11", "World", nil},                                     // Extract "World"
		{"hello", "0", "50", "Hello, World", nil},                              // End index exceeds length
		{"hello", "-50", "4", "llo", nil},                                      // Start index negative out of bounds
		{"unknown", "0", "2", "", errors.New("ERR key not found, or expired")}, // Key not found
	}

	for _, tt := range tests {
		result, err := strEssentia.GetRange(tt.key, tt.start, tt.end)

		// Check error equality
		if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) ||
			(err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("GetRange(%q, %q, %q) error = %v, expected %v", tt.key, tt.start, tt.end, err, tt.err)
		}

		// Check expected result
		if result != tt.expected {
			t.Errorf("GetRange(%q, %q, %q) = %q, expected %q", tt.key, tt.start, tt.end, result, tt.expected)
		}
	}
}

func TestGetRangeInvalidStartIndex(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("hello", "Hello, World!")
	_, err := strEssentia.GetRange("hello", "invalid_start_index", "20")
	expected := "ERR invalid start index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeInvalidEndIndex(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("hello", "Hello, World!")
	_, err := strEssentia.GetRange("hello", "0", "invalid_end_index")
	expected := "ERR invalid end index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeStartIndexLargerThanEndIndex(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("hello", "Hello, World!")
	_, err := strEssentia.GetRange("hello", "6", "5")
	expected := "ERR start index greater than end index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeEmptyVal(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("hello", "")
	val, err := strEssentia.GetRange("hello", "0", "20")
	expected := ""
	if err != nil || val != expected {
		t.Errorf("Expected value: %v, got error: %v", expected, err)
	}
}

func TestDecr(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.Decr("num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("num")
	if val != "9" {
		t.Errorf("Expected 9, got %v", val)
	}
}

func TestDecrBy(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.DecrBy("num", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("num")
	if val != "7" {
		t.Errorf("Expected 7, got %v", val)
	}
}

func TestDecrByNewKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()

	err := strEssentia.DecrBy("TestDecrByNewKey", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("TestDecrByNewKey")
	if val != "-3" {
		t.Errorf("Expected -3, got %v", val)
	}
}

func TestDecrByInvalidValue(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "Z")

	err := strEssentia.DecrBy("num", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestDecrByInvalid(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.DecrBy("num", "invalid")
	if err == nil || err.Error() != "ERR invalid decrement value" {
		t.Errorf("Expected error for invalid decrement, got: %v", err)
	}
}

func TestGetSet(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "value1")

	val, err := strEssentia.GetSet("key1", "value2")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}

	val, err = strEssentia.Get("key1")
	if err != nil || val != "value2" {
		t.Errorf("Expected val: value2, got val: %v, error: %v", val, err)
	}

}

func TestGetSetInvalidKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	val, err := strEssentia.GetSet("invalid", "value1")
	expected := "ERR key not found, or expired"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %s, got val: %v, error: %v", expected, val, err)
	}
}

func TestIncr(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.Incr("num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("num")
	if val != "11" {
		t.Errorf("Expected 11, got %v", val)
	}
}

func TestIncrBy(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.IncrBy("num", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("num")
	if val != "13" {
		t.Errorf("Expected 13, got %v", val)
	}
}

func TestIncrByNewKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()

	err := strEssentia.IncrBy("TestIncrByNewKey", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("TestIncrByNewKey")
	if val != "3" {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestIncrByInvalidValue(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "Z")

	err := strEssentia.IncrBy("num", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestIncrByInvalid(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.IncrBy("num", "invalid")
	if err == nil || err.Error() != "ERR invalid increment value" {
		t.Errorf("Expected error for invalid increment, got: %v", err)
	}
}

func TestIncrByFloat(t *testing.T) {
	tests := []struct {
		key      string
		val      string
		incr     string
		expected float64
		err      error
	}{
		{"hello", "19", "4", 23, nil},
		{"hello", "19", "4.4", 23.4, nil},
		{"hello", "19.4", "4", 23.4, nil},
		{"hello", "19.4", "4.4", 23.8, nil},
	}
	for _, tt := range tests {
		strEssentia := essentias.NewStringEssentia()
		strEssentia.Set(tt.key, tt.val)

		err := strEssentia.IncrByFloat(tt.key, tt.incr)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		val, _ := strEssentia.Get(tt.key)
		floatVal, _ := strconv.ParseFloat(val, 64)
		if !floatsAlmostEqual(floatVal, tt.expected) {
			t.Errorf("Expected %f, got %v", tt.expected, val)
		}
	}
}

func TestIncrByNewKeyFloat(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()

	err := strEssentia.IncrByFloat("TestIncrByNewKeyFloat", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strEssentia.Get("TestIncrByNewKeyFloat")
	if val != "3" {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestIncrByInvalidValueFloat(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("TestIncrByInvalidValueFloat", "Z")

	err := strEssentia.IncrByFloat("TestIncrByInvalidValueFloat", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestIncrByInvalidFloat(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("num", "10")

	err := strEssentia.IncrByFloat("num", "invalid")
	if err == nil || err.Error() != "ERR invalid increment value" {
		t.Errorf("Expected error for invalid increment, got: %v", err)
	}
}

func TestLcs(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "lasagna")
	strEssentia.Set("key2", "baigan")
	lcs, err := strEssentia.Lcs("key1", "key2", []string{})
	if err != nil || lcs != "aga" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsLen(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "lasagna")
	strEssentia.Set("key2", "baigan")
	lcsLen, err := strEssentia.Lcs("key1", "key2", []string{"len"})
	if err != nil || lcsLen != "3" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsLen2(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "lasagna")
	strEssentia.Set("key2", "baigan")
	lcsLen, err := strEssentia.Lcs("key2", "key1", []string{"len"})
	if err != nil || lcsLen != "3" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsInvalidFirstKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	_, err := strEssentia.Lcs("TestLcsInvalidFirstKey", "key2", []string{})
	expected := "ERR key not found, or expired"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error for lcs: %s, got: %v", expected, err)
	}
}

func TestLcsInvalidSecondKey(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.Set("key1", "lasagna")
	_, err := strEssentia.Lcs("key1", "TestLcsInvalidSecondKey", []string{})
	expected := "ERR key not found, or expired"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error for lcs: %s, got: %v", expected, err)
	}
}

func TestMGetAndMSet(t *testing.T) {
	strEssentia := essentias.NewStringEssentia()
	strEssentia.MSet(&map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})

	values := strEssentia.MGet(&[]string{"key1", "key2", "key3", "invalid"})
	expected := []string{"value1", "value2", "value3", ""}
	for i := range *values {
		if (*values)[i] != expected[i] {
			t.Errorf("Expected values: %v got %v", expected[i], (*values)[i])
		}
	}
}
