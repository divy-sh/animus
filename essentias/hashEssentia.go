package essentias

import (
	"errors"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

func HGet(hash, key string) (string, error) {
	lock := store.GetLock(key)
	lock.RLock()
	defer lock.RUnlock()
	value, ok := store.Get[string, map[string]string](hash)
	if !ok {
		return "", errors.New(common.ERR_HASH_NOT_FOUND)
	}
	if val, ok := value[key]; ok {
		return val, nil
	}
	return "", errors.New(common.ERR_HASH_NOT_FOUND)
}

func HSet(hash, key, value string) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	hashVal, ok := store.Get[string, map[string]string](hash)
	if ok {
		hashVal[key] = value
	} else {
		hashVal = map[string]string{key: value}
	}
	store.Set(hash, hashVal)
}
