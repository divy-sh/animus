package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/google/btree"
)

type TTLItem struct {
	Key     string
	Expires time.Time
}

type LRUItem struct {
	Key       string
	Timestamp time.Time
}

type StringType struct {
	strings map[string]string
	ttl     *btree.BTree
	lru     *btree.BTree
	muLock  sync.RWMutex
	maxKeys int
}

func NewStringType() *StringType {
	return &StringType{
		strings: make(map[string]string),
		ttl:     btree.New(2),
		lru:     btree.New(2),
		muLock:  sync.RWMutex{},
		maxKeys: 1000,
	}

}

// public functions
func (s *StringType) Append(key, value string) {
	s.muLock.Lock()
	s.strings[key] += value
	s.updateLRU(key)
	s.muLock.Unlock()
}

func (s *StringType) Decr(key string) error {
	s.muLock.Lock()
	if val, ok := s.strings[key]; ok {
		val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return errors.New("ERR cannot decrement a non integer value")
		}
		s.strings[key] = fmt.Sprint(val - 1)
		s.updateLRU(key)
	} else {
		s.strings[key] = "0"
		s.updateLRU(key)
	}
	s.muLock.Unlock()
	return nil
}

func (s *StringType) DecrBy(key, value string) error {
	decrVal, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return errors.New("ERR invalid decrement value")
	}

	s.muLock.Lock()
	if val, ok := s.strings[key]; ok {
		val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return errors.New("ERR cannot decrement a non integer value")
		}
		s.strings[key] = fmt.Sprint(val - decrVal)
		s.updateLRU(key)
	} else {
		s.strings[key] = "0"
		s.updateLRU(key)
	}
	s.muLock.Unlock()
	return nil
}

func (s *StringType) Get(key string) (string, error) {
	s.muLock.RLock()
	value, ok := s.strings[key]
	s.muLock.RUnlock()
	if !ok {
		return "", errors.New("ERR not found")
	}
	s.muLock.Lock()
	s.updateLRU(key)
	s.lazyEvict()
	s.muLock.Unlock()
	return value, nil
}

func (s *StringType) GetDel(key string) (string, error) {
	s.muLock.Lock()
	value, ok := s.strings[key]
	delete(s.strings, key)
	s.removeTTL(key)
	s.removeLRU(key)
	s.lazyEvict()
	s.muLock.Unlock()
	if !ok {
		return "", errors.New("ERR not found")
	}
	return value, nil
}

func (s *StringType) Set(key, value string) {
	s.muLock.Lock()
	s.strings[key] = value
	s.updateLRU(key)
	s.lazyEvict()
	s.muLock.Unlock()
}

// private functions
func (s *StringType) updateLRU(key string) {
	s.lru.ReplaceOrInsert(LRUItem{Key: key, Timestamp: time.Now()})
	s.lazyEvictLRU()
}

func (s *StringType) removeLRU(key string) {
	var itemToDelete LRUItem
	s.lru.Ascend(func(item btree.Item) bool {
		if item.(LRUItem).Key == key {
			itemToDelete = item.(LRUItem)
			return false
		}
		return true
	})
	if itemToDelete.Key != "" {
		s.lru.Delete(itemToDelete)
	}
}

func (s *StringType) lazyEvictLRU() {
	for s.lru.Len() > s.maxKeys {
		min := s.lru.Min()
		if min == nil {
			return
		}
		item := min.(LRUItem)
		delete(s.strings, item.Key)
		s.removeTTL(item.Key)
		s.lru.Delete(item)
	}
}

func (s *StringType) setTTL(key string, expires time.Time) {
	s.ttl.ReplaceOrInsert(TTLItem{Key: key, Expires: expires})
}

func (s *StringType) removeTTL(key string) {
	var itemToDelete TTLItem
	s.ttl.Ascend(func(item btree.Item) bool {
		if item.(TTLItem).Key == key {
			itemToDelete = item.(TTLItem)
			return false
		}
		return true
	})
	if itemToDelete.Key != "" {
		s.ttl.Delete(itemToDelete)
	}
}

func (s *StringType) lazyEvict() {
	now := time.Now()
	for {
		item := s.ttl.Min()
		if item == nil {
			return
		}
		ttlItem := item.(TTLItem)
		if ttlItem.Expires.After(now) {
			return
		}
		delete(s.strings, ttlItem.Key)
		s.ttl.Delete(ttlItem)
		s.removeLRU(ttlItem.Key)
	}
}

// utility functions
func (item TTLItem) Less(other btree.Item) bool {
	return item.Expires.Before(other.(TTLItem).Expires)
}

func (item LRUItem) Less(other btree.Item) bool {
	return item.Timestamp.Before(other.(LRUItem).Timestamp)
}
