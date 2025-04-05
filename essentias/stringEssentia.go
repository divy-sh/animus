package essentias

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/divy-sh/animus/common"
	"github.com/divy-sh/animus/store"
)

// public functions
func Append(key, value string) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		store.Set(key, value)
		return
	}
	store.Set(key, val+value)
}

func Decr(key string) error {
	return DecrBy(key, "1")
}

func DecrBy(key, value string) error {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	decrVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("ERR invalid decrement value")
	}
	val, ok := store.Get[string, string](key)
	if !ok {
		store.Set(key, fmt.Sprint(-decrVal))
		return nil
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return errors.New("ERR value is not an integer or out of range")
	}
	store.Set(key, fmt.Sprint(intVal-decrVal))
	return nil
}

func Get(key string) (string, error) {
	lock := store.GetLock(key)
	lock.RLock()
	defer lock.RUnlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	return val, nil
}

func GetDel(key string) (string, error) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	store.DeleteWithKey(key)
	return val, nil
}

func GetEx(key, exp string) (string, error) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	expSeconds, err := strconv.ParseInt(exp, 10, 64)
	if err != nil {
		return "", errors.New("ERR invalid expire time")
	}
	store.SetWithTTL(key, val, time.Now().Unix()+expSeconds)
	return val, nil
}

func GetRange(key, start, end string) (string, error) {
	lock := store.GetLock(key)
	lock.RLock()
	defer lock.RUnlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
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

func GetSet(key, value string) (string, error) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	val, ok := store.Get[string, string](key)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	store.Set(key, value)
	return val, nil
}

func Incr(key string) error {
	return IncrBy(key, "1")
}

func IncrBy(key, value string) error {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	incrVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("ERR invalid increment value")
	}
	val, ok := store.Get[string, string](key)
	if !ok {
		store.Set(key, value)
		return nil
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return errors.New("ERR value is not an integer or out of range")
	}
	store.Set(key, fmt.Sprint(intVal+incrVal))
	return nil
}

func IncrByFloat(key, value string) error {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	incrVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("ERR invalid increment value")
	}
	val, ok := store.Get[string, string](key)
	if !ok {
		store.Set(key, value)
		return nil
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return errors.New("ERR value is not a float or out of range")
	}
	store.Set(key, fmt.Sprint(floatVal+incrVal))
	return nil
}

func Set(key, value string) {
	lock := store.GetLock(key)
	lock.Lock()
	defer lock.Unlock()
	store.Set(key, value)
}

func Lcs(key1 string, key2 string, commands []string) (string, error) {
	lock := store.GetLock(key1)
	lock.RLock()
	lock2 := store.GetLock(key2)
	lock2.RLock()
	defer lock.RUnlock()
	defer lock2.RUnlock()

	val1, ok := store.Get[string, string](key1)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	val2, ok := store.Get[string, string](key2)
	if !ok {
		return "", errors.New(common.ERROR_STRING_NOT_FOUND)
	}
	lcs, lcsLen := findLcs(val1, val2)
	if len(commands) == 1 && strings.ToUpper(commands[0]) == "LEN" {
		return fmt.Sprint(lcsLen), nil
	}
	return lcs, nil
}

func MGet(keys *[]string) *[]string {
	values := make([]string, len(*keys))
	for i, key := range *keys {
		lock := store.GetLock(key)
		lock.Lock()
		val, ok := store.Get[string, string](key)
		if !ok {
			values[i] = ""
		} else {
			values[i] = val
		}
		lock.Unlock()
	}
	return &values
}

func MSet(kvPairs *map[string]string) {
	for key, val := range *kvPairs {
		lock := store.GetLock(key)
		lock.Lock()
		store.Set(key, val)
		lock.Unlock()
	}
}

/* PRIVATE FUNCTIONS */

func findLcs(str1, str2 string) (string, int) {
	m, n := len(str1), len(str2)
	if m < n {
		str1, str2 = str2, str1
		m, n = n, m
	}

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := range m {
		for j := range n {
			if str1[i] == str2[j] {
				curr[j+1] = prev[j] + 1
			} else {
				curr[j+1] = max(curr[j], prev[j+1])
			}
		}
		prev, curr = curr, prev
	}

	lcsLen := prev[n]
	lcs := make([]byte, lcsLen)
	i, j, k := m, n, lcsLen

	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			k--
			lcs[k] = str1[i-1]
			i--
			j--
		} else if prev[j] > prev[j-1] {
			i--
		} else {
			j--
		}
	}
	return string(lcs), lcsLen
}
