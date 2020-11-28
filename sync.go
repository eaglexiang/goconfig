package goconfig

import "sync"

// syncConfig 对 config 的同步封装
type syncConfig struct {
	data Config
	l    sync.RWMutex
}

func newSyncConfig(base Config) *syncConfig {
	return &syncConfig{
		data: base,
	}
}

func (c *syncConfig) Set(key string, value interface{}) (err error) {
	c.l.Lock()
	defer c.l.Unlock()

	err = c.data.Set(key, value)

	return
}

func (c *syncConfig) Get(key string) (value interface{}, err error) {
	c.l.RLock()
	defer c.l.RUnlock()

	value, err = c.data.Get(key)

	return
}

func (c *syncConfig) Bytes() (b []byte, err error) {
	c.l.RLock()
	defer c.l.RUnlock()

	b, err = c.data.Bytes()

	return
}

func (c *syncConfig) Import(b []byte) (err error) {
	c.l.Lock()
	defer c.l.RUnlock()

	err = c.data.Import(b)

	return
}
