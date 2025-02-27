package types

import (
	"errors"
	"time"

	"github.com/divy-sh/animus/store"
)

type HashType struct {
	hashes store.Store[string, map[string]string]
}

func NewHashType() *HashType {
	return &HashType{
		hashes: *store.NewStore[string, map[string]string](),
	}
}

func (h *HashType) HGet(hash, key string) (string, error) {
	value, ok := h.hashes.Get(hash)
	if !ok {
		return "", errors.New("ERR not found")
	}
	if val, ok := value[key]; ok {
		return val, nil
	}
	return "", errors.New("ERR not found")
}

func (h *HashType) HSet(hash, key, value string) {
	hashVal, ok := h.hashes.Get(key)
	if ok {
		hashVal[key] = value
	} else {
		hashVal = map[string]string{key: value}
	}
	h.hashes.Set(hash, hashVal, time.Now().AddDate(1000, 0, 0))
}
