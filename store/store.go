package store

import (
	"container/list"
	"sync"
	"time"

	"github.com/google/btree"
)

/* Essentias */

type Store[K comparable, V any] struct {
	dict    map[K]*list.Element
	queue   *list.List
	expTree *btree.BTree
	maxSize int
	locks   sync.Map
}

type DataNode[K comparable, V any] struct {
	key  K
	data V
	ttl  time.Time
}

type BTreeItem[K comparable] struct {
	ttl time.Time
	key K
}

func (a BTreeItem[K]) Less(b btree.Item) bool {
	return a.ttl.Before(b.(BTreeItem[K]).ttl)
}

/* Public functions */

func NewStore[K comparable, V any]() *Store[K, V] {
	return &Store[K, V]{
		dict:    make(map[K]*list.Element),
		queue:   list.New(),
		expTree: btree.New(2),
		maxSize: 100_000,
	}
}

func (s *Store[K, V]) getLock(key K) *sync.RWMutex {
	actual, _ := s.locks.LoadOrStore(key, &sync.RWMutex{})
	return actual.(*sync.RWMutex)
}

func (s *Store[K, V]) Get(key K) (V, bool) {
	lock := s.getLock(key)
	lock.RLock()
	defer lock.Unlock()

	node, found := s.dict[key]
	if !found {
		var zero V
		return zero, false
	} else if node.Value.(*DataNode[K, V]).ttl.Before(time.Now()) {
		s.evictWithKey(key)
		var zero V
		return zero, false
	}

	s.queue.MoveToFront(node)
	return node.Value.(*DataNode[K, V]).data, true
}

func (s *Store[K, V]) Set(key K, value V, ttl time.Time) {
	lock := s.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	node, found := s.dict[key]
	if found {
		node.Value.(*DataNode[K, V]).data = value
		node.Value.(*DataNode[K, V]).ttl = ttl
		s.queue.MoveToFront(node)
		s.expTree.Delete(BTreeItem[K]{ttl: node.Value.(*DataNode[K, V]).ttl, key: key})
		s.expTree.ReplaceOrInsert(BTreeItem[K]{ttl: ttl, key: key})
		return
	}

	storeItem := &DataNode[K, V]{
		key:  key,
		data: value,
		ttl:  ttl,
	}
	queueItem := s.queue.PushFront(storeItem)
	s.dict[key] = queueItem
	s.expTree.ReplaceOrInsert(BTreeItem[K]{ttl: ttl, key: key})
	s.lazyEvict()
}

func (s *Store[K, V]) DeleteWithKey(key K) bool {
	lock := s.getLock(key)
	lock.Lock()
	defer lock.Unlock()

	s.evictWithKey(key)
	return true
}

/* private functions */

func (s *Store[K, V]) lazyEvict() {
	if len(s.dict) <= s.maxSize {
		return
	}
	if !s.evictTtlExpired() {
		s.evictLru()
	}
}

func (s *Store[K, V]) evictTtlExpired() bool {
	item := s.expTree.Min()
	if item != nil && item.(BTreeItem[K]).ttl.Before(time.Now()) {
		key := s.expTree.Delete(item).(BTreeItem[K]).key
		s.queue.Remove(s.dict[key])
		delete(s.dict, key)
		return true
	}
	return false
}

func (s *Store[K, V]) evictLru() bool {
	lruItem := s.queue.Back()
	if lruItem != nil {
		dataNode := lruItem.Value.(*DataNode[K, V])
		s.expTree.Delete(BTreeItem[K]{ttl: dataNode.ttl, key: dataNode.key})
		s.queue.Remove(lruItem)
		delete(s.dict, dataNode.key)
		return true
	}
	return false
}

func (s *Store[K, V]) evictWithKey(key K) bool {
	s.expTree.Delete(BTreeItem[K]{ttl: s.dict[key].Value.(*DataNode[K, V]).ttl, key: key})
	s.queue.Remove(s.dict[key])
	delete(s.dict, key)
	return true
}
