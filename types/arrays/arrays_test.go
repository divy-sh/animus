package arrays_test

import (
	"reflect"
	"testing"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
	"github.com/divy-sh/animus/types/arrays"
)

func TestArGet(t *testing.T) {
	t.Run("returns the value at a valid index", func(t *testing.T) {
		key := "testArrayGetType"
		store.Set(key, []any{1, 2, 3, 4, 5})

		value, err := arrays.ArGet(key, 2)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if value != 3 {
			t.Fatalf("Expected value to be 3, got %v", value)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArGet("missingArray", 0)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an out-of-bounds index", func(t *testing.T) {
		key := "outOfBoundsArrayType"
		store.Set(key, []any{1, 2, 3})

		_, err := arrays.ArGet(key, 5)
		if err == nil {
			t.Fatalf("Expected an error for out-of-bounds index, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArCount(t *testing.T) {
	t.Run("returns the length of an existing array", func(t *testing.T) {
		key := "testArrayCountType"
		store.Set(key, []any{1, 2, 3, 4})

		count, err := arrays.ArCount(key)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if count != 4 {
			t.Fatalf("Expected count to be 4, got %d", count)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArCount("missingCountArray")
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})
}

func TestArDel(t *testing.T) {
	t.Run("removes the value at a valid index", func(t *testing.T) {
		key := "testArrayDelType"
		store.Set(key, []any{1, 2, 3, 4})

		err := arrays.ArDel(key, 1)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		updated, ok := store.Get[string, []any](key)
		if !ok {
			t.Fatalf("Expected array to remain stored")
		}
		if !reflect.DeepEqual(updated, []any{1, 3, 4}) {
			t.Fatalf("Expected array to become [1 3 4], got %v", updated)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		err := arrays.ArDel("missingDeleteArray", 0)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an out-of-bounds index", func(t *testing.T) {
		key := "testArrayDelOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		err := arrays.ArDel(key, 5)
		if err == nil {
			t.Fatalf("Expected an error for out-of-bounds index, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArDelRange(t *testing.T) {
	t.Run("removes the values in a valid range", func(t *testing.T) {
		key := "testArrayDelRangeType"
		store.Set(key, []any{10, 20, 30, 40, 50})

		err := arrays.ArDelRange(key, 1, 3)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		updated, ok := store.Get[string, []any](key)
		if !ok {
			t.Fatalf("Expected array to remain stored")
		}
		if !reflect.DeepEqual(updated, []any{10, 50}) {
			t.Fatalf("Expected array to become [10 50], got %v", updated)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		err := arrays.ArDelRange("missingDeleteRangeArray", 0, 1)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an invalid range", func(t *testing.T) {
		key := "testArrayDelRangeOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		err := arrays.ArDelRange(key, 1, 5)
		if err == nil {
			t.Fatalf("Expected an error for invalid range, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArGetRange(t *testing.T) {
	t.Run("returns the values in a valid range", func(t *testing.T) {
		key := "testArrayGetRangeType"
		store.Set(key, []any{1, 2, 3, 4, 5})

		values, err := arrays.ArGetRange(key, 1, 3)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{2, 3, 4}) {
			t.Fatalf("Expected values [2 3 4], got %v", values)
		}
	})

	t.Run("returns an error for a missing array", func(t *testing.T) {
		_, err := arrays.ArGetRange("missingArrayRange", 0, 1)
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("returns an error for an invalid range", func(t *testing.T) {
		key := "testArrayGetRangeOutOfBounds"
		store.Set(key, []any{1, 2, 3})

		_, err := arrays.ArGetRange(key, 1, 5)
		if err == nil {
			t.Fatalf("Expected an error for invalid range, got nil")
		}
		if err.Error() != common.ERR_INDEX_OUT_OF_BOUNDS {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_INDEX_OUT_OF_BOUNDS, err.Error())
		}
	})
}

func TestArGrep(t *testing.T) {
	t.Run("matches array values with asterisk wildcard", func(t *testing.T) {
		key := "testArrayGrepAsterisk"
		store.Set(key, []any{"hello", "world", "hallo", "hi", "help"})

		values, err := arrays.ArGrep(key, "h*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"hello", "hallo", "hi", "help"}) {
			t.Fatalf("Expected values [hello hallo hi help], got %v", values)
		}
	})

	t.Run("matches array values with asterisk wildcard in middle", func(t *testing.T) {
		key := "testArrayGrepAsteriskMiddle"
		store.Set(key, []any{"cat", "car", "card", "dog", "cast"})

		values, err := arrays.ArGrep(key, "ca*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"cat", "car", "card", "cast"}) {
			t.Fatalf("Expected values [cat car card cast], got %v", values)
		}
	})

	t.Run("matches array values with question mark wildcard", func(t *testing.T) {
		key := "testArrayGrepQuestion"
		store.Set(key, []any{"cat", "cot", "cut", "dog", "ca"})

		values, err := arrays.ArGrep(key, "c?t")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"cat", "cot", "cut"}) {
			t.Fatalf("Expected values [cat cot cut], got %v", values)
		}
	})

	t.Run("matches array values with multiple question marks", func(t *testing.T) {
		key := "testArrayGrepMultipleQuestion"
		store.Set(key, []any{"ab", "abc", "abcd", "a", "acd"})

		values, err := arrays.ArGrep(key, "a??")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"abc", "acd"}) {
			t.Fatalf("Expected values [abc acd], got %v", values)
		}
	})

	t.Run("matches array values with character class", func(t *testing.T) {
		key := "testArrayGrepCharClass"
		store.Set(key, []any{"cat", "cot", "cut", "dog", "cbt"})

		values, err := arrays.ArGrep(key, "c[ao]t")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"cat", "cot"}) {
			t.Fatalf("Expected values [cat cot], got %v", values)
		}
	})

	t.Run("matches array values with character range in class", func(t *testing.T) {
		key := "testArrayGrepCharRange"
		store.Set(key, []any{"cat", "cot", "cut", "cet", "dog"})

		values, err := arrays.ArGrep(key, "c[a-o]t")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(values) != 3 || !contains(values, "cat") || !contains(values, "cet") || !contains(values, "cot") {
			t.Fatalf("Expected values to contain [cat cet cot], got %v", values)
		}
	})

	t.Run("matches array values with negated character class", func(t *testing.T) {
		key := "testArrayGrepNegatedClass"
		store.Set(key, []any{"cat", "cot", "cut", "c1t", "c#t"})

		values, err := arrays.ArGrep(key, "c[^ao]t")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"cut", "c1t", "c#t"}) {
			t.Fatalf("Expected values [cut c1t c#t], got %v", values)
		}
	})

	t.Run("matches array values with escaped asterisk", func(t *testing.T) {
		key := "testArrayGrepEscapedAsterisk"
		store.Set(key, []any{"a*b", "aab", "abb", "a*c"})

		values, err := arrays.ArGrep(key, "a\\*b")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"a*b"}) {
			t.Fatalf("Expected values [a*b], got %v", values)
		}
	})

	t.Run("matches array values with escaped question mark", func(t *testing.T) {
		key := "testArrayGrepEscapedQuestion"
		store.Set(key, []any{"a?b", "aab", "abb", "a?c"})

		values, err := arrays.ArGrep(key, "a\\?b")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"a?b"}) {
			t.Fatalf("Expected values [a?b], got %v", values)
		}
	})

	t.Run("matches with complex pattern combining wildcards and classes", func(t *testing.T) {
		key := "testArrayGrepComplex"
		store.Set(key, []any{"test1.go", "test2.go", "hello.go", "test_a.go", "testing.py"})

		values, err := arrays.ArGrep(key, "test?.go")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"test1.go", "test2.go"}) {
			t.Fatalf("Expected values [test1.go test2.go], got %v", values)
		}
	})

	t.Run("returns empty result when no matches found", func(t *testing.T) {
		key := "testArrayGrepNoMatch"
		store.Set(key, []any{"hello", "world", "foo"})

		values, err := arrays.ArGrep(key, "z*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if values != nil && len(values) != 0 {
			t.Fatalf("Expected empty result, got %v", values)
		}
	})

	t.Run("returns empty result for empty array", func(t *testing.T) {
		key := "testArrayGrepEmpty"
		store.Set(key, []any{})

		values, err := arrays.ArGrep(key, "h*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if values != nil && len(values) != 0 {
			t.Fatalf("Expected empty result, got %v", values)
		}
	})

	t.Run("skips non-string elements in array", func(t *testing.T) {
		key := "testArrayGrepMixedTypes"
		store.Set(key, []any{"hello", 42, "hallo", 3.14, "hi", true})

		values, err := arrays.ArGrep(key, "h*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"hello", "hallo", "hi"}) {
			t.Fatalf("Expected values [hello hallo hi], got %v", values)
		}
	})

	t.Run("returns error for non-existent array", func(t *testing.T) {
		_, err := arrays.ArGrep("nonExistentArray", "h*")
		if err == nil {
			t.Fatalf("Expected an error for non-existent array, got nil")
		}
		if err.Error() != common.ERR_ARRAY_NOT_FOUND {
			t.Fatalf("Expected error message '%s', got '%s'", common.ERR_ARRAY_NOT_FOUND, err.Error())
		}
	})

	t.Run("matches exact strings", func(t *testing.T) {
		key := "testArrayGrepExact"
		store.Set(key, []any{"exact", "exacto", "exact match", "notexact"})

		values, err := arrays.ArGrep(key, "exact")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"exact"}) {
			t.Fatalf("Expected values [exact], got %v", values)
		}
	})

	t.Run("matches with asterisk at end", func(t *testing.T) {
		key := "testArrayGrepAsteriskEnd"
		store.Set(key, []any{"testing", "test", "tested", "tasting", "tea"})

		values, err := arrays.ArGrep(key, "test*")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"testing", "test", "tested"}) {
			t.Fatalf("Expected values [testing test tested], got %v", values)
		}
	})

	t.Run("matches with asterisk at start", func(t *testing.T) {
		key := "testArrayGrepAsteriskStart"
		store.Set(key, []any{"testing", "parsing", "working", "coding"})

		values, err := arrays.ArGrep(key, "*ing")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"testing", "parsing", "working", "coding"}) {
			t.Fatalf("Expected values [testing parsing working coding], got %v", values)
		}
	})

	t.Run("matches with multiple asterisks", func(t *testing.T) {
		key := "testArrayGrepMultipleAsterisk"
		store.Set(key, []any{"a1b2c", "a1b", "ac", "a1c", "axbxc"})

		values, err := arrays.ArGrep(key, "a*b*c")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"a1b2c", "axbxc"}) {
			t.Fatalf("Expected values [a1b2c axbxc], got %v", values)
		}
	})

	t.Run("matches with negated character class using exclamation", func(t *testing.T) {
		key := "testArrayGrepNegatedExclamation"
		store.Set(key, []any{"cat", "cot", "cut", "c1t", "c#t"})

		values, err := arrays.ArGrep(key, "c[!ao]t")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !reflect.DeepEqual(values, []any{"cut", "c1t", "c#t"}) {
			t.Fatalf("Expected values [cut c1t c#t], got %v", values)
		}
	})
}

func contains(slice []any, value string) bool {
	for _, v := range slice {
		if str, ok := v.(string); ok && str == value {
			return true
		}
	}
	return false
}
