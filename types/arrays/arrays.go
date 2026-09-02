package arrays

import (
	"errors"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func ArCount(key string) (int64, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return 0, errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	return int64(len(arr)), nil
}

func ArDel(key string, index int64) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	if index < 0 || index >= int64(len(arr)) {
		return errors.New(common.ERR_INDEX_OUT_OF_BOUNDS)
	}

	arr = append(arr[:index], arr[index+1:]...)
	store.Set(key, arr)

	return nil
}

func ArDelRange(key string, start, end int64) error {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	if start < 0 || end >= int64(len(arr)) || start > end {
		return errors.New(common.ERR_INDEX_OUT_OF_BOUNDS)
	}

	arr = append(arr[:start], arr[end+1:]...)
	store.Set(key, arr)

	return nil
}

func ArGet(key string, index int64) (any, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return nil, errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	if index < 0 || index >= int64(len(arr)) {
		return nil, errors.New(common.ERR_INDEX_OUT_OF_BOUNDS)
	}

	return arr[index], nil
}

func ArGetRange(key string, start, end int64) ([]any, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return nil, errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	if start < 0 || end >= int64(len(arr)) || start > end {
		return nil, errors.New(common.ERR_INDEX_OUT_OF_BOUNDS)
	}

	return arr[start : end+1], nil
}

func ArGrep(key string, pattern string) ([]any, error) {
	store.LockKeys(key)
	defer store.UnlockKeys(key)

	arr, ok := store.Get[string, []any](key)
	if !ok {
		return nil, errors.New(common.ERR_ARRAY_NOT_FOUND)
	}

	var result []any
	for _, v := range arr {
		str, ok := v.(string)
		if !ok {
			continue
		}
		if matchPattern(str, pattern) {
			result = append(result, str)
		}
	}

	return result, nil
}

// Private helper functions

func matchPattern(str, pattern string) bool {
	return matchPatternRecursive([]rune(str), []rune(pattern), 0, 0)
}

func matchPatternRecursive(str, pattern []rune, strIndex, patternIndex int) bool {
	if patternIndex == len(pattern) {
		return strIndex == len(str)
	}

	char := pattern[patternIndex]
	switch char {
	case '*':
		if patternIndex+1 == len(pattern) {
			return true
		}
		return matchPatternRecursive(str, pattern, strIndex, patternIndex+1) ||
			(strIndex < len(str) && matchPatternRecursive(str, pattern, strIndex+1, patternIndex))
	case '?':
		return strIndex < len(str) && matchPatternRecursive(str, pattern, strIndex+1, patternIndex+1)
	case '\\':
		if patternIndex+1 >= len(pattern) {
			return strIndex < len(str) && str[strIndex] == '\\'
		}
		if strIndex < len(str) && str[strIndex] == pattern[patternIndex+1] {
			return matchPatternRecursive(str, pattern, strIndex+1, patternIndex+2)
		}
		return false
	case '[':
		matched, nextIndex, ok := matchClass(str, pattern, strIndex, patternIndex)
		if !ok {
			return false
		}
		return matched && matchPatternRecursive(str, pattern, strIndex+1, nextIndex)
	default:
		return strIndex < len(str) && str[strIndex] == char && matchPatternRecursive(str, pattern, strIndex+1, patternIndex+1)
	}
}

func matchClass(str, pattern []rune, strIndex, patternIndex int) (bool, int, bool) {
	if strIndex >= len(str) {
		return false, patternIndex, false
	}

	end := patternIndex + 1
	for end < len(pattern) && pattern[end] != ']' {
		end++
	}
	if end >= len(pattern) {
		return false, patternIndex, false
	}

	negated := false
	start := patternIndex + 1
	if start < end && (pattern[start] == '^' || pattern[start] == '!') {
		negated = true
		start++
	}

	matched := false
	for i := start; i < end; {
		if pattern[i] == '\\' && i+1 < end {
			if str[strIndex] == pattern[i+1] {
				matched = true
			}
			i += 2
			continue
		}

		if i+2 < end && pattern[i+1] == '-' {
			if pattern[i] <= str[strIndex] && str[strIndex] <= pattern[i+2] {
				matched = true
			}
			i += 3
			continue
		}

		if str[strIndex] == pattern[i] {
			matched = true
		}
		i++
	}

	if negated {
		matched = !matched
	}
	return matched, end + 1, true
}
