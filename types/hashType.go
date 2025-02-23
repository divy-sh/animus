package types

import (
	"errors"
	"sync"
)

type HashType struct {
	hashes  map[string]map[string]string
	muLock  sync.RWMutex
	expiry  Expiry
	maxKeys int
}

func NewHashType() *HashType {
	return &HashType{
		hashes:  make(map[string]map[string]string),
		muLock:  sync.RWMutex{},
		expiry:  NewExpiry(),
		maxKeys: 1000,
	}
}

func (h *HashType) HGet(hash, key string) (string, error) {
	h.muLock.RLock()
	value, ok := h.hashes[hash][key]
	if !ok {
		return "", errors.New("ERR not found")
	}
	h.muLock.RUnlock()
	h.muLock.Lock()
	h.expiry.updateLRU(hash)
	key, ok = h.expiry.lazyEvict()
	if ok {
		delete(h.hashes, key)
	}
	h.muLock.Unlock()
	return value, nil
}

func (h *HashType) HSet(hash, key, value string) {
	h.muLock.Lock()
	h.hashes[hash][key] = value
	h.expiry.updateLRU(key)
	hash, ok := h.expiry.lazyEvict()
	if ok {
		delete(h.hashes, hash)
	}
	h.muLock.Unlock()
}
