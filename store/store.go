package store

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/google/btree"
)

/* Types */

type Store struct {
	dict    map[interface{}]*list.Element
	queue   *list.List
	expTree *btree.BTree
	maxSize int
	lock    sync.RWMutex
}

type DataNode struct {
	key  interface{}
	data interface{}
	ttl  time.Time
}

type BTreeItem struct {
	ttl time.Time
	key interface{}
}

func (a BTreeItem) Less(b btree.Item) bool {
	return a.ttl.Before(b.(*BTreeItem).ttl)
}

/* Public functions */

func NewStore() *Store {
	return &Store{
		dict:    make(map[interface{}]*list.Element),
		queue:   list.New(),
		expTree: btree.New(2),
		maxSize: 100_000,
	}
}

func (s *Store) Get(key interface{}) (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	node, found := s.dict[key]
	if !found {
		return nil, false
	} else if node.Value.(*DataNode).ttl.Before(time.Now()) {
		s.evictWithKey(key)
		return nil, false
	}
	fmt.Println(node.Value.(*DataNode).ttl, time.Now())
	s.queue.MoveToFront(node)
	return node.Value.(*DataNode).data, true
}

func (s *Store) Set(key interface{}, value interface{}, ttl time.Time) {
	s.lock.Lock()
	defer s.lock.Unlock()

	node, found := s.dict[key]
	if found {
		node.Value.(*DataNode).data = value
		node.Value.(*DataNode).ttl = ttl
		s.queue.MoveToFront(node)
		s.expTree.Delete(&BTreeItem{ttl: node.Value.(*DataNode).ttl, key: key})
		s.expTree.ReplaceOrInsert(&BTreeItem{ttl, key})
		return
	} else {
		storeItem := &DataNode{
			key:  key,
			data: value,
			ttl:  ttl,
		}
		queueItem := s.queue.PushFront(storeItem)
		s.dict[key] = queueItem
		s.lazyEvict()
	}
}

func (s *Store) DeleteWithKey(key interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.evictWithKey(key)
	return true
}

/* private functions */

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
	if item != nil && item.(*BTreeItem).ttl.Before(time.Now()) {
		key := s.expTree.Delete(item).(*BTreeItem).key
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
		s.expTree.Delete(&BTreeItem{ttl: dataNode.ttl, key: dataNode.key})
		s.queue.Remove(lruItem)
		delete(s.dict, dataNode.key)
		return true
	}
	return false
}

func (s *Store) evictWithKey(key interface{}) bool {
	s.expTree.Delete(&BTreeItem{ttl: s.dict[key].Value.(*DataNode).ttl, key: key})
	s.queue.Remove(s.dict[key])
	delete(s.dict, key)
	return true
}
