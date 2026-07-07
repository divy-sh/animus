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