package generics

import (
	"errors"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func Copy(source, destination string) (int, error) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	value, ok := store.Get[any, any](source)
	if !ok {
		return 0, errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}
	store.Set(destination, value)
	return 1, nil
}

func Delete(keys *[]string) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	for _, key := range *keys {
		store.Delete(key)
	}
}

func Exists(keys *[]string) int {
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
	validKeyCount := 0
	for _, key := range *keys {
		_, exists := store.Get[any, any](key)
		if exists {
			validKeyCount++
		}
	}
	return validKeyCount
}
