package generics

import (
	"errors"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func Copy(source, destination string) (int, error) {
	lock := store.GetLock(source)
	lock.RLock()
	value, ok := store.Get[any, any](source)
	lock.RUnlock()

	if !ok {
		return 0, errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}

	lock2 := store.GetLock(destination)
	lock2.Lock()
	defer lock2.Unlock()
	store.Set(destination, value)
	return 1, nil
}

func Delete(keys *[]string) {
	for _, key := range *keys {
		lock := store.GetLock(key)
		lock.Lock()
		store.DeleteWithKey(key)
		lock.Unlock()
	}
}

func Exists(keys *[]string) int {
	validKeyCount := 0
	for _, key := range *keys {
		lock := store.GetLock(key)
		lock.RLock()
		_, exists := store.Get[any, any](key)
		if exists {
			validKeyCount++
		}
		lock.RUnlock()
	}
	return validKeyCount
}
