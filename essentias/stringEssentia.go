package essentias

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/divy-sh/animus/store"
)

type StringEssentia struct {
	strs store.Store[string, string]
}

func NewStringEssentia() *StringEssentia {
	return &StringEssentia{
		strs: *store.NewStore[string, string](),
	}
}

// public functions
func (s *StringEssentia) Append(key, value string) {
	val, ok := s.strs.Get(key)
	if !ok {
		s.strs.Set(key, value, time.Now().AddDate(1000, 0, 0))
		return
	}
	s.strs.Set(key, val+value, time.Now().AddDate(1000, 0, 0))
}

func (s *StringEssentia) Decr(key string) error {
	return s.DecrBy(key, "1")
}

func (s *StringEssentia) DecrBy(key, value string) error {
	decrVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("ERR invalid decrement value")
	}
	val, ok := s.strs.Get(key)
	if !ok {
		s.strs.Set(key, value, time.Now().AddDate(1000, 0, 0))
		return nil
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return errors.New("ERR value is not an integer or out of range")
	}
	s.strs.Set(key, fmt.Sprint(intVal-decrVal), time.Now().AddDate(1000, 0, 0))
	return nil
}

func (s *StringEssentia) Get(key string) (string, error) {
	val, ok := s.strs.Get(key)
	if !ok {
		return "", errors.New("ERR key not found, or expired")
	}
	return val, nil
}

func (s *StringEssentia) GetDel(key string) (string, error) {
	val, ok := s.strs.Get(key)
	if !ok {
		return "", errors.New("ERR key not found, or expired")
	}
	s.strs.DeleteWithKey(key)
	return val, nil
}

func (s *StringEssentia) GetEx(key, exp string) (string, error) {
	val, ok := s.strs.Get(key)
	if !ok {
		return "", errors.New("ERR key not found, or expired")
	}
	expSeconds, err := strconv.ParseInt(exp, 10, 64)
	if err != nil {
		return "", errors.New("ERR invalid expire time")
	}
	s.strs.Set(key, val, time.Now().Add(time.Duration(expSeconds)*time.Second))
	return val, nil
}

func (s *StringEssentia) GetRange(key, start, end string) (string, error) {
	val, ok := s.strs.Get(key)
	if !ok {
		return "", errors.New("ERR key not found, or expired")
	}
	startInd, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		return "", errors.New("ERR invalid start index")
	}
	endInd, err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		return "", errors.New("ERR invalid end index")
	}
	length := int64(len(val))
	if length == 0 {
		return "", nil
	}
	startInd = (startInd%length + length) % length
	endInd = (endInd%length + length) % length
	if startInd > endInd {
		return "", errors.New("ERR start index greater than end index")
	}
	return val[startInd : endInd+1], nil
}

func (s *StringEssentia) Set(key, value string) {
	s.strs.Set(key, value, time.Now().AddDate(1000, 0, 0))
}
