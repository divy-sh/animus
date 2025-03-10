package essentias

import (
	"errors"
	"sync"
	"time"

	"github.com/divy-sh/animus/store"
)

type HashEssentia struct {
	hashes store.Store[string, map[string]string]
	lock   sync.RWMutex
}

func NewHashEssentia() *HashEssentia {
	return &HashEssentia{
		hashes: *store.NewStore[string, map[string]string](),
	}
}

func (h *HashEssentia) HGet(hash, key string) (string, error) {
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
	h.lock.Lock()
	defer h.lock.Unlock()
	hashVal, ok := h.hashes.Get(hash)
	if ok {
		hashVal[key] = value
	} else {
		hashVal = map[string]string{key: value}
	}
	h.hashes.Set(hash, hashVal, time.Now().AddDate(1000, 0, 0))
}
