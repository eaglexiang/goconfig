package goconfig

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Config interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (value interface{}, err error)
	Bytes() (b []byte, err error)
	Import(b []byte) (err error)
}

func New(filename string) (c EasyConfig, err error) {
	config := newConfig()
	persistC, err := newPersistConfig(config, filename)
	syncC := newSyncConfig(persistC)
	easyC := newEasyConfig(syncC)
	c = easyC
	return
}

type config map[string]interface{}

func newConfig() config {
	return make(config)
}

// Set 更新配置
func (c config) Set(key string, value interface{}) (err error) {
	key = strings.TrimSpace(key)
	if key == "" {
		err = ErrEmptyKey
		return
	}

	path, key := splitPath(key)

	sub, err := c.forceGetLastSub(path)
	if err != nil {
		return
	}

	sub[key] = value

	return
}

func (c config) forceGetLastSub(path []string) (sub config, err error) {
	sub = c

	for i, subKey := range path {
		sub, err = sub.forceGetSub(subKey)
		if err != nil {
			err = fmt.Errorf("%w: %s", err, strings.Join(path[:i], "."))
			break
		}
	}

	return
}

// forceGetSub 获取子配置，获取不到则构造新的
func (c config) forceGetSub(key string) (sub config, err error) {
	subItem, ok := c[key]
	if !ok {
		// 子配置不存在，构造新的
		sub = newConfig()
		c[key] = sub
		return
	}

	sub, ok = subItem.(config)
	if !ok {
		sub, ok = subItem.(map[string]interface{})
	}
	if !ok {
		err = errNotSubConfig(sub, key)
	}

	return
}

// Get 获取配置
func (c config) Get(key string) (value interface{}, err error) {
	key = strings.TrimSpace(key)
	if key == "" {
		err = ErrEmptyKey
		return
	}

	path, key := splitPath(key)

	sub, err := c.getLastSub(path)
	if err != nil {
		return
	}

	value, ok := sub[key]
	if !ok {
		err = errNoKey(key)
	}

	return
}

func (c config) getLastSub(path []string) (sub config, err error) {
	sub = c

	for i, subKey := range path {
		sub, err = sub.getSub(subKey)
		if err != nil {
			err = fmt.Errorf("%w: %s", err, strings.Join(path[:i], "."))
			break
		}
	}

	return
}

func (c config) getSub(key string) (sub config, err error) {
	subItem, ok := c[key]
	if !ok {
		err = ErrNoKey
		return
	}

	sub, ok = subItem.(config)
	if !ok {
		sub, ok = subItem.(map[string]interface{})
	}
	if !ok {
		err = errNotSubConfig(subItem, key)
	}

	return
}

func (c config) Bytes() (b []byte, err error) {
	b, err = json.MarshalIndent(c, "", "    ")
	return
}

func (c config) Import(b []byte) (err error) {
	err = json.Unmarshal(b, &c)
	return
}
