package essentias

import (
	"errors"
	"time"

	"github.com/divy-sh/animus/store"
)

func HGet(hash, key string) (string, error) {
	lock := store.GetLock(key)
	lock.RLock()
	defer lock.RUnlock()
	value, ok := store.Get[string, map[string]string](hash)
	if !ok {
		return "", errors.New("ERR hash does not exist")
	}
	if val, ok := value[key]; ok {
		return val, nil
	}
	return "", errors.New("ERR hash does not exist")
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
	store.Set(hash, hashVal, time.Now().AddDate(1000, 0, 0))
}
