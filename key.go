package goconfig

import (
	"errors"
	"fmt"
	"strings"
)

const keySplit = "."

var (
	ErrEmptyKey = errors.New("empty key")
	ErrNoKey    = errors.New("key not found")
)

func splitKeys(key string) (keys []string) {
	keys = strings.Split(key, keySplit)
	return
}

func splitPath(key string) (pathKeys []string, realKey string) {
	keys := splitKeys(key)
	pathKeys = keys[:len(keys)-1]
	realKey = keys[len(keys)-1]
	return
}

func errNoKey(key string) (err error) {
	err = fmt.Errorf("%w: %s", ErrNoKey, key)
	return
}
