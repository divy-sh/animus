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
