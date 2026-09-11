package strings_test

import (
	"errors"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/types/strings"
)

const float64EqualityThreshold = 1e-9

func floatsAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestSetAndGet(t *testing.T) {
	strings.Set("key1", "value1")

	val, err := strings.Get("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
}

func TestAppend(t *testing.T) {
	strings.Set("key1", "value1")
	strings.Append("key1", "_appended")

	val, err := strings.Get("key1")
	if err != nil || val != "value1_appended" {
		t.Errorf("Expected value1_appended, got %v, err: %v", val, err)
	}
}

func TestAppendNewKey(t *testing.T) {
	strings.Append("TestAppendNewKey", "appended")

	val, err := strings.Get("TestAppendNewKey")
	if err != nil || val != "appended" {
		t.Errorf("Expected appended, got %v, err: %v", val, err)
	}
}

func TestGetDel(t *testing.T) {
	strings.Set("key1", "value1")

	val, err := strings.GetDel("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}

	_, err = strings.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetDelInvalidKey(t *testing.T) {

	_, err := strings.GetDel("invalid_key")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %v, got: %v", common.ERR_STRING_NOT_FOUND, err)
	}
}

func TestGetEx(t *testing.T) {
	strings.Set("key1", "value1")

	val, err := strings.GetEx("key1", "0")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
	time.Sleep(50 * time.Millisecond)
	_, err = strings.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestGetExInvalidKey(t *testing.T) {

	_, err := strings.GetEx("invalid_key", "1")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %v, got: %v", common.ERR_STRING_NOT_FOUND, err)
	}
}

func TestGetExInvalidTime(t *testing.T) {
	strings.Set("key1", "value1")

	_, err := strings.GetEx("key1", "invalid")
	if err == nil || err.Error() != "ERR invalid expire time" {
		t.Errorf("Expected error: %v, got: %v", "ERR invalid expire time", err)
	}
}

func TestGetRange(t *testing.T) {
	// Create a StringEssentia instance and populate it
	strings.Set("hello", "Hello, World!")
	strings.Set("empty", "")

	tests := []struct {
		key      string
		start    string
		end      string
		expected string
		err      error
	}{
		{"hello", "0", "4", "Hello", nil},                                  // Normal range
		{"hello", "-6", "-1", "World!", nil},                               // Negative index wraparound
		{"hello", "7", "11", "World", nil},                                 // Extract "World"
		{"hello", "0", "50", "Hello, World", nil},                          // End index exceeds length
		{"hello", "-50", "4", "llo", nil},                                  // Start index negative out of bounds
		{"unknown", "0", "2", "", errors.New(common.ERR_STRING_NOT_FOUND)}, // Key not found
	}

	for _, tt := range tests {
		result, err := strings.GetRange(tt.key, tt.start, tt.end)

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
	strings.Set("hello", "Hello, World!")
	_, err := strings.GetRange("hello", "invalid_start_index", "20")
	expected := "ERR invalid start index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeInvalidEndIndex(t *testing.T) {
	strings.Set("hello", "Hello, World!")
	_, err := strings.GetRange("hello", "0", "invalid_end_index")
	expected := "ERR invalid end index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeStartIndexLargerThanEndIndex(t *testing.T) {
	strings.Set("hello", "Hello, World!")
	_, err := strings.GetRange("hello", "6", "5")
	expected := "ERR start index greater than end index"
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %v, got: %v", expected, err)
	}
}

func TestGetRangeEmptyVal(t *testing.T) {
	strings.Set("hello", "")
	val, err := strings.GetRange("hello", "0", "20")
	expected := ""
	if err != nil || val != expected {
		t.Errorf("Expected value: %v, got error: %v", expected, err)
	}
}

func TestDecr(t *testing.T) {
	strings.Set("num", "10")

	err := strings.Decr("num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("num")
	if val != "9" {
		t.Errorf("Expected 9, got %v", val)
	}
}

func TestDecrBy(t *testing.T) {
	strings.Set("num", "10")

	err := strings.DecrBy("num", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("num")
	if val != "7" {
		t.Errorf("Expected 7, got %v", val)
	}
}

func TestDecrByNewKey(t *testing.T) {

	err := strings.DecrBy("TestDecrByNewKey", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("TestDecrByNewKey")
	if val != "-3" {
		t.Errorf("Expected -3, got %v", val)
	}
}

func TestDecrByInvalidValue(t *testing.T) {
	strings.Set("num", "Z")

	err := strings.DecrBy("num", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestDecrByInvalid(t *testing.T) {
	strings.Set("num", "10")

	err := strings.DecrBy("num", "invalid")
	if err == nil || err.Error() != "ERR invalid decrement value" {
		t.Errorf("Expected error for invalid decrement, got: %v", err)
	}
}

func TestGetSet(t *testing.T) {
	strings.Set("key1", "value1")

	val, err := strings.GetSet("key1", "value2")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}

	val, err = strings.Get("key1")
	if err != nil || val != "value2" {
		t.Errorf("Expected val: value2, got val: %v, error: %v", val, err)
	}

}

func TestGetSetInvalidKey(t *testing.T) {
	val, err := strings.GetSet("invalid", "value1")
	expected := common.ERR_STRING_NOT_FOUND
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error: %s, got val: %v, error: %v", expected, val, err)
	}
}

func TestIncr(t *testing.T) {
	strings.Set("num", "10")

	err := strings.Incr("num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("num")
	if val != "11" {
		t.Errorf("Expected 11, got %v", val)
	}
}

func TestIncrBy(t *testing.T) {
	strings.Set("num", "10")

	err := strings.IncrBy("num", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("num")
	if val != "13" {
		t.Errorf("Expected 13, got %v", val)
	}
}

func TestIncrByNewKey(t *testing.T) {

	err := strings.IncrBy("TestIncrByNewKey", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("TestIncrByNewKey")
	if val != "3" {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestIncrByInvalidValue(t *testing.T) {
	strings.Set("num", "Z")

	err := strings.IncrBy("num", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestIncrByInvalid(t *testing.T) {
	strings.Set("num", "10")

	err := strings.IncrBy("num", "invalid")
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
		strings.Set(tt.key, tt.val)

		err := strings.IncrByFloat(tt.key, tt.incr)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		val, _ := strings.Get(tt.key)
		floatVal, _ := strconv.ParseFloat(val, 64)
		if !floatsAlmostEqual(floatVal, tt.expected) {
			t.Errorf("Expected %f, got %v", tt.expected, val)
		}
	}
}

func TestIncrByNewKeyFloat(t *testing.T) {

	err := strings.IncrByFloat("TestIncrByNewKeyFloat", "3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := strings.Get("TestIncrByNewKeyFloat")
	if val != "3" {
		t.Errorf("Expected 3, got %v", val)
	}
}

func TestIncrByInvalidValueFloat(t *testing.T) {
	strings.Set("TestIncrByInvalidValueFloat", "Z")

	err := strings.IncrByFloat("TestIncrByInvalidValueFloat", "3")
	if err == nil {
		t.Errorf("Expected error for invalid value, got: %v", err)
	}
}

func TestIncrByInvalidFloat(t *testing.T) {
	strings.Set("num", "10")

	err := strings.IncrByFloat("num", "invalid")
	if err == nil || err.Error() != "ERR invalid increment value" {
		t.Errorf("Expected error for invalid increment, got: %v", err)
	}
}

func TestLcs(t *testing.T) {
	strings.Set("key1", "lasagna")
	strings.Set("key2", "baigan")
	lcs, err := strings.Lcs("key1", "key2", []string{})
	if err != nil || lcs != "aga" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsLen(t *testing.T) {
	strings.Set("key1", "lasagna")
	strings.Set("key2", "baigan")
	lcsLen, err := strings.Lcs("key1", "key2", []string{"len"})
	if err != nil || lcsLen != "3" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsLen2(t *testing.T) {
	strings.Set("key1", "lasagna")
	strings.Set("key2", "baigan")
	lcsLen, err := strings.Lcs("key2", "key1", []string{"len"})
	if err != nil || lcsLen != "3" {
		t.Errorf("Expected value for lcs, got: %v", err)
	}
}

func TestLcsInvalidFirstKey(t *testing.T) {
	_, err := strings.Lcs("TestLcsInvalidFirstKey", "key2", []string{})
	expected := common.ERR_STRING_NOT_FOUND
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error for lcs: %s, got: %v", expected, err)
	}
}

func TestLcsInvalidSecondKey(t *testing.T) {
	strings.Set("key1", "lasagna")
	_, err := strings.Lcs("key1", "TestLcsInvalidSecondKey", []string{})
	expected := common.ERR_STRING_NOT_FOUND
	if err == nil || err.Error() != expected {
		t.Errorf("Expected error for lcs: %s, got: %v", expected, err)
	}
}

func TestMGetAndMSet(t *testing.T) {
	strings.MSet(&map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})

	values := strings.MGet(&[]string{"key1", "key2", "key3", "invalid"})
	expected := []string{"value1", "value2", "value3", ""}
	for i := range *values {
		if (*values)[i] != expected[i] {
			t.Errorf("Expected values: %v got %v", expected[i], (*values)[i])
		}
	}
}

func TestSetEx(t *testing.T) {
	strings.SetEx("key1", "value1", "1")

	val, err := strings.Get("key1")
	if err != nil || val != "value1" {
		t.Errorf("Expected value1, got %v, err: %v", val, err)
	}
	time.Sleep(time.Second)
	_, err = strings.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestSetExInvalidExpiryTime(t *testing.T) {
	err := strings.SetEx("TestSetExInvalidExpiryTime", "TestSetExInvalidExpiryTime", "invalid")

	if err == nil || err.Error() != common.ERR_INVALID_TIME_SECONDS {
		t.Errorf("Expected %s, got %v", common.ERR_INVALID_TIME_SECONDS, err)
	}
}

func TestStrLen(t *testing.T) {
	strings.Set("key1", "value1")

	length, err := strings.StrLen("key1")
	if err != nil || length != 6 {
		t.Errorf("Expected length 6, got %v, err: %v", length, err)
	}
}

func TestStrLenInvalidKey(t *testing.T) {
	_, err := strings.StrLen("invalid_key")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %v, got: %v", common.ERR_STRING_NOT_FOUND, err)
	}
}

func TestSetRange(t *testing.T) {
	strings.Set("key1", "Hello World")

	err := strings.SetRange("key1", "5", ",")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, err := strings.Get("key1")
	if err != nil || val != "Hello,World" {
		t.Errorf("Expected 'Hello,World', got '%v', err: %v", val, err)
	}
}

func TestSetRangeNewKey(t *testing.T) {
	err := strings.SetRange("TestSetRangeNewKey", "5", "World")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, err := strings.Get("TestSetRangeNewKey")
	if err != nil || val != "\x00\x00\x00\x00\x00World" {
		t.Errorf("Expected '\\x00\\x00\\x00\\x00\\x00World', got '%v', err: %v", val, err)
	}
}

func TestSetRangeInvalidOffset(t *testing.T) {
	strings.Set("key1", "HelloWorld")

	err := strings.SetRange("key1", "-1", "Test")
	if err == nil || err.Error() != common.ERR_OUT_OF_RANGE {
		t.Errorf("Expected error: %v, got: %v", common.ERR_OUT_OF_RANGE, err)
	}

	err = strings.SetRange("key1", "invalid_offset", "Test")
	if err == nil || err.Error() != common.ERR_OUT_OF_RANGE {
		t.Errorf("Expected error: %v, got: %v", common.ERR_OUT_OF_RANGE, err)
	}
}

func TestSetRangeOffsetExceedsLength(t *testing.T) {
	strings.Set("key1", "Hello")
	err := strings.SetRange("key1", "10", "World")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, err := strings.Get("key1")
	if err != nil || val != "Hello\x00\x00\x00\x00\x00World" {
		t.Errorf("Expected 'Hello\\x00\\x00\\x00\\x00\\x00World', got '%v', err: %v", val, err)
	}
}

func TestSetRangeInvalidKey(t *testing.T) {
	err := strings.SetRange("TestSetRangeInvalidKey", "0", "Test")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, err := strings.Get("TestSetRangeInvalidKey")
	if err != nil || val != "Test" {
		t.Errorf("Expected 'Test', got '%v', err: %v", val, err)
	}
}

func TestDigest(t *testing.T) {
	d1 := strings.Digest("value1")
	d2 := strings.Digest("value1")
	d3 := strings.Digest("value2")

	if d1 != d2 {
		t.Errorf("Expected digest of the same value to be equal, got %v and %v", d1, d2)
	}
	if d1 == d3 {
		t.Errorf("Expected digest of different values to differ, got %v for both", d1)
	}
	if d1 == "" {
		t.Errorf("Expected non-empty digest")
	}
}

func TestDelExNoExpression(t *testing.T) {
	strings.Set("TestDelExNoExpression", "value1")

	err := strings.DelEx("TestDelExNoExpression", "", "")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExNoExpression")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestDelExKeyNotFound(t *testing.T) {
	err := strings.DelEx("TestDelExKeyNotFound", "", "")
	if err == nil || err.Error() != common.ERR_STRING_NOT_FOUND {
		t.Errorf("Expected error: %v, got: %v", common.ERR_STRING_NOT_FOUND, err)
	}
}

func TestDelExInvalidExpression(t *testing.T) {
	strings.Set("TestDelExInvalidExpression", "value1")

	err := strings.DelEx("TestDelExInvalidExpression", "INVALID", "value1")
	if err == nil || err.Error() != "ERR invalid expression" {
		t.Errorf("Expected error: %v, got: %v", "ERR invalid expression", err)
	}

	val, err := strings.Get("TestDelExInvalidExpression")
	if err != nil || val != "value1" {
		t.Errorf("Expected key to remain untouched, got val: %v, err: %v", val, err)
	}
}

func TestDelExIfEqMatch(t *testing.T) {
	strings.Set("TestDelExIfEqMatch", "value1")

	err := strings.DelEx("TestDelExIfEqMatch", "IFEQ", "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExIfEqMatch")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestDelExIfEqNoMatch(t *testing.T) {
	strings.Set("TestDelExIfEqNoMatch", "value1")

	err := strings.DelEx("TestDelExIfEqNoMatch", "IFEQ", "value2")
	if err == nil || err.Error() != "ERR condition not met" {
		t.Errorf("Expected error: %v, got: %v", "ERR condition not met", err)
	}

	val, err := strings.Get("TestDelExIfEqNoMatch")
	if err != nil || val != "value1" {
		t.Errorf("Expected key to remain untouched, got val: %v, err: %v", val, err)
	}
}

func TestDelExIfNeMatch(t *testing.T) {
	strings.Set("TestDelExIfNeMatch", "value1")

	err := strings.DelEx("TestDelExIfNeMatch", "IFNE", "value2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExIfNeMatch")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestDelExIfNeNoMatch(t *testing.T) {
	strings.Set("TestDelExIfNeNoMatch", "value1")

	err := strings.DelEx("TestDelExIfNeNoMatch", "IFNE", "value1")
	if err == nil || err.Error() != "ERR condition not met" {
		t.Errorf("Expected error: %v, got: %v", "ERR condition not met", err)
	}

	val, err := strings.Get("TestDelExIfNeNoMatch")
	if err != nil || val != "value1" {
		t.Errorf("Expected key to remain untouched, got val: %v, err: %v", val, err)
	}
}

func TestDelExIfdEqMatch(t *testing.T) {
	strings.Set("TestDelExIfdEqMatch", "value1")

	err := strings.DelEx("TestDelExIfdEqMatch", "IFDEQ", strings.Digest("value1"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExIfdEqMatch")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestDelExIfdEqNoMatch(t *testing.T) {
	strings.Set("TestDelExIfdEqNoMatch", "value1")

	// Passing the raw value instead of its digest must not match, proving
	// that IFDEQ compares hash digests rather than raw values.
	err := strings.DelEx("TestDelExIfdEqNoMatch", "IFDEQ", "value1")
	if err == nil || err.Error() != "ERR condition not met" {
		t.Errorf("Expected error: %v, got: %v", "ERR condition not met", err)
	}

	val, err := strings.Get("TestDelExIfdEqNoMatch")
	if err != nil || val != "value1" {
		t.Errorf("Expected key to remain untouched, got val: %v, err: %v", val, err)
	}

	err = strings.DelEx("TestDelExIfdEqNoMatch", "IFDEQ", strings.Digest("value2"))
	if err == nil || err.Error() != "ERR condition not met" {
		t.Errorf("Expected error: %v, got: %v", "ERR condition not met", err)
	}
}

func TestDelExIfdNeMatch(t *testing.T) {
	strings.Set("TestDelExIfdNeMatch", "value1")

	err := strings.DelEx("TestDelExIfdNeMatch", "IFDNE", strings.Digest("value2"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExIfdNeMatch")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}

func TestDelExIfdNeNoMatch(t *testing.T) {
	strings.Set("TestDelExIfdNeNoMatch", "value1")

	err := strings.DelEx("TestDelExIfdNeNoMatch", "IFDNE", strings.Digest("value1"))
	if err == nil || err.Error() != "ERR condition not met" {
		t.Errorf("Expected error: %v, got: %v", "ERR condition not met", err)
	}

	val, err := strings.Get("TestDelExIfdNeNoMatch")
	if err != nil || val != "value1" {
		t.Errorf("Expected key to remain untouched, got val: %v, err: %v", val, err)
	}
}

func TestDelExCaseInsensitiveExpression(t *testing.T) {
	strings.Set("TestDelExCaseInsensitiveExpression", "value1")

	err := strings.DelEx("TestDelExCaseInsensitiveExpression", "ifeq", "value1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = strings.Get("TestDelExCaseInsensitiveExpression")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got none")
	}
}
