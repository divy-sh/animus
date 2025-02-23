package types

import (
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

type Expiry struct {
	ttl     *btree.BTree
	lru     *btree.BTree
	maxKeys int
}

func NewExpiry() Expiry {
	return Expiry{
		ttl:     btree.New(2),
		lru:     btree.New(2),
		maxKeys: 1000,
	}
}
func (exp *Expiry) updateLRU(key string) {
	exp.lru.ReplaceOrInsert(LRUItem{Key: key, Timestamp: time.Now()})
	exp.lazyEvictLRU()
}

func (exp *Expiry) removeLRU(key string) {
	var itemToDelete LRUItem
	exp.lru.Ascend(func(item btree.Item) bool {
		if item.(LRUItem).Key == key {
			itemToDelete = item.(LRUItem)
			return false
		}
		return true
	})
	if itemToDelete.Key != "" {
		exp.lru.Delete(itemToDelete)
	}
}

func (exp *Expiry) lazyEvictLRU() (string, bool) {
	if exp.lru.Len() > exp.maxKeys {
		min := exp.lru.Min()
		if min == nil {
			return "", false
		}
		item := min.(LRUItem)
		exp.removeTTL(item.Key)
		exp.lru.Delete(item)
		return item.Key, true
	}
	return "", false
}

func (exp *Expiry) setTTL(key string, expires time.Time) {
	exp.ttl.ReplaceOrInsert(TTLItem{Key: key, Expires: expires})
}

func (exp *Expiry) removeTTL(key string) {
	var itemToDelete TTLItem
	exp.ttl.Ascend(func(item btree.Item) bool {
		if item.(TTLItem).Key == key {
			itemToDelete = item.(TTLItem)
			return false
		}
		return true
	})
	if itemToDelete.Key != "" {
		exp.ttl.Delete(itemToDelete)
	}
}

func (exp *Expiry) lazyEvict() (string, bool) {
	now := time.Now()
	item := exp.ttl.Min()
	if item == nil {
		return "", false
	}
	ttlItem := item.(TTLItem)
	if ttlItem.Expires.After(now) {
		return "", false
	}
	exp.ttl.Delete(ttlItem)
	exp.removeLRU(ttlItem.Key)
	return ttlItem.Key, true
}

// utility functions
func (item TTLItem) Less(other btree.Item) bool {
	return item.Expires.Before(other.(TTLItem).Expires)
}

func (item LRUItem) Less(other btree.Item) bool {
	return item.Timestamp.Before(other.(LRUItem).Timestamp)
}
