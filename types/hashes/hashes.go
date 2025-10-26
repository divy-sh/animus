package hashes

import (
	"errors"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
	"github.com/divy-sh/animus/types/generics"
)

func HGet(hash, key string) (string, error) {
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
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
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	hashVal, ok := store.Get[string, map[string]string](hash)
	if ok {
		hashVal[key] = value
	} else {
		hashVal = map[string]string{key: value}
	}
	store.Set(hash, hashVal)
}

func HExists(hash, key string) (int64, error) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	_, ok := store.Get[string, map[string]string](hash)
	if ok {
		return 1, nil
	} else {
		return 0, errors.New(common.ERR_HASH_NOT_FOUND)
	}
}

func HExpire(key, seconds, flag string) error {
	return generics.Expire(key, seconds, flag)
}

func HDel(hash, key string) error {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	hashVal, ok := store.Get[string, map[string]string](hash)
	if !ok {
		return errors.New(common.ERR_HASH_NOT_FOUND)
	}

	if _, ok := hashVal[key]; !ok {
		return errors.New(common.ERR_KEY_NOT_FOUND)
	}
	delete(hashVal, key)
	if len(hashVal) == 0 {
		store.Delete(hash)
		return nil
	} else {
		store.Set(hash, hashVal)
	}
	return nil
}

func HGetAll(key string) (map[string]string, error) {
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
	hashVal, ok := store.Get[string, map[string]string](key)
	if !ok {
		return nil, errors.New(common.ERR_HASH_NOT_FOUND)
	}
	return hashVal, nil
}
