package essentias

import (
	"errors"
	"sync"
	"time"

	"github.com/divy-sh/animus/store"
)

type HashEssentia struct {
	hashes store.Store[string, map[string]string]
	locks  sync.Map
}

func NewHashEssentia() *HashEssentia {
	return &HashEssentia{
		hashes: *store.NewStore[string, map[string]string](),
	}
}

func (h *HashEssentia) getLock(key string) *sync.RWMutex {
	actual, _ := h.locks.LoadOrStore(key, &sync.RWMutex{})
	return actual.(*sync.RWMutex)
}

func (h *HashEssentia) HGet(hash, key string) (string, error) {
	lock := h.getLock(key)
	lock.RLock()
	defer lock.RUnlock()
	value, ok := h.hashes.Get(hash)
	if !ok {
		return "", errors.New("ERR not found")
	}
	if val, ok := value[key]; ok {
		return val, nil
	}
	return "", errors.New("ERR not found")
}

func (h *HashEssentia) HSet(hash, key, value string) {
	lock := h.getLock(key)
	lock.Lock()
	defer lock.Unlock()
	hashVal, ok := h.hashes.Get(hash)
	if ok {
		hashVal[key] = value
	} else {
		hashVal = map[string]string{key: value}
	}
	h.hashes.Set(hash, hashVal, time.Now().AddDate(1000, 0, 0))
}
