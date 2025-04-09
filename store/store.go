package store

import (
	"container/list"
	"math"
	"sync"
	"time"

	"github.com/google/btree"
)

/* Essentias */

type Store struct {
	dict    map[any]*list.Element
	queue   *list.List
	expTree *btree.BTree
	maxSize int
	locks   sync.Map
}

type DataNode struct {
	key  any
	data any
	ttl  int64
}

type BTreeItem struct {
	ttl int64
	key any
}

func (a BTreeItem) Less(b btree.Item) bool {
	return a.ttl >= 0 && a.ttl < b.(BTreeItem).ttl
}

/* Singleton instance */
var (
	sharedInstance *Store
	once           sync.Once
	storeConfig    = StoreConfig{MaxSize: 100_000} // Default config
)

/* Singleton Getter */

// GetSharedStore ensures a single global store instance
func GetSharedStore() *Store {
	if sharedInstance == nil {
		NewStoreBuilder().Build() // Fallback to default config if not explicitly built
	}
	return sharedInstance
}

/* Public functions */

func GetLock(key any) *sync.RWMutex {
	store := GetSharedStore()
	actual, _ := store.locks.LoadOrStore(key, &sync.RWMutex{})
	return actual.(*sync.RWMutex)
}

func GetLocks(keys *[]any) []*sync.RWMutex {
	store := GetSharedStore()
	locks := []*sync.RWMutex{}
	for _, key := range *keys {
		actual, _ := store.locks.LoadOrStore(key, &sync.RWMutex{})
		locks = append(locks, actual.(*sync.RWMutex))
	}
	return locks
}

// Get retrieves a value with type inference
func Get[K comparable, V any](key K) (V, bool) {
	store := GetSharedStore()

	node, found := store.dict[key]
	if !found {
		var zero V
		return zero, false
	} else if node.Value.(*DataNode).ttl <= time.Now().Unix() {
		store.evictWithKey(key)
		var zero V
		return zero, false
	}

	store.queue.MoveToFront(node)
	value, ok := node.Value.(*DataNode).data.(V)
	if !ok {
		var zero V
		return zero, false
	}
	return value, true
}

// Set stores a value with TTL
func Set[K comparable, V any](key K, value V) {
	store := GetSharedStore()
	ttl := int64(math.MaxInt64)
	node, found := store.dict[key]
	if found {
		node.Value.(*DataNode).data = value
		node.Value.(*DataNode).ttl = ttl
		store.queue.MoveToFront(node)
		store.expTree.Delete(BTreeItem{ttl: node.Value.(*DataNode).ttl, key: key})
		store.expTree.ReplaceOrInsert(BTreeItem{ttl: ttl, key: key})
		return
	}

	storeItem := &DataNode{
		key:  key,
		data: value,
		ttl:  ttl,
	}
	queueItem := store.queue.PushFront(storeItem)
	store.dict[key] = queueItem
	store.expTree.ReplaceOrInsert(BTreeItem{ttl: ttl, key: key})
	store.lazyEvict()
}

func SetWithTTL[K comparable, V any](key K, value V, ttl int64) {
	store := GetSharedStore()
	ttl = time.Now().Unix() + ttl
	node, found := store.dict[key]
	if found {
		node.Value.(*DataNode).data = value
		node.Value.(*DataNode).ttl = ttl
		store.queue.MoveToFront(node)
		store.expTree.Delete(BTreeItem{ttl: node.Value.(*DataNode).ttl, key: key})
		store.expTree.ReplaceOrInsert(BTreeItem{ttl: ttl, key: key})
		return
	}

	storeItem := &DataNode{
		key:  key,
		data: value,
		ttl:  ttl,
	}
	queueItem := store.queue.PushFront(storeItem)
	store.dict[key] = queueItem
	store.expTree.ReplaceOrInsert(BTreeItem{ttl: ttl, key: key})
	store.lazyEvict()
}

// DeleteWithKey removes a key
func DeleteWithKey[K comparable](key K) bool {
	store := GetSharedStore()

	store.evictWithKey(key)
	return true
}

/* Private functions */

func (s *Store) lazyEvict() {
	if len(s.dict) <= s.maxSize {
		return
	}
	if !s.evictTtlExpired() {
		s.evictLru()
	}
}

func (s *Store) evictTtlExpired() bool {
	item := s.expTree.Min()
	if item != nil && item.(BTreeItem).ttl < time.Now().Unix() {
		key := s.expTree.Delete(item).(BTreeItem).key
		s.queue.Remove(s.dict[key])
		delete(s.dict, key)
		return true
	}
	return false
}

func (s *Store) evictLru() bool {
	lruItem := s.queue.Back()
	if lruItem != nil {
		dataNode := lruItem.Value.(*DataNode)
		s.expTree.Delete(BTreeItem{ttl: dataNode.ttl, key: dataNode.key})
		s.queue.Remove(lruItem)
		delete(s.dict, dataNode.key)
		return true
	}
	return false
}

func (s *Store) evictWithKey(key any) bool {
	s.expTree.Delete(BTreeItem{ttl: s.dict[key].Value.(*DataNode).ttl, key: key})
	s.queue.Remove(s.dict[key])
	delete(s.dict, key)
	return true
}
