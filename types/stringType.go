package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type StringType struct {
	strings map[string]string
	muLock  sync.RWMutex
	expiry  Expiry
}

func NewStringType() *StringType {
	return &StringType{
		strings: make(map[string]string),
		expiry:  NewExpiry(),
		muLock:  sync.RWMutex{},
	}

}

// public functions
func (s *StringType) Append(key, value string) {
	s.muLock.Lock()
	s.strings[key] += value
	s.expiry.updateLRU(key)
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
		s.expiry.updateLRU(key)
	} else {
		s.strings[key] = "0"
		s.expiry.updateLRU(key)
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
	} else {
		s.strings[key] = "0"
	}
	s.expiry.updateLRU(key)
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
	s.expiry.updateLRU(key)
	key, ok = s.expiry.lazyEvict()
	if ok {
		delete(s.strings, key)
	}
	s.muLock.Unlock()
	return value, nil
}

func (s *StringType) GetDel(key string) (string, error) {
	s.muLock.Lock()
	value := s.strings[key]
	delete(s.strings, key)
	s.expiry.removeTTL(key)
	s.expiry.removeLRU(key)
	key, ok := s.expiry.lazyEvict()
	if ok {
		delete(s.strings, key)
	}
	s.muLock.Unlock()
	if !ok {
		return "", errors.New("ERR not found")
	}
	return value, nil
}

func (s *StringType) Set(key, value string) {
	s.muLock.Lock()
	s.strings[key] = value
	s.expiry.updateLRU(key)
	key, ok := s.expiry.lazyEvict()
	if ok {
		delete(s.strings, key)
	}
	s.muLock.Unlock()
}
