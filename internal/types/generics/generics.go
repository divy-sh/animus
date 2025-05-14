package generics

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/divy-sh/animus/internal/common"
	"github.com/divy-sh/animus/internal/store"
)

func Copy(source, destination string) (int, error) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	value, ok := store.Get[any, any](source)
	if !ok {
		return 0, errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}
	store.Set(destination, value)
	return 1, nil
}

func Delete(keys *[]string) {
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	for _, key := range *keys {
		store.Delete(key)
	}
}

func Exists(keys *[]string) int {
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
	validKeyCount := 0
	for _, key := range *keys {
		_, exists := store.Get[any, any](key)
		if exists {
			validKeyCount++
		}
	}
	return validKeyCount
}

func Expire(key, seconds, flag string) error {
	secs, _ := strconv.ParseInt(seconds, 10, 64)
	store.GlobalLock.Lock()
	defer store.GlobalLock.Unlock()
	val, ttl, ok := store.GetWithTTL[any, any](key)
	if !ok {
		return errors.New(common.ERR_SOURCE_KEY_NOT_FOUND)
	}
	if ttl == -1 && strings.ToUpper(flag) == common.EXP_XX {
		return errors.New(common.ERR_EXPIRY_TYPE)
	}
	if ttl >= 0 && strings.ToUpper(flag) == common.EXP_NX {
		return errors.New(common.ERR_EXPIRY_TYPE)
	}
	if ttl > time.Now().Unix()+secs && strings.ToUpper(flag) == common.EXP_GT {
		return errors.New(common.ERR_EXPIRY_TYPE)
	}
	if ttl < time.Now().Unix()+secs && strings.ToUpper(flag) == common.EXP_LT {
		return errors.New(common.ERR_EXPIRY_TYPE)
	}
	store.SetWithTTL(key, val, secs)
	return nil
}

func Keys(pattern string) (*[]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil, errors.New(common.ERR_INVALID_REGEX)
	}
	store.GlobalLock.RLock()
	defer store.GlobalLock.RUnlock()
	allKeys := store.GetKeys[string]()
	matchedKeys := []string{}
	for _, key := range *allKeys {
		if re.MatchString(key) {
			matchedKeys = append(matchedKeys, key)
		}
	}
	return &matchedKeys, nil
}
